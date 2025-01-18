package command

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/account/aggregate"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/account/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

type CreateAccountHandler interface {
	CreateAccount(context.Context, CreateAccountCommand) (CreateAccountCommandResult, error)
}

type CreateAccountCommand struct {
	ID          string
	Email       string
	PhoneNumber string
}

type CreateAccountCommandResult struct {
	Account *aggregate.Account
}

type CreateAccountCommandHandler struct {
	logger   logging.Logger
	accounts repository.AccountRepository
}

var _ CreateAccountHandler = (*CreateAccountCommandHandler)(nil)

func NewCreateAccountCommandHandler(logger logging.Logger, accounts repository.AccountRepository) *CreateAccountCommandHandler {
	return &CreateAccountCommandHandler{
		logger:   logger.WithName("CreateAccountCommandHandler"),
		accounts: accounts,
	}
}

func (h CreateAccountCommandHandler) CreateAccount(ctx context.Context, cmd CreateAccountCommand) (CreateAccountCommandResult, error) {
	account := aggregate.NewAccount(cmd.ID, cmd.Email, cmd.PhoneNumber)
	err := h.accounts.Save(ctx, account)
	if err != nil {
		return CreateAccountCommandResult{}, err
	}
	return CreateAccountCommandResult{Account: account}, nil
}
