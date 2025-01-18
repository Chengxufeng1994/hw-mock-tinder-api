package user

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/valueobject"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/user/model"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/transaction"
)

type InterestRepository struct {
	tm *transaction.TransactionManager
}

var _ repository.InterestRepository = (*InterestRepository)(nil)

func NewInterestRepository(tm *transaction.TransactionManager) *InterestRepository {
	return &InterestRepository{tm: tm}
}

func (r InterestRepository) Exists(ctx context.Context, id int) (bool, error) {
	tx := r.tm.GetGormTransaction(ctx)

	var count int64
	err := tx.WithContext(ctx).
		Model(&model.Interest{}).
		Where("id = ?", id).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *InterestRepository) GetInterestByID(ctx context.Context, id int) (valueobject.Interest, error) {
	tx := r.tm.GetGormTransaction(ctx)

	var interest model.Interest
	err := tx.WithContext(ctx).
		Model(&model.Interest{}).
		Where("id = ?", id).
		First(&interest).Error
	if err != nil {
		return valueobject.Interest{}, err
	}

	return valueobject.NewInterest(interest.ID, interest.Name), nil
}
