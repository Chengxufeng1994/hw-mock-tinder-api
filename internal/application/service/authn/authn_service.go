package authn

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/port/out"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/authn/command"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

type AuthenticateService struct {
	logger        logging.Logger
	accounts      out.AccountClient
	users         out.UserClient
	oauthProvider out.OAuthProvider
}

var _ AuthenticateUseCase = (*AuthenticateService)(nil)

func NewAuthenticateService(
	logger logging.Logger,
	accounts out.AccountClient,
	users out.UserClient,
	oauthProvider out.OAuthProvider,
) *AuthenticateService {
	return &AuthenticateService{
		logger:        logger.WithName("authn_service"),
		accounts:      accounts,
		users:         users,
		oauthProvider: oauthProvider,
	}
}

func (svc AuthenticateService) LoginWithFacebook(
	ctx context.Context,
	cmd command.LoginWithFacebookCommand,
) (command.LoginWithFacebookCommandResult, error) {
	// TODO: transaction
	// get user info from OAuth provider
	userInfo := svc.oauthProvider.GetUserInfo(ctx, cmd.Token)
	// check if account exists, if not, create account and user
	account, err := svc.accounts.FindAccountByEmail(ctx, userInfo.Email)
	if err != nil {
		return command.LoginWithFacebookCommandResult{}, err
	}
	if account.IsEmpty() {
		if err := svc._createAccount(); err != nil {
			return command.LoginWithFacebookCommandResult{}, err
		}
	}

	// check if user exists, if not, create user
	user, err := svc.users.FindUserByAccountID(ctx, account.ID)
	if err != nil {
		return command.LoginWithFacebookCommandResult{}, err
	}
	if user.IsEmpty() {
		if err := svc.users.CreateUser(ctx, account.ID, userInfo.ID); err != nil {
			return command.LoginWithFacebookCommandResult{}, err
		}
	}

	return command.LoginWithFacebookCommandResult{
		AccessToken: userInfo.ID,
	}, nil
}

func (svc AuthenticateService) LoginWithSms(
	ctx context.Context,
	cmd command.LoginWithSMSCommand,
) (command.LoginWithSMSCommandResult, error) {
	panic("unimplemented")
}

func (svc AuthenticateService) _createAccount() error {
	panic("unimplemented")
}
