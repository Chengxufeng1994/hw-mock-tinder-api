package repository

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/aggregate"
)

type SearchParams struct {
	Latitude  float64
	Longitude float64
	Distance  uint
	AgeMin    uint
	AgeMax    uint
	Gender    string
	Count     int
}

type UserRepository interface {
	Save(ctx context.Context, user *aggregate.User) (*aggregate.User, error)
	GetUserByID(ctx context.Context, id string) (*aggregate.User, error)
	GetUserByAccountID(ctx context.Context, accountID string) (*aggregate.User, error)
	GetRecommendations(ctx context.Context, searchParams SearchParams) ([]*aggregate.User, error)
}
