package oauth

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/port/out"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/shard/valueobject"
)

type FacebookProviderOptions struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

type FacebookProvider struct {
	options FacebookProviderOptions
}

var _ out.OAuthProvider = (*FacebookProvider)(nil)

// Mock Facebook Provider
func NewFacebookProvider(options FacebookProviderOptions) *FacebookProvider {
	return &FacebookProvider{options: options}
}

func (f *FacebookProvider) GetUserInfo(ctx context.Context, token string) valueobject.UserInfo {
	panic("unimplemented")
}
