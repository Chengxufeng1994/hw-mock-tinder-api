package command

import "context"

type UpdateUserHandler interface {
	UpdateUser(context.Context, UpdateUserCommand) (UpdateUserCommandResult, error)
}

type UpdateUserCommand struct {
}

type UpdateUserCommandResult struct{}

type UpdateUserCommandHandler struct{}

var _ UpdateUserHandler = (*UpdateUserCommandHandler)(nil)

func NewUpdateUserCommandHandler() *UpdateUserCommandHandler {
	return &UpdateUserCommandHandler{}
}

func (u *UpdateUserCommandHandler) UpdateUser(context.Context, UpdateUserCommand) (UpdateUserCommandResult, error) {
	panic("unimplemented")
}
