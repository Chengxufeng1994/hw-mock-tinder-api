package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/config"
)

type Cache struct {
	redisClient *redis.Client
}

func New(config *config.Cache) *Cache {
	addr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)

	return &Cache{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: config.Redis.Password,
			DB:       config.Redis.DB,
		}),
	}
}

// Get implements cache.Cache.
func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.redisClient.Get(ctx, key).Val(), nil
}

// Set implements cache.Cache.
func (c *Cache) Set(ctx context.Context, key string, value string) error {
	return c.redisClient.Set(ctx, key, value, 0).Err()
}

// SetWithExpire implements cache.Cache.
func (c *Cache) SetWithExpire(ctx context.Context, key string, value string, expire time.Duration) error {
	return c.redisClient.Set(ctx, key, value, expire).Err()
}
