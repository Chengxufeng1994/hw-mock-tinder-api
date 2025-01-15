package client

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/port/out"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/account"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/shard/valueobject"
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

func (client AccountClient) CreateAccount(ctx context.Context, id string, email string, phoneNumber string) (valueobject.AccountInfo, error) {
	panic("unimplemented")
}

func (client AccountClient) FindAccountByEmail(ctx context.Context, email string) (valueobject.AccountInfo, error) {
	panic("unimplemented")
}

func (client AccountClient) FindAccountByPhoneNumber(ctx context.Context, phoneNumber string) (valueobject.AccountInfo, error) {
	panic("unimplemented")
}
