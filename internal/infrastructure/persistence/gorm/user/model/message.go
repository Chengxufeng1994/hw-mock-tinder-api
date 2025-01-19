package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ChatID   string `gorm:"column:chat_id;not null"`
	SenderID string `gorm:"column:sender_id;not null"`
	Content  string `gorm:"column:content;not null"`
}

func (Message) TableName() string {
	return "users.messages"
}
