package client

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/port/out"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/user"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/user/command"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/user/query"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/aggregate"
	useraggregate "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/aggregate"
)

type UserClient struct {
	userService user.UserUseCase
}

var _ out.UserClient = (*UserClient)(nil)

func NewUserClient(userService user.UserUseCase) *UserClient {
	return &UserClient{
		userService: userService,
	}
}

func (client UserClient) CreateUser(ctx context.Context, user *aggregate.User) (*aggregate.User, error) {
	result, err := client.userService.CreateUser(ctx, command.CreateUserCommand{
		User: useraggregate.NewUser(user.ID, user.AccountID),
	})

	if err != nil {
		return nil, err
	}

	return &aggregate.User{
		ID:        result.User.ID(),
		AccountID: result.User.AccountID,
	}, nil
}

func (client UserClient) FindUserByAccountID(ctx context.Context, accountID string) (*aggregate.User, error) {
	result, err := client.userService.GetUserByAccountID(ctx, query.GetUserByAccountIDQuery{
		AccountID: accountID,
	})

	if err != nil {
		return nil, err
	}
	if result.User == nil {
		return &aggregate.User{}, nil
	}

	return &aggregate.User{
		ID:        result.User.ID(),
		AccountID: result.User.AccountID,
	}, nil
}
