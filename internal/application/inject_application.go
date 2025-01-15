package application

import (
	"github.com/google/wire"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/account"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/authn"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/hello"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/user"
)

type ApplicationService struct {
	HelloService   hello.HelloUseCase
	AuthnService   authn.AuthenticateUseCase
	AccountService account.AccountUseCase
	UserService    user.UserUseCase
}

var ApplicationSet = wire.NewSet(
	hello.NewHelloService,
	wire.Bind(new(hello.HelloUseCase), new(*hello.HelloService)),
	authn.NewAuthenticateService,
	wire.Bind(new(authn.AuthenticateUseCase), new(*authn.AuthenticateService)),
	account.NewAccountService,
	wire.Bind(new(account.AccountUseCase), new(*account.AccountService)),
	user.NewUserService,
	wire.Bind(new(user.UserUseCase), new(*user.UserService)),
	ProviderService,
)

func ProviderService(
	helloService hello.HelloUseCase,
	authnService authn.AuthenticateUseCase,
	accountService account.AccountUseCase,
	userService user.UserUseCase,
) *ApplicationService {
	return &ApplicationService{
		HelloService:   helloService,
		AuthnService:   authnService,
		AccountService: accountService,
		UserService:    userService,
	}
}
