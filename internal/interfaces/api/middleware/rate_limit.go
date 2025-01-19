package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/ratelimiter"
)

func RateLimit(ratelimiter ratelimiter.RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !ratelimiter.Allow(c.ClientIP()) {
			return
		}
		c.Next()
	}
}
