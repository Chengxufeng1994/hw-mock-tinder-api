package valueobject

type AuthDetail struct {
	AccountID string `json:"account_id"`
	UserID    string `json:"user_id"`
}

func NewAuthDetail(accountID, userID string) AuthDetail {
	return AuthDetail{
		AccountID: accountID,
		UserID:    userID,
	}
}
