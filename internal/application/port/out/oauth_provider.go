package out

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/shard/valueobject"
)

type OAuthProvider interface {
	GetUserInfo(ctx context.Context, token string) valueobject.UserInfo
}
