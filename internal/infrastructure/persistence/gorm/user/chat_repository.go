package user

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/aggregate"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/valueobject"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/user/model"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/transaction"
)

type ChatRepository struct {
	tm *transaction.TransactionManager
}

var _ repository.ChatRepository = (*ChatRepository)(nil)

func NewChatRepository(tm *transaction.TransactionManager) *ChatRepository {
	return &ChatRepository{tm: tm}
}

func (r *ChatRepository) Save(ctx context.Context, chat *aggregate.Chat) error {
	tx := r.tm.GetGormTransaction(ctx)

	chatPO := new(model.Chat)
	chatPO.ID = chat.ID
	chatPO.MatchID = chat.MatchID
	chatPO.Messages = make([]model.Message, 0, len(chat.Messages))
	for _, message := range chat.Messages {
		chatPO.Messages = append(chatPO.Messages, model.Message{
			Model:    gorm.Model{ID: message.ID},
			SenderID: message.SenderID,
			Content:  message.Content,
			ChatID:   chat.ID,
		})
	}

	err := tx.Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).
			Save(chatPO).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (r ChatRepository) GetChatByID(ctx context.Context, id string) (*aggregate.Chat, error) {
	tx := r.tm.GetGormTransaction(ctx)
	var chat model.Chat
	err := tx.WithContext(ctx).
		Preload("Messages").
		Where("id = ?", id).
		First(&chat).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	messages := make([]valueobject.Message, 0, len(chat.Messages))
	for _, message := range chat.Messages {
		messages = append(messages, valueobject.NewMessage(message.ID, message.SenderID, message.Content, message.CreatedAt))
	}

	return aggregate.HydrateChat(chat.ID, chat.MatchID, messages), nil
}
