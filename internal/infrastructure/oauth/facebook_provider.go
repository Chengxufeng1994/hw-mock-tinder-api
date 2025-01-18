package oauth

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/port/out"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/valueobject"
)

type FacebookProviderOptions struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

type facebookUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
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
	response := &facebookUserResponse{
		ID:    "1",
		Name:  "John Doe",
		Email: "john_doe@example.com",
	}

	return valueobject.NewUserInfo(response.ID, response.Name, response.Email)
}
