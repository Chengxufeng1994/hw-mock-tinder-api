package cache

import (
	"context"
	"time"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/cache/redis"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/config"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
	SetWithExpire(ctx context.Context, key string, value string, expire time.Duration) error
}

func NewCache(config *config.Cache) Cache {
	switch config.Type {
	case "redis":
		return redis.New(config)
	default:
		return nil
	}
}
