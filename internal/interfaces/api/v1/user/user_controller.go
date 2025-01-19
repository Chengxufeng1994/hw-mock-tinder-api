package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/user"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/user/command"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/user/query"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/interfaces/api/middleware"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/interfaces/api/response"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/errors"
)

type UserController struct {
	userService user.UserUseCase
	userMapper  UserMapper
}

func NewUserController(userService user.UserUseCase) UserController {
	return UserController{
		userService: userService,
		userMapper:  NewUserMapper(),
	}
}

func (ctrl UserController) GetCurrentUser(c *gin.Context) {
	authDetail := middleware.GetAuthDetailFromContext(c)

	result, err := ctrl.userService.GetUserByID(c.Request.Context(), query.GetUserByIDQuery{
		ID: authDetail.UserID,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}

	userDto := ctrl.userMapper.ToDto(result.User)

	c.JSON(http.StatusOK, response.OKResponse{Data: userDto})
}

func (ctrl UserController) UpdateCurrentUser(c *gin.Context) {
	var req UpdateCurrentUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		apErr := errors.MakeBindError("authController.LoginWithFacebook", err.Error())
		_ = c.Error(apErr)
		return
	}

	authDetail := middleware.GetAuthDetailFromContext(c)
	_, err := ctrl.userService.UpdateUser(c.Request.Context(), command.UpdateUserCommand{
		ID:           authDetail.UserID,
		Name:         req.Name,
		Age:          req.Age,
		Gender:       req.Gender,
		Photos:       req.Photos,
		Interests:    req.Interests,
		Longitude:    req.Longitude,
		Latitude:     req.Latitude,
		AgeMin:       req.AgeMin,
		AgeMax:       req.AgeMax,
		GenderFilter: req.GenderFilter,
		Distance:     req.Distance,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OKResponse{})
}

func (ctrl UserController) GetPreferenceRecommendations(c *gin.Context) {
	authDetail := middleware.GetAuthDetailFromContext(c)
	result, err := ctrl.userService.GetPreferenceRecommendations(c.Request.Context(), query.GetPreferenceRecommendationsQuery{UserID: authDetail.UserID})
	if err != nil {
		_ = c.Error(err)
		return
	}

	usersDto := make([]User, len(result.Users))
	for i, user := range result.Users {
		userDto := ctrl.userMapper.ToDto(user)
		usersDto[i] = userDto
	}

	c.JSON(http.StatusOK, response.OKResponse{Data: usersDto})
}

func (ctrl UserController) GetRecommendations(c *gin.Context) {
	latitude := c.DefaultQuery("latitude", "0.0")
	longitude := c.DefaultQuery("longitude", "0.0")
	ageMin := c.DefaultQuery("age_min", "18")
	ageMax := c.DefaultQuery("age_max", "30")
	gender := c.DefaultQuery("gender", "any")
	distance := c.DefaultQuery("distance", "5.0")
	count := c.DefaultQuery("count", "10")

	lat, _ := strconv.ParseFloat(latitude, 64)
	lon, _ := strconv.ParseFloat(longitude, 64)
	minAge, _ := strconv.ParseUint(ageMin, 10, 64)
	maxAge, _ := strconv.ParseUint(ageMax, 10, 64)
	dist, _ := strconv.ParseUint(distance, 10, 64)
	cnt, _ := strconv.Atoi(count)

	result, err := ctrl.userService.GetRecommendations(c.Request.Context(), query.GetRecommendationsQuery{
		Latitude:  lat,
		Longitude: lon,
		AgeMin:    uint(minAge),
		AgeMax:    uint(maxAge),
		Gender:    gender,
		Distance:  uint(dist),
		Count:     cnt,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}
	usersDto := make([]User, len(result.Users))
	for i, user := range result.Users {
		userDto := ctrl.userMapper.ToDto(user)
		usersDto[i] = userDto
	}

	c.JSON(http.StatusOK, response.OKResponse{Data: usersDto})
}

func (ctrl UserController) SendMessage(c *gin.Context) {
	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apErr := errors.MakeBindError("userController.SendMessage", err.Error())
		_ = c.Error(apErr)
		return
	}

	authDetail := middleware.GetAuthDetailFromContext(c)

	_, err := ctrl.userService.SendMessage(c.Request.Context(), command.SendMessageCommand{
		ChatID:     req.ChatID,
		SenderID:   authDetail.UserID,
		ReceiverID: req.ReceiverID,
		Content:    req.Content,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OKResponse{})
}
