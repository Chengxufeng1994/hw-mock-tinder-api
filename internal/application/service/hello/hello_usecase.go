package hello

import "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/hello/query"

type (
	HelloUseCase interface {
		query.Queries
	}
)
