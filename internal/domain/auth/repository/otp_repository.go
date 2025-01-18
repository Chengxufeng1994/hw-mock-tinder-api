package repository

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/entity"
)

type OTPRepository interface {
	Save(ctx context.Context, otp *entity.OTP) error
	GetOTPByAccountID(ctx context.Context, accountID string) (*entity.OTP, error)
}
