package query

import "context"

type GetUserHandler interface {
	GetUser(context.Context, GetUserQuery) (GetUserQueryResult, error)
}

type GetUserQuery struct {
	ID string
}

type GetUserQueryResult struct{}

type GetUserQueryHandler struct{}

var _ GetUserHandler = (*GetUserQueryHandler)(nil)

func NewGetUserQueryHandler() *GetUserQueryHandler {
	return &GetUserQueryHandler{}
}

func (GetUserQueryHandler) GetUser(context.Context, GetUserQuery) (GetUserQueryResult, error) {
	return GetUserQueryResult{}, nil
}
