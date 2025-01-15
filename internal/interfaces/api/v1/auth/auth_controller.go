package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/authn"
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

func (ctrl AuthnController) LoginWithFacebook(c *gin.Context) {
}

func (ctrl AuthnController) LoginWithSMS(c *gin.Context) {
}
