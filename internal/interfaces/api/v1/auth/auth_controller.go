package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/authn"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/authn/command"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/interfaces/api/response"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/errors"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

type AuthnController struct {
	logger       logging.Logger
	authnService authn.AuthenticateUseCase
}

func NewAuthnController(logger logging.Logger, authnService authn.AuthenticateUseCase) *AuthnController {
	return &AuthnController{
		logger:       logger.WithName("authn_handler"),
		authnService: authnService,
	}
}

type LoginWithFacebookRequest struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (ctrl AuthnController) LoginWithFacebook(c *gin.Context) {
	var req LoginWithFacebookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		apErr := errors.MakeBindError("authController.LoginWithFacebook", err.Error())
		_ = c.Error(apErr)
		return
	}

	result, err := ctrl.authnService.LoginWithFacebook(c.Request.Context(), command.LoginWithFacebookCommand{
		Token: req.Token,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}

	resp := LoginResponse{
		AccessToken: result.AccessToken,
	}

	c.JSON(http.StatusOK, response.OKResponse{Data: resp})
}

type LoginWithSmsRequest struct {
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
}

func (ctrl AuthnController) LoginWithSMS(c *gin.Context) {
	var req LoginWithSmsRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		apErr := errors.MakeBindError("authController.LoginWithSMS", err.Error())
		_ = c.Error(apErr)
		return
	}

	result, err := ctrl.authnService.LoginWithSms(c.Request.Context(), command.LoginWithSMSCommand{
		PhoneNumber: req.PhoneNumber,
		Code:        req.Code,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}

	resp := LoginResponse{
		AccessToken: result.AccessToken,
	}

	c.JSON(http.StatusOK, response.OKResponse{Data: resp})
}
