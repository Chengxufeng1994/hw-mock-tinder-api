package infrastructure

import (
	"github.com/google/wire"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/port/out"
	accountrepo "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/account/repository"
	authrepo "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/repository"
	userrepo "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/client"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/config"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/db"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/oauth"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/account"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/auth"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/user"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/server"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/transaction"
)

var InfrastructureSet = wire.NewSet(
	ProviderLoggerConfig,
	ProviderAuthConfig,
	ProviderServerConfig,
	ProviderDatabaseConfig,
	ProviderFacebookProviderOptionsFromConfig,
	oauth.NewFacebookProvider,
	wire.Bind(new(out.OAuthProvider), new(*oauth.FacebookProvider)),
	client.NewUserClient,
	wire.Bind(new(out.UserClient), new(*client.UserClient)),
	client.NewAccountClient,
	wire.Bind(new(out.AccountClient), new(*client.AccountClient)),
	user.NewUserRepository,
	wire.Bind(new(userrepo.UserRepository), new(*user.UserRepository)),
	user.NewInterestRepository,
	wire.Bind(new(userrepo.InterestRepository), new(*user.InterestRepository)),
	account.NewAccountRepository,
	wire.Bind(new(accountrepo.AccountRepository), new(*account.AccountRepository)),
	auth.NewOTPRepository,
	wire.Bind(new(authrepo.OTPRepository), new(*auth.OTPRepository)),
	db.NewDB,
	transaction.NewTransactionManager,
	server.NewServer,
)

func ProviderLoggerConfig(config *config.Config) *config.Logging {
	return &config.Logging
}

func ProviderServerConfig(config *config.Config) *config.Server {
	return &config.Server
}

func ProviderDatabaseConfig(config *config.Config) *config.Database {
	return &config.Database
}

func ProviderAuthConfig(config *config.Config) *config.Auth {
	return &config.Auth
}

func ProviderFacebookProviderOptionsFromConfig(config *config.Config) oauth.FacebookProviderOptions {
	return oauth.FacebookProviderOptions{}
}
