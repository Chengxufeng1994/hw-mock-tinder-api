package repository

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/aggregate"
)

type MatchRepository interface {
	GetMatchByID(ctx context.Context, matchID string) (*aggregate.Match, error)
	GetMatchedByUserID(ctx context.Context, userID string) ([]*aggregate.Match, error)
}
