package repository

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/account/aggregate"
)

type AccountRepository interface {
	GetAccountByEmail(ctx context.Context, email string) (*aggregate.Account, error)
	GetAccountByPhoneNumber(ctx context.Context, phoneNumber string) (*aggregate.Account, error)
}
