package infrastructure

import (
	"github.com/google/wire"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/port/out"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/client"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/config"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/db"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/oauth"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/user"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/server"
)

var InfrastructureSet = wire.NewSet(
	ProviderLoggerConfig,
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
	wire.Bind(new(repository.UserRepository), new(*user.UserRepository)),
	db.NewDB,
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

func ProviderFacebookProviderOptionsFromConfig(config *config.Config) oauth.FacebookProviderOptions {
	return oauth.FacebookProviderOptions{}
}
