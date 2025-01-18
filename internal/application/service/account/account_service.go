package account

import (
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/account/command"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/account/query"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/account/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

type (
	AccountService struct {
		appCommands
		appQueries
	}
	appCommands struct {
		command.CreateAccountHandler
	}
	appQueries struct {
		query.GetAccountByEmailHandler
		query.GetAccountByPhoneNumberHandler
	}
)

var _ AccountUseCase = (*AccountService)(nil)

func NewAccountService(logger logging.Logger, accounts repository.AccountRepository) *AccountService {
	return &AccountService{
		appCommands: appCommands{
			CreateAccountHandler: command.NewCreateAccountCommandHandler(logger, accounts),
		},
		appQueries: appQueries{
			GetAccountByEmailHandler:       query.NewGetAccountByEmailQueryHandler(logger, accounts),
			GetAccountByPhoneNumberHandler: query.NewGetAccountByPhoneNumberQueryHandler(logger, accounts),
		},
	}
}
