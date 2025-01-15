package out

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/shard/valueobject"
)

type AccountClient interface {
	CreateAccount(ctx context.Context, id string, email string, phoneNumber string) (valueobject.AccountInfo, error)
	FindAccountByEmail(ctx context.Context, email string) (valueobject.AccountInfo, error)
	FindAccountByPhoneNumber(ctx context.Context, phoneNumber string) (valueobject.AccountInfo, error)
}
