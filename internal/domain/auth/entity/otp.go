package entity

import (
	"strings"
	"time"
)

type OTP struct {
	ID         int
	AccountID  string
	Code       string
	IsVerified bool
	ExpiresAt  time.Time
}

func NewOTP(id int, accountID, code string, isVerified bool, expiresAt time.Time) *OTP {
	return &OTP{
		ID:         id,
		AccountID:  accountID,
		Code:       code,
		IsVerified: isVerified,
		ExpiresAt:  expiresAt,
	}
}

func (otp *OTP) VerifyCode(code string) bool {
	if time.Now().After(otp.ExpiresAt) {
		return false
	}

	if !strings.EqualFold(otp.Code, code) {
		return false
	}

	otp.IsVerified = true
	return true
}
