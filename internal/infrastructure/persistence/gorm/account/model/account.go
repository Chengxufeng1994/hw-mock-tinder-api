package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Email       string `gorm:"column:email;unique"`
	PhoneNumber string `gorm:"column:phone_number;unique"`
}

func (Account) TableName() string {
	return "accounts.accounts"
}
