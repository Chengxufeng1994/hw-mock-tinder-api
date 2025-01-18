package auth

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/entity"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/auth/model"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/transaction"
)

type OTPRepository struct {
	tm *transaction.TransactionManager
}

var _ repository.OTPRepository = (*OTPRepository)(nil)

func NewOTPRepository(tm *transaction.TransactionManager) *OTPRepository {
	return &OTPRepository{
		tm: tm,
	}
}

func (r *OTPRepository) GetOTPByAccountID(ctx context.Context, accountID string) (*entity.OTP, error) {
	tx := r.tm.GetGormTransaction(ctx)
	var otp model.OTP
	err := tx.WithContext(ctx).
		Where("account_id = ?", accountID).
		Find(&otp).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return entity.NewOTP(otp.ID, otp.AccountID, otp.Code, otp.IsVerified, otp.Expiration), nil
}

func (r *OTPRepository) Save(ctx context.Context, otp *entity.OTP) error {
	return nil
}
