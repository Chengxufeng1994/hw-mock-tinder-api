package authn

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/authn/command"
)

type AuthenticateUseCase interface {
	LoginWithFacebook(ctx context.Context, cmd command.LoginWithFacebookCommand) (command.LoginWithFacebookCommandResult, error)
	LoginWithSms(ctx context.Context, cmd command.LoginWithSMSCommand) (command.LoginWithSMSCommandResult, error)
	VerifyAccessToken(ctx context.Context, cmd command.VerifyAccessTokenCommand) (command.VerifyAccessTokenCommandResult, error)
}
