package application

import (
	"github.com/google/wire"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/service"
)

var DomainServiceSet = wire.NewSet(
	service.NewJWTTokenService,
)
