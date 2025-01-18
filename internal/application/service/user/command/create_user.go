package command

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/aggregate"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

type CreateUserHandler interface {
	CreateUser(context.Context, CreateUserCommand) (CreateUserCommandResult, error)
}

type CreateUserCommand struct {
	User *aggregate.User
}

type CreateUserCommandResult struct {
	User *aggregate.User
}

type CreateUserCommandHandler struct {
	logger logging.Logger
	users  repository.UserRepository
}

func NewCreateUserCommandHandler(logger logging.Logger, users repository.UserRepository) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		logger: logger.WithName("CreateUserCommandHandler"),
		users:  users,
	}
}

func (h CreateUserCommandHandler) CreateUser(ctx context.Context, cmd CreateUserCommand) (CreateUserCommandResult, error) {
	newUser := aggregate.NewUser(cmd.User.ID(), cmd.User.AccountID)
	user, err := h.users.Save(ctx, newUser)
	if err != nil {
		return CreateUserCommandResult{}, err
	}
	return CreateUserCommandResult{User: user}, nil
}
