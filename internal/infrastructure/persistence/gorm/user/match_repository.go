package user

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/aggregate"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/valueobject"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/user/model"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/transaction"
)

type MatchRepository struct {
	tm *transaction.TransactionManager
}

var _ repository.MatchRepository = (*MatchRepository)(nil)

func NewMatchRepository(tm *transaction.TransactionManager) *MatchRepository {
	return &MatchRepository{tm: tm}
}

func (r *MatchRepository) GetMatchedByUserID(ctx context.Context, userID string) ([]*aggregate.Match, error) {
	var matches []model.Match
	tx := r.tm.GetGormTransaction(ctx)
	tx.WithContext(ctx).
		Model(&model.Match{}).
		Select("CASE WHEN user_a_id = ? THEN user_b_id ELSE user_a_id END", userID).
		Where("user_a_id = ? OR user_b_id = ?", userID, userID).
		Where("status = ?", "accepted").
		Find(&matches)

	result := make([]*aggregate.Match, len(matches))
	for i, match := range matches {
		result[i] = aggregate.NewMatch(match.ID, match.UserAID, match.UserBID, valueobject.NewMatchStatus(match.Status))
	}

	return result, nil
}

func (r *MatchRepository) GetMatchByID(ctx context.Context, matchID string) (*aggregate.Match, error) {
	tx := r.tm.GetGormTransaction(ctx)
	var match model.Match
	err := tx.WithContext(ctx).
		Model(&model.Match{}).
		Where("id = ?", matchID).
		Where("status = ?", "accepted").
		First(&match).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return aggregate.NewMatch(match.ID, match.UserAID, match.UserBID, valueobject.NewMatchStatus(match.Status)), nil
}
