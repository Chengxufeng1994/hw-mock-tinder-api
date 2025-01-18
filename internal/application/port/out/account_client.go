package out

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/aggregate"
)

type AccountClient interface {
	CreateAccount(ctx context.Context, account *aggregate.Account) (*aggregate.Account, error)
	FindAccountByEmail(ctx context.Context, email string) (*aggregate.Account, error)
	FindAccountByPhoneNumber(ctx context.Context, phoneNumber string) (*aggregate.Account, error)
}
