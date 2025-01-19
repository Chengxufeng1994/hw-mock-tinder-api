package model

import (
	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	UserAID string `gorm:"column:user_a_id;not null"`
	UserBID string `gorm:"column:user_b_id;not null"`
	Status  string `gorm:"column:status;not null"`
}

func (Match) TableName() string {
	return "users.matches"
}
