package valueobject

type UserStatus string

const (
	UserStatusNew     UserStatus = "new"
	UserStatusActive  UserStatus = "active"
	UserStatusBlocked UserStatus = "blocked"
	UserStatusDeleted UserStatus = "deleted"
)

func NewUserStatus(status string) UserStatus {
	switch status {
	case UserStatusNew.String():
		return UserStatusNew
	case UserStatusActive.String():
		return UserStatusActive
	case UserStatusBlocked.String():
		return UserStatusBlocked
	case UserStatusDeleted.String():
		return UserStatusDeleted
	default:
		return ""
	}
}

func (s UserStatus) String() string {
	switch s {
	case UserStatusNew, UserStatusActive, UserStatusBlocked, UserStatusDeleted:
		return string(s)
	default:
		return ""
	}
}
