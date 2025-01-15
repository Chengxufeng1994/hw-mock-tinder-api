// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/account"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/authn"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/hello"
	user2 "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/user"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/client"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/config"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/db"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/oauth"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/user"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/server"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/interfaces"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

// Injectors from wire.go:

func InitializeApplication(logger logging.Logger, config2 *config.Config) (app, func(), error) {
	configServer := infrastructure.ProviderServerConfig(config2)
	helloService := hello.NewHelloService(logger)
	accountService := account.NewAccountService()
	accountClient := client.NewAccountClient(accountService)
	facebookProviderOptions := infrastructure.ProviderFacebookProviderOptionsFromConfig(config2)
	facebookProvider := oauth.NewFacebookProvider(facebookProviderOptions)
	authenticateService := authn.NewAuthenticateService(logger, accountClient, facebookProvider)
	database := infrastructure.ProviderDatabaseConfig(config2)
	gormDB, cleanup, err := db.NewDB(database)
	if err != nil {
		return app{}, nil, err
	}
	userRepository := user.NewUserRepository(gormDB)
	userService := user2.NewUserService(logger, userRepository)
	applicationService := application.ProviderService(helloService, authenticateService, accountService, userService)
	handler := interfaces.ProviderRouter(logger, config2, applicationService)
	serverServer := server.NewServer(logger, configServer, handler)
	mainApp := newApp(serverServer)
	return mainApp, func() {
		cleanup()
	}, nil
}
