package aggregate

type User struct {
	ID        string
	AccountID string
}

func NewUser(id string, accountID string) *User {
	return &User{
		ID:        id,
		AccountID: accountID,
	}
}

func (u User) IsEmpty() bool {
	return u.ID == "" && u.AccountID == ""
}
