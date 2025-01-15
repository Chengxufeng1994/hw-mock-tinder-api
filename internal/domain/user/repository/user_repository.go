package repository

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/aggregate"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*aggregate.User, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*aggregate.User, error)
}
