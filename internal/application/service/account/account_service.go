package account

import "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/account/query"

type (
	AccountService struct {
		appCommands
		appQueries
	}
	appCommands struct {
	}
	appQueries struct {
		query.GetAccountByEmailHandler
		query.GetAccountByPhoneNumberHandler
	}
)

var _ AccountUseCase = (*AccountService)(nil)

func NewAccountService() *AccountService {
	return &AccountService{
		appQueries: appQueries{
			GetAccountByEmailHandler: query.NewGetAccountByEmailQueryHandler(),
		},
		appCommands: appCommands{},
	}
}
