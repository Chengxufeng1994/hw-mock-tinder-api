package aggregate

import "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/shard/valueobject"

type Account struct {
	ID          string
	Email       valueobject.Email
	PhoneNumber valueobject.PhoneNumber
}

func NewAccount(id string, email string, phoneNumber string) *Account {
	return &Account{
		ID:          id,
		Email:       valueobject.NewEmail(email),
		PhoneNumber: valueobject.NewPhoneNumber(phoneNumber),
	}
}
