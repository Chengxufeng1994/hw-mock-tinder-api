package query

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/account/aggregate"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/account/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

type GetAccountByEmailHandler interface {
	GetAccountByEmail(context.Context, GetAccountByEmailQuery) (GetAccountByEmailQueryResult, error)
}

type GetAccountByEmailQuery struct {
	Email string
}

type GetAccountByEmailQueryResult struct {
	Account *aggregate.Account
}

type GetAccountByEmailQueryHandler struct {
	logger   logging.Logger
	accounts repository.AccountRepository
}

var _ GetAccountByEmailHandler = (*GetAccountByEmailQueryHandler)(nil)

func NewGetAccountByEmailQueryHandler(
	logger logging.Logger, accounts repository.AccountRepository,
) *GetAccountByEmailQueryHandler {
	return &GetAccountByEmailQueryHandler{
		logger:   logger.WithName("GetAccountByEmailQueryHandler"),
		accounts: accounts,
	}
}

func (h *GetAccountByEmailQueryHandler) GetAccountByEmail(
	ctx context.Context,
	query GetAccountByEmailQuery,
) (GetAccountByEmailQueryResult, error) {
	account, err := h.accounts.GetAccountByEmail(ctx, query.Email)
	if err != nil {
		return GetAccountByEmailQueryResult{}, err
	}

	return GetAccountByEmailQueryResult{Account: account}, nil
}
