package infrastructure

import (
	"github.com/google/wire"
	"go.uber.org/ratelimit"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/port/out"
	accountrepo "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/account/repository"
	authrepo "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/repository"
	userrepo "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/cache"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/client"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/config"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/db"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/oauth"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/account"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/auth"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/user"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/ratelimiter"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/server"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/transaction"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/ws"
)

var InfrastructureSet = wire.NewSet(
	ProviderLoggerConfig,
	ProviderAuthConfig,
	ProviderServerConfig,
	ProviderDatabaseConfig,
	ProviderCacheConfig,
	ProviderFacebookProviderOptionsFromConfig,
	ProviderUberRateLimiter,
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
	user.NewMatchRepository,
	wire.Bind(new(userrepo.MatchRepository), new(*user.MatchRepository)),
	user.NewChatRepository,
	wire.Bind(new(userrepo.ChatRepository), new(*user.ChatRepository)),
	account.NewAccountRepository,
	wire.Bind(new(accountrepo.AccountRepository), new(*account.AccountRepository)),
	auth.NewOTPRepository,
	wire.Bind(new(authrepo.OTPRepository), new(*auth.OTPRepository)),
	ratelimiter.NewUberRateLimiter,
	wire.Bind(new(ratelimiter.RateLimiter), new(*ratelimiter.UberRateLimiter)),
	ws.NewHub,
	db.NewDB,
	cache.NewCache,
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

func ProviderCacheConfig(config *config.Config) *config.Cache {
	return &config.Cache
}

func ProviderFacebookProviderOptionsFromConfig(config *config.Config) oauth.FacebookProviderOptions {
	return oauth.FacebookProviderOptions{}
}

const rateLimit = 100

func ProviderUberRateLimiter() ratelimit.Limiter {
	return ratelimit.New(rateLimit)
}
