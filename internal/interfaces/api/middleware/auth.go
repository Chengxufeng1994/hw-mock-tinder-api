package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/authn"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/authn/command"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/valueobject"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/errors"
)

const (
	authenticateHeader    = "Authorization"
	validateAuthDetailKey = "authDetail"
	minTokenParts         = 2
)

func Authn(authnService authn.AuthenticateUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := getTokenFromRequest(c)
		if tokenValue == "" {
			_ = c.Error(errors.MakeTokenInvalid("AuthnMiddleware"))
			return
		}

		cmd := command.VerifyAccessTokenCommand{AccessToken: tokenValue}
		result, err := authnService.VerifyAccessToken(c, cmd)
		if err != nil {
			_ = c.Error(errors.MakeTokenInvalid("AuthnMiddleware.VerifyAccessToken"))
			return
		}

		SetAuthDetailToContext(c, result.AuthDetail)

		c.Next()
	}
}

func getTokenFromRequest(c *gin.Context) string {
	authorization := c.Request.Header.Get("Authorization")
	tokenParts := strings.Split(authorization, " ")
	if len(tokenParts) < minTokenParts {
		return ""
	}

	return tokenParts[1]
}

func SetAuthDetailToContext(c *gin.Context, authDetail valueobject.AuthDetail) {
	c.Set(validateAuthDetailKey, authDetail)
}

func GetAuthDetailFromContext(c *gin.Context) valueobject.AuthDetail {
	if authDetail, ok := c.Get(validateAuthDetailKey); ok {
		if authDetailObj, ok := authDetail.(valueobject.AuthDetail); ok {
			return authDetailObj
		}
	}
	return valueobject.AuthDetail{}
}
