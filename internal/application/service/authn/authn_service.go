package authn

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/port/out"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/authn/command"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/aggregate"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/service"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/valueobject"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/cache"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/config"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/transaction"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/errors"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

// AccountPolicy 定義搜尋與建立帳戶的策略介面
type AccountPolicy interface {
	FindAccount(ctx context.Context, svc *AuthenticateService, identifier string) (*aggregate.Account, error)
	CreateAccount(ctx context.Context, svc *AuthenticateService, identifier string) (*aggregate.Account, error)
}

// EmailAccountPolicy 用於根據 Email 搜尋或建立帳戶
type EmailAccountPolicy struct{}

func (p *EmailAccountPolicy) FindAccount(ctx context.Context, svc *AuthenticateService, email string) (*aggregate.Account, error) {
	return svc.accounts.FindAccountByEmail(ctx, email)
}

func (p *EmailAccountPolicy) CreateAccount(ctx context.Context, svc *AuthenticateService, email string) (*aggregate.Account, error) {
	accountID := uuid.Must(uuid.NewV7()).String()
	newAccount := aggregate.NewAccountWithEmail(accountID, email)
	return svc.accounts.CreateAccount(ctx, newAccount)
}

// PhoneNumberAccountPolicy 用於根據手機號碼搜尋或建立帳戶
type PhoneNumberAccountPolicy struct{}

func (p *PhoneNumberAccountPolicy) FindAccount(ctx context.Context, svc *AuthenticateService, phoneNumber string) (*aggregate.Account, error) {
	return svc.accounts.FindAccountByPhoneNumber(ctx, phoneNumber)
}

func (p *PhoneNumberAccountPolicy) CreateAccount(ctx context.Context, svc *AuthenticateService, phoneNumber string) (*aggregate.Account, error) {
	accountID := uuid.Must(uuid.NewV7()).String()
	newAccount := aggregate.NewAccountWithPhoneNumber(accountID, phoneNumber)
	return svc.accounts.CreateAccount(ctx, newAccount)
}

type AuthenticateService struct {
	logger        logging.Logger
	cfg           *config.Auth
	accounts      out.AccountClient
	users         out.UserClient
	oauthProvider out.OAuthProvider
	tokenService  service.TokenService
	otpRepository repository.OTPRepository
	transaction   *transaction.TransactionManager
	cache         cache.Cache
}

var _ AuthenticateUseCase = (*AuthenticateService)(nil)

func NewAuthenticateService(
	logger logging.Logger,
	cfg *config.Auth,
	accounts out.AccountClient,
	users out.UserClient,
	oauthProvider out.OAuthProvider,
	tokenService service.TokenService,
	otpRepository repository.OTPRepository,
	transaction *transaction.TransactionManager,
	cache cache.Cache,
) *AuthenticateService {
	return &AuthenticateService{
		logger:        logger.WithName("AuthenticateService"),
		cfg:           cfg,
		accounts:      accounts,
		users:         users,
		oauthProvider: oauthProvider,
		otpRepository: otpRepository,
		tokenService:  tokenService,
		transaction:   transaction,
		cache:         cache,
	}
}

func (svc AuthenticateService) LoginWithFacebook(
	ctx context.Context,
	cmd command.LoginWithFacebookCommand,
) (command.LoginWithFacebookCommandResult, error) {
	// get user info from OAuth provider
	userInfo := svc.oauthProvider.GetUserInfo(ctx, cmd.Token)
	accessToken, err := svc.login(ctx, &EmailAccountPolicy{}, userInfo.Email)
	if err != nil {
		return command.LoginWithFacebookCommandResult{}, err
	}
	return command.LoginWithFacebookCommandResult{AccessToken: accessToken.Value()}, nil
}

func (svc AuthenticateService) LoginWithSms(
	ctx context.Context,
	cmd command.LoginWithSMSCommand,
) (command.LoginWithSMSCommandResult, error) {
	accessToken, err := svc.login(ctx, &PhoneNumberAccountPolicy{}, cmd.PhoneNumber)
	if err != nil {
		return command.LoginWithSMSCommandResult{}, err
	}
	return command.LoginWithSMSCommandResult{AccessToken: accessToken.Value()}, nil
}

