package command

import "context"

type Commands interface {
	CreateAccount(context.Context, CreateAccountCommand) (CreateAccountCommandResult, error)
}
