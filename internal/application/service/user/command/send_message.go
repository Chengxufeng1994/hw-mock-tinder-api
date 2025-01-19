package command

import (
	"context"
	"net/http"
	"time"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/valueobject"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/transaction"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/ws"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/errors"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

type SendMessageHandler interface {
	SendMessage(context.Context, SendMessageCommand) (SendMessageCommandResult, error)
}

type SendMessageCommand struct {
	ChatID     string
	SenderID   string
	ReceiverID string
	Content    string
}

type SendMessageCommandResult struct{}

type SendMessageCommandHandler struct {
	logger  logging.Logger
	tm      *transaction.TransactionManager
	users   repository.UserRepository
	matches repository.MatchRepository
	chats   repository.ChatRepository
	hub     *ws.Hub
}

var _ SendMessageHandler = (*SendMessageCommandHandler)(nil)

func NewSendMessageCommandHandler(logger logging.Logger, tm *transaction.TransactionManager, users repository.UserRepository, matches repository.MatchRepository, chats repository.ChatRepository, hub *ws.Hub) *SendMessageCommandHandler {
	return &SendMessageCommandHandler{
		logger:  logger.WithName("SendMessageCommandHandler"),
		tm:      tm,
		users:   users,
		matches: matches,
		chats:   chats,
		hub:     hub,
	}
}

func (s *SendMessageCommandHandler) SendMessage(ctx context.Context, cmd SendMessageCommand) (SendMessageCommandResult, error) {
	chat, err := s.chats.GetChatByID(ctx, cmd.ChatID)
	if err != nil {
		return SendMessageCommandResult{}, errors.NewAppError("SendMessageCommand", "app.chat.not.found", nil, "", http.StatusNotFound, errors.ErrChatNotFound)
	}
	match, err := s.matches.GetMatchByID(ctx, chat.MatchID)
	if err != nil {
		return SendMessageCommandResult{}, errors.NewAppError("SendMessageCommand", "app.match.not.found", nil, "", http.StatusNotFound, errors.ErrMatchNotFound)
	}
	if match.UserAID != cmd.SenderID && match.UserBID != cmd.SenderID {
		return SendMessageCommandResult{}, errors.NewAppError("SendMessageCommand", "app.match.not.found", nil, "", http.StatusNotFound, errors.ErrMatchNotFound)
	}

	chat.AddMessage(valueobject.NewMessage(0, cmd.SenderID, cmd.Content, time.Now()))

	err = s.chats.Save(ctx, chat)
	if err != nil {
		return SendMessageCommandResult{}, err
	}

	message := ws.Message{
		SenderID:   cmd.SenderID,
		ReceiverID: cmd.ReceiverID,
		Message:    cmd.Content,
	}

	s.hub.SendMessage(message)

	return SendMessageCommandResult{}, nil
}
