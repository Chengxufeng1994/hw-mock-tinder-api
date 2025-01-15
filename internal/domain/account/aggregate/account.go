package aggregate

import "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/account/valueobject"

type Account struct {
	ID          string
	Email       valueobject.Email
	PhoneNumber valueobject.PhoneNumber
}
