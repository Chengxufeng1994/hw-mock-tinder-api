package valueobject

type AccountInfo struct {
	ID          string
	Email       string
	PhoneNumber string
}

func NewAccountInfo(id string, email string, phoneNumber string) AccountInfo {
	return AccountInfo{
		ID:          id,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}

func (accountInfo AccountInfo) IsEmpty() bool {
	return accountInfo.ID == "" && accountInfo.Email == "" && accountInfo.PhoneNumber == ""
}
