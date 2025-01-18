package account

import (
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/account/command"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/account/query"
)

type AccountUseCase interface {
	command.Commands
	query.Queries
}
