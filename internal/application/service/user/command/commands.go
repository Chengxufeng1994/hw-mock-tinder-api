package command

import "context"

type Commands interface {
	UpdateUser(context.Context, UpdateUserCommand) (UpdateUserCommandResult, error)
}
