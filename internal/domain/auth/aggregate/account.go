package aggregate

import "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/auth/entity"

type Account struct {
	ID          string
	Email       string
	PhoneNumber string
	OTP         *entity.OTP
}

func NewAccountWithEmail(id string, email string) *Account {
	return &Account{
		ID:    id,
		Email: email,
	}
}

func NewAccountWithPhoneNumber(id string, phoneNumber string) *Account {
	return &Account{
		ID:          id,
		PhoneNumber: phoneNumber,
	}
}

func (a Account) IsEmpty() bool {
	return a.ID == "" && a.Email == "" && a.PhoneNumber == ""
}
