package repository

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/valueobject"
)

type InterestRepository interface {
	Exists(ctx context.Context, id int) (bool, error)
	GetInterestByID(ctx context.Context, id int) (valueobject.Interest, error)
}