func (svc *AuthenticateService) login(ctx context.Context, policy AccountPolicy, identifier string) (valueobject.Token, error) {
	var accessToken valueobject.Token
	err := svc.transaction.Execute(ctx, func(txCtx context.Context) error {
		account, err := svc.getOrCreateAccount(txCtx, policy, identifier)
		if err != nil {
			return errors.NewAppError("login", "app.account.not.found", nil, "", http.StatusNotFound, errors.ErrAccountNotFound)
		}

		user, err := svc.getOrCreateUser(txCtx, account.ID)
		if err != nil {
			return errors.NewAppError("login", "app.user.not.found", nil, "", http.StatusNotFound, errors.ErrUserNotFound)
		}

		authDetail := valueobject.NewAuthDetail(account.ID, user.ID)
		accessToken, err = svc.tokenService.GenerateAccessToken(txCtx, authDetail)
		if err != nil {
			return errors.MakeTokenInvalid("login.GenerateAccessToken")
		}

		return svc.cacheAccessToken(ctx, user.ID, accessToken)
	})

	if err != nil {
		return valueobject.Token{}, err
	}

	return accessToken, nil
}

// getOrCreateAccount checks if the account exists, otherwise creates it
func (svc *AuthenticateService) getOrCreateAccount(ctx context.Context, accountPolicy AccountPolicy, identifier string) (*aggregate.Account, error) {
	// check if account exists, if not, create account and user
	account, err := accountPolicy.FindAccount(ctx, svc, identifier)
	if err != nil {
		return nil, err
	}
	if !account.IsEmpty() {
		return account, nil
	}

	svc.logger.
		WithField("identifier", identifier).
		Info("Account not found, creating new account")
	account, err = accountPolicy.CreateAccount(ctx, svc, identifier)
	if err != nil {
		return nil, err
	}
	svc.logger.
		WithField("accountID", account.ID).
		Info("Account created successfully")
	return account, nil
}

// getOrCreateUser checks if the user profile exists, otherwise creates it
func (svc *AuthenticateService) getOrCreateUser(ctx context.Context, accountID string) (*aggregate.User, error) {
	// check if user exists, if not, create user
	user, err := svc.users.FindUserByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}
	if !user.IsEmpty() {
		return user, nil
	}

	svc.logger.
		WithField("accountID", accountID).
		Info("User not found, creating new user profile")
	userID := uuid.Must(uuid.NewV7()).String()
	newUser := aggregate.NewUser(userID, accountID)
	user, err = svc.users.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}
	svc.logger.
		WithField("userID", userID).
		Info("User profile created successfully")
	return user, nil
}

func (svc *AuthenticateService) cacheAccessToken(ctx context.Context, userID string, accessToken valueobject.Token) error {
	return svc.cache.SetWithExpire(ctx,
		fmt.Sprintf("user:%s:accessToken", userID),
		accessToken.Value(),
		time.Duration(svc.cfg.ExpiresTime)*time.Second,
	)
}

func (svc AuthenticateService) VerifyAccessToken(
	ctx context.Context,
	cmd command.VerifyAccessTokenCommand,
) (command.VerifyAccessTokenCommandResult, error) {
	var err error
	var result command.VerifyAccessTokenCommandResult
	result.AuthDetail, err = svc.tokenService.VerifyAccessToken(ctx, valueobject.NewToken(cmd.AccessToken))
	if err != nil {
		return command.VerifyAccessTokenCommandResult{}, err
	}
	accessToken, err := svc.cache.Get(ctx, fmt.Sprintf("user:%s:accessToken", result.AuthDetail.UserID))
	if err != nil {
		return command.VerifyAccessTokenCommandResult{}, err
	}
	if accessToken != cmd.AccessToken {
		return command.VerifyAccessTokenCommandResult{}, errors.MakeTokenInvalid("VerifyAccessTokenCommand")
	}
	return result, nil
}
