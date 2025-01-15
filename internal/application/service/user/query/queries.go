package query

import "context"

type Queries interface {
	GetUser(context.Context, GetUserQuery) (GetUserQueryResult, error)
}
