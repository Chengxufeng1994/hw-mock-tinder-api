package query

import "context"

type GetAccountByEmailHandler interface {
	GetAccountByEmail(context.Context, GetAccountByEmailQuery) (GetAccountByEmailQueryResult, error)
}

type GetAccountByEmailQuery struct {
	Email string
}

type GetAccountByEmailQueryResult struct{}

type GetAccountByEmailQueryHandler struct{}

var _ GetAccountByEmailHandler = (*GetAccountByEmailQueryHandler)(nil)

func NewGetAccountByEmailQueryHandler() *GetAccountByEmailQueryHandler {
	return &GetAccountByEmailQueryHandler{}
}

func (h *GetAccountByEmailQueryHandler) GetAccountByEmail(
	ctx context.Context,
	query GetAccountByEmailQuery,
) (GetAccountByEmailQueryResult, error) {
	panic("unimplemented")
}
