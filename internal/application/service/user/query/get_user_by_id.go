package query

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/aggregate"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

type GetUserByIDHandler interface {
	GetUserByID(context.Context, GetUserByIDQuery) (GetUserByIDQueryResult, error)
}

type GetUserByIDQuery struct {
	ID string
}

type GetUserByIDQueryResult struct {
	User *aggregate.User
}

type GetUserByIDQueryHandler struct {
	logger logging.Logger
	users  repository.UserRepository
}

var _ GetUserByIDHandler = (*GetUserByIDQueryHandler)(nil)

func NewGetUserQueryHandler(logger logging.Logger, users repository.UserRepository) *GetUserByIDQueryHandler {
	return &GetUserByIDQueryHandler{
		logger: logger.WithName("GetUserByIDQueryHandler"),
		users:  users,
	}
}

func (h GetUserByIDQueryHandler) GetUserByID(ctx context.Context, query GetUserByIDQuery) (GetUserByIDQueryResult, error) {
	user, err := h.users.GetUserByID(ctx, query.ID)
	if err != nil {
		return GetUserByIDQueryResult{}, err
	}
	return GetUserByIDQueryResult{
		User: user,
	}, nil
}
