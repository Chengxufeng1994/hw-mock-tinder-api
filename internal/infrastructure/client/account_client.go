package client

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/port/out"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/account"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/account/command"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/account/query"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/aggregate"
)

type AccountClient struct {
	accounts account.AccountUseCase
}

var _ out.AccountClient = (*AccountClient)(nil)

func NewAccountClient(accounts account.AccountUseCase) *AccountClient {
	return &AccountClient{
		accounts: accounts,
	}
}

// CreateAccount implements out.AccountClient.
func (client *AccountClient) CreateAccount(ctx context.Context, account *aggregate.Account) (*aggregate.Account, error) {
	result, err := client.accounts.CreateAccount(ctx, command.CreateAccountCommand{
		ID:          account.ID,
		Email:       account.Email,
		PhoneNumber: account.PhoneNumber,
	})

	if err != nil {
		return nil, err
	}

	return &aggregate.Account{
		ID:          result.Account.ID,
		Email:       result.Account.Email.Value(),
		PhoneNumber: result.Account.PhoneNumber.Value(),
	}, nil
}

// FindAccountByEmail implements out.AccountClient.
func (client *AccountClient) FindAccountByEmail(ctx context.Context, email string) (*aggregate.Account, error) {
	response, err := client.accounts.GetAccountByEmail(ctx, query.GetAccountByEmailQuery{
		Email: email,
	})
	if err != nil {
		return nil, err
	}
	return &aggregate.Account{
		ID:          response.Account.ID,
		Email:       response.Account.Email.Value(),
		PhoneNumber: response.Account.PhoneNumber.Value(),
	}, nil
}

// FindAccountByPhoneNumber implements out.AccountClient.
func (client *AccountClient) FindAccountByPhoneNumber(ctx context.Context, phoneNumber string) (*aggregate.Account, error) {
	response, err := client.accounts.GetAccountByPhoneNumber(ctx, query.GetAccountByPhoneNumberQuery{
		PhoneNumber: phoneNumber,
	})
	if err != nil {
		return nil, err
	}
	return &aggregate.Account{
		ID:          response.Account.ID,
		Email:       response.Account.Email.Value(),
		PhoneNumber: response.Account.PhoneNumber.Value(),
	}, nil
}
