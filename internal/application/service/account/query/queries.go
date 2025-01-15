package query

import "context"

type Queries interface {
	GetAccountByEmail(context.Context, GetAccountByEmailQuery) (GetAccountByEmailQueryResult, error)
}
