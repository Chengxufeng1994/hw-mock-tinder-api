package model

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	MatchID  string `gorm:"column:match_id;not null"`
	Messages []Message
}

func (Chat) TableName() string {
	return "users.chats"
}
