package repository

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/aggregate"
)

type ChatRepository interface {
	Save(ctx context.Context, chat *aggregate.Chat) error
	GetChatByID(ctx context.Context, chatID string) (*aggregate.Chat, error)
}
