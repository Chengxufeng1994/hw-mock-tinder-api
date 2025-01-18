package command

import "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/valueobject"

type VerifyAccessTokenCommand struct {
	AccessToken string
}

type VerifyAccessTokenCommandResult struct {
	AuthDetail valueobject.AuthDetail
}
