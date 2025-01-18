package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/interfaces/api/response"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/errors"
)

// Global error handler middleware
func ErrorHandler(c *gin.Context) {
	c.Next() // Execute handlers

	// Capture the first error
	err := c.Errors.Last()
	if err == nil {
		return
	}

	// Detect language from header
	// lang := c.GetHeader("Accept-Language")
	// if lang == "" {
	// 	lang = "en" // Default to English
	// }

	// Custom error handling
	switch e := err.Err.(type) {
	case *errors.AppError:
		rid, ok := c.Get(XRequestIDKey)
		if ok {
			if ridStr, ok := rid.(string); ok {
				e.RequestID = ridStr
			}
		}
		fmt.Printf("%#v\n", e)
		c.JSON(e.StatusCode, response.ErrorResponse{
			Code:    e.Code,
			Message: e.Message,
			Details: e.DetailedError,
		})
	default:
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    errors.ErrUnknown,
			Message: "Unknown error",
			Details: err.Error(),
		})
	}
}
