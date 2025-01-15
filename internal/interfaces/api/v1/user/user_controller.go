package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/user"
)

type UserController struct {
	userService user.UserUseCase
}

func NewUserController(userService user.UserUseCase) UserController {
	return UserController{
		userService: userService,
	}
}

func (u *UserController) GetCurrentUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "hello world"})
}

func (u *UserController) UpdateCurrentUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "hello world"})
}
