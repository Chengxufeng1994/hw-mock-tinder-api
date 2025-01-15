package query

import "context"

type Queries interface {
	SayHello(context.Context, SayHelloQuery) (SayHelloQueryResult, error)
}
