package query

import "context"

type GetAccountByPhoneNumberHandler interface {
	GetAccountByPhoneNumber(context.Context, GetAccountByPhoneNumberQuery) (GetAccountByPhoneNumberQueryResult, error)
}

type GetAccountByPhoneNumberQuery struct {
	PhoneNumber string
}

type GetAccountByPhoneNumberQueryResult struct{}

type GetAccountByPhoneNumberQueryHandler struct{}

var _ GetAccountByPhoneNumberHandler = (*GetAccountByPhoneNumberQueryHandler)(nil)

func NewGetAccountByPhoneNumberQueryHandler() *GetAccountByPhoneNumberQueryHandler {
	return &GetAccountByPhoneNumberQueryHandler{}
}

func (h *GetAccountByPhoneNumberQueryHandler) GetAccountByPhoneNumber(
	ctx context.Context,
	query GetAccountByPhoneNumberQuery,
) (GetAccountByPhoneNumberQueryResult, error) {
	panic("unimplemented")
}
