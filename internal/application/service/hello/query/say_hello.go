package query

import "context"

type SayHelloHandler interface {
	Handle(context.Context, SayHelloQuery) (SayHelloQueryResult, error)
}

type SayHelloQuery struct {
	Name string
}

type SayHelloQueryResult struct {
	Message string
}

type SayHelloQueryHandler struct{}

var _ SayHelloHandler = (*SayHelloQueryHandler)(nil)

func NewSayHelloQueryHandler() *SayHelloQueryHandler {
	return &SayHelloQueryHandler{}
}

func (SayHelloQueryHandler) Handle(ctx context.Context, query SayHelloQuery) (SayHelloQueryResult, error) {
	return SayHelloQueryResult{
		Message: "Hello " + query.Name,
	}, nil
}
