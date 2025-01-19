package command

import "context"

type Commands interface {
	CreateUser(context.Context, CreateUserCommand) (CreateUserCommandResult, error)
	UpdateUser(context.Context, UpdateUserCommand) (UpdateUserCommandResult, error)
	SendMessage(context.Context, SendMessageCommand) (SendMessageCommandResult, error)
}
