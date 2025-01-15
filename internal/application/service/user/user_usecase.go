package user

import (
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/user/command"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/user/query"
)

type (
	UserUseCase interface {
		command.Commands
		query.Queries
	}
)
