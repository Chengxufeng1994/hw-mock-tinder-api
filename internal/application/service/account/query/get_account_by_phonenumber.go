package query

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/account/aggregate"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/account/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

type GetAccountByPhoneNumberHandler interface {
	GetAccountByPhoneNumber(context.Context, GetAccountByPhoneNumberQuery) (GetAccountByPhoneNumberQueryResult, error)
}

type GetAccountByPhoneNumberQuery struct {
	PhoneNumber string
}

type GetAccountByPhoneNumberQueryResult struct {
	Account *aggregate.Account
}

type GetAccountByPhoneNumberQueryHandler struct {
	logger   logging.Logger
	accounts repository.AccountRepository
}

var _ GetAccountByPhoneNumberHandler = (*GetAccountByPhoneNumberQueryHandler)(nil)

func NewGetAccountByPhoneNumberQueryHandler(
	logger logging.Logger, accounts repository.AccountRepository,
) *GetAccountByPhoneNumberQueryHandler {
	return &GetAccountByPhoneNumberQueryHandler{
		logger:   logger.WithName("GetAccountByPhoneNumberQueryHandler"),
		accounts: accounts,
	}
}

func (h *GetAccountByPhoneNumberQueryHandler) GetAccountByPhoneNumber(
	ctx context.Context,
	query GetAccountByPhoneNumberQuery,
) (GetAccountByPhoneNumberQueryResult, error) {
	account, err := h.accounts.GetAccountByPhoneNumber(ctx, query.PhoneNumber)
	if err != nil {
		return GetAccountByPhoneNumberQueryResult{}, err
	}

	return GetAccountByPhoneNumberQueryResult{Account: account}, nil
}
