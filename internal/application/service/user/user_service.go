package user

import (
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/user/command"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/user/query"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/transaction"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

type (
	UserService struct {
		appCommands
		appQueries
	}
	appCommands struct {
		command.CreateUserHandler
		command.UpdateUserHandler
	}
	appQueries struct {
		query.GetUserByAccountIDHandler
		query.GetUserByIDHandler
		query.GetRecommendationsHandler
		query.GetPreferenceRecommendationsHandler
	}
)

var _ UserUseCase = (*UserService)(nil)

func NewUserService(logger logging.Logger, tm *transaction.TransactionManager, users repository.UserRepository, interests repository.InterestRepository) *UserService {
	return &UserService{
		appCommands: appCommands{
			CreateUserHandler: command.NewCreateUserCommandHandler(logger, users),
			UpdateUserHandler: command.NewUpdateUserCommandHandler(logger, tm, users, interests),
		},
		appQueries: appQueries{
			GetUserByAccountIDHandler:           query.NewGetUserByAccountIDQueryHandler(logger, users),
			GetUserByIDHandler:                  query.NewGetUserQueryHandler(logger, users),
			GetRecommendationsHandler:           query.NewGetRecommendationsQueryHandler(logger, users),
			GetPreferenceRecommendationsHandler: query.NewGetPreferenceRecommendationsQueryHandler(logger, users),
		},
	}
}
