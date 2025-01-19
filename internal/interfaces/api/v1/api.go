package api

import (
	"github.com/gin-gonic/gin"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/ratelimiter"
	wsinfra "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/ws"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/interfaces/api/middleware"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/interfaces/api/v1/auth"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/interfaces/api/v1/user"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/interfaces/ws"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

func SetupRouter(
	logger logging.Logger,
	router *gin.RouterGroup,
	application *application.ApplicationService,
	rl ratelimiter.RateLimiter,
	hub *wsinfra.Hub,
) {
	authController := auth.NewAuthnController(logger, application.AuthnService)
	{
		authGroup := router.Group("/auth")
		authGroup.Use(middleware.RateLimit(rl))
		authGroup.POST("/login/facebook", authController.LoginWithFacebook)
		authGroup.POST("/login/sms", authController.LoginWithSMS)
	}

	userController := user.NewUserController(application.UserService)
	{
		userGroup := router.Group("/users")
		userGroup.Use(middleware.Authn(application.AuthnService))
		userGroup.GET("/chats", ws.WebsocketHandler(hub))
		userGroup.POST("/chats/message", userController.SendMessage)
		userGroup.GET("", userController.GetCurrentUser)
		userGroup.PUT("", userController.UpdateCurrentUser)
		userGroup.GET("/preferences/recommendations", userController.GetPreferenceRecommendations)
		userGroup.GET("/recommendations", userController.GetRecommendations)
	}
}
