package client

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/port/out"
)

type UserClient struct{}

var _ out.UserClient = (*UserClient)(nil)

func NewUserClient() *UserClient {
	return &UserClient{}
}

// GetUserByEmail implements client.UserClient.
func (u *UserClient) GetUserByEmail(ctx context.Context, email string) {
	panic("unimplemented")
}

// GetUserByPhoneNumber implements client.UserClient.
func (u *UserClient) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) {
	panic("unimplemented")
}
