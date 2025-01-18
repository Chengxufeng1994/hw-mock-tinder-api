package service

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/valueobject"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/config"
)

type TokenService interface {
	GenerateAccessToken(ctx context.Context, authDetail valueobject.AuthDetail) (valueobject.Token, error)
	VerifyAccessToken(ctx context.Context, token valueobject.Token) (valueobject.AuthDetail, error)
}

type JWTTokenService struct {
	config *config.Auth
}

var _ TokenService = (*JWTTokenService)(nil)

func NewJWTTokenService(config *config.Auth) TokenService {
	return &JWTTokenService{config: config}
}

type CustomTokenClaims struct {
	AuthDetail valueobject.AuthDetail `json:"auth_detail"`
	jwt.StandardClaims
}

func (svc *JWTTokenService) GenerateAccessToken(ctx context.Context, authDetail valueobject.AuthDetail) (valueobject.Token, error) {
	duration, err := time.ParseDuration(strconv.Itoa(svc.config.ExpiresTime) + "s")
	if err != nil {
		return valueobject.Token{}, err
	}
	expiresTime := time.Now().Add(duration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomTokenClaims{
		AuthDetail: authDetail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresTime.Unix(),
		},
	})
	tokenString, err := token.SignedString([]byte(svc.config.SecretKey))
	if err != nil {
		return valueobject.Token{}, err
	}
	return valueobject.NewToken(tokenString), nil
}

func (svc *JWTTokenService) VerifyAccessToken(ctx context.Context, token valueobject.Token) (valueobject.AuthDetail, error) {
	tokenClaims, err := jwt.ParseWithClaims(token.Value(), &CustomTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(svc.config.SecretKey), nil
	})
	if err != nil {
		return valueobject.AuthDetail{}, err
	}

	claims, ok := tokenClaims.Claims.(*CustomTokenClaims)
	if !ok {
		return valueobject.AuthDetail{}, err
	}

	return valueobject.NewAuthDetail(claims.AuthDetail.AccountID, claims.AuthDetail.UserID), nil
}
