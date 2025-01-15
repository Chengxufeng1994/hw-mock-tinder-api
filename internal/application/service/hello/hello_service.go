package hello

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/hello/query"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

type (
	HelloService struct {
		appQueries
		appCommand
	}

	appQueries struct {
		*query.SayHelloQueryHandler
	}

	appCommand struct{}
)

var _ HelloUseCase = (*HelloService)(nil)

func NewHelloService(logger logging.Logger) *HelloService {
	sayHelloQueryHandler := query.NewSayHelloQueryHandler()

	return &HelloService{
		appQueries: appQueries{
			SayHelloQueryHandler: sayHelloQueryHandler,
		},
		appCommand: appCommand{},
	}
}

func (h *HelloService) SayHello(ctx context.Context, query query.SayHelloQuery) (query.SayHelloQueryResult, error) {
	return h.SayHelloQueryHandler.Handle(ctx, query)
}
