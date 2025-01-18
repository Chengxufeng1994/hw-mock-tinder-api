//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/config"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/interfaces"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
	"github.com/google/wire"
)

func InitializeApplication(logger logging.Logger, config *config.Config) (app, func(), error) {
	wire.Build(
		infrastructure.InfrastructureSet,

		application.DomainServiceSet,

		application.ApplicationSet,

		interfaces.InterfacesSet,

		newApp,
	)

	return app{}, nil, nil
}
