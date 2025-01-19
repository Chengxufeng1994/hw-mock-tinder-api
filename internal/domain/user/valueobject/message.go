package valueobject

import "time"

type Message struct {
	ID        uint
	SenderID  string
	Content   string
	CreatedAt time.Time
}

func NewMessage(id uint, senderID, content string, createdAt time.Time) Message {
	return Message{
		ID:        id,
		SenderID:  senderID,
		Content:   content,
		CreatedAt: createdAt,
	}
}
