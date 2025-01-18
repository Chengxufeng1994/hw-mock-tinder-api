package query

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/aggregate"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

type GetUserByAccountIDHandler interface {
	GetUserByAccountID(context.Context, GetUserByAccountIDQuery) (GetUserByAccountIDQueryResult, error)
}

type GetUserByAccountIDQuery struct {
	AccountID string
}

type GetUserByAccountIDQueryResult struct {
	User *aggregate.User
}

type GetUserByAccountIDQueryHandler struct {
	logger logging.Logger
	users  repository.UserRepository
}

var _ GetUserByAccountIDHandler = (*GetUserByAccountIDQueryHandler)(nil)

func NewGetUserByAccountIDQueryHandler(
	logger logging.Logger, users repository.UserRepository,
) *GetUserByAccountIDQueryHandler {
	return &GetUserByAccountIDQueryHandler{
		logger: logger.WithName("GetUserByAccountIDQueryHandler"),
		users:  users,
	}
}

func (h GetUserByAccountIDQueryHandler) GetUserByAccountID(ctx context.Context, query GetUserByAccountIDQuery) (GetUserByAccountIDQueryResult, error) {
	user, err := h.users.GetUserByAccountID(ctx, query.AccountID)
	if err != nil {
		return GetUserByAccountIDQueryResult{}, err
	}
	return GetUserByAccountIDQueryResult{User: user}, nil
}
