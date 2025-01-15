package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

func LoggerMiddleware(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
