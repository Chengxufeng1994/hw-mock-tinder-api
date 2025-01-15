package user

import (
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/user/command"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/user/query"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

type (
	UserService struct {
		appCommands
		appQueries
	}
	appCommands struct {
		command.UpdateUserHandler
	}
	appQueries struct {
		query.GetUserHandler
	}
)

var _ UserUseCase = (*UserService)(nil)

func NewUserService(logger logging.Logger, users repository.UserRepository) *UserService {
	return &UserService{
		appCommands: appCommands{
			UpdateUserHandler: command.NewUpdateUserCommandHandler(),
		},
		appQueries: appQueries{
			GetUserHandler: query.NewGetUserQueryHandler(),
		},
	}
}
