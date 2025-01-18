package account

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/account/aggregate"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/account/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/shard/valueobject"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/account/model"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/transaction"
)

type AccountRepository struct {
	tm *transaction.TransactionManager
}

var _ repository.AccountRepository = (*AccountRepository)(nil)

func NewAccountRepository(tm *transaction.TransactionManager) *AccountRepository {
	return &AccountRepository{
		tm: tm,
	}
}

func (r *AccountRepository) Save(ctx context.Context, account *aggregate.Account) error {
	tx := r.tm.GetGormTransaction(ctx)

	accountPO := new(model.Account)
	accountPO.ID = account.ID
	accountPO.Email = account.Email.Value()
	accountPO.PhoneNumber = account.PhoneNumber.Value()

	err := tx.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"email",
				"phone_number",
			}),
		}).
		Create(accountPO).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *AccountRepository) GetAccountByEmail(ctx context.Context, email string) (*aggregate.Account, error) {
	tx := r.tm.GetGormTransaction(ctx)

	var account model.Account
	err := tx.WithContext(ctx).
		Where("email = ?", email).
		Find(&account).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &aggregate.Account{
		ID:          account.ID,
		Email:       valueobject.NewEmail(account.Email),
		PhoneNumber: valueobject.NewPhoneNumber(account.PhoneNumber),
	}, nil
}

func (r *AccountRepository) GetAccountByPhoneNumber(ctx context.Context, phoneNumber string) (*aggregate.Account, error) {
	tx := r.tm.GetGormTransaction(ctx)

	var account model.Account
	err := tx.WithContext(ctx).
		Where("phone_number = ?", phoneNumber).
		Find(&account).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &aggregate.Account{
		ID:          account.ID,
		Email:       valueobject.NewEmail(account.Email),
		PhoneNumber: valueobject.NewPhoneNumber(account.PhoneNumber),
	}, nil
}
