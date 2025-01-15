package api

import (
	"github.com/gin-gonic/gin"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/interfaces/api/v1/auth"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/interfaces/api/v1/user"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

func SetupRouter(
	logger logging.Logger,
	router *gin.RouterGroup,
	application *application.ApplicationService,
) {
	authController := auth.NewAuthnController(logger, application.AuthnService)
	{
		authGroup := router.Group("/auth")
		authGroup.POST("/login/facebook", authController.LoginWithFacebook)
		authGroup.POST("/login/sms", authController.LoginWithSMS)
	}

	userController := user.NewUserController(application.UserService)
	{
		userGroup := router.Group("/user")
		userGroup.GET("/", userController.GetCurrentUser)
		userGroup.POST("/", userController.UpdateCurrentUser)
	}
}
