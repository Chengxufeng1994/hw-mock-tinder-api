package out

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/aggregate"
)

type UserClient interface {
	CreateUser(ctx context.Context, user *aggregate.User) (*aggregate.User, error)
	FindUserByAccountID(ctx context.Context, accountID string) (*aggregate.User, error)
}
