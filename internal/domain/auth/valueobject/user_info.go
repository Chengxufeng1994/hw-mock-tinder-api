package valueobject

type UserInfo struct {
	ID    string
	Name  string
	Email string
}

func NewUserInfo(id string, name string, email string) UserInfo {
	return UserInfo{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
