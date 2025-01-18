package model

import "gorm.io/gorm"

type Preference struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	UserID   string `gorm:"column:user_id;not null"`
	MinAge   uint   `gorm:"column:min_age"`
	MaxAge   uint   `gorm:"column:max_age"`
	Gender   string `gorm:"column:gender"`
	Distance uint   `gorm:"column:distance"`
}

func (Preference) TableName() string {
	return "users.preferences"
}
