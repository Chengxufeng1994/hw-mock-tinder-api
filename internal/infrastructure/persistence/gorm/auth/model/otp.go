package model

import "time"

type OTP struct {
	ID         int       `gorm:"primaryKey"`
	AccountID  string    `gorm:"column:account_id;unique;not null"`
	Code       string    `gorm:"column:code;not null"`
	IsVerified bool      `gorm:"column:is_verified;not null"`
	Expiration time.Time `gorm:"column:expiration;not null"`
}

func (OTP) TableName() string {
	return "auths.otps"
}
