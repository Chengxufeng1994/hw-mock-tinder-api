package errors

const (
	// ErrUserNotFound - 404: User not found.
	ErrUserNotFound int = iota + 130001

	// ErrUserAlreadyExists - 400: User already exists.
	ErrUserAlreadyExists

	// ErrCreateUserFailed - 400: Create user failed.
	ErrCreateUserFailed
)
