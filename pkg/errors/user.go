package errors

const (
	// ErrUserNotFound - 404: User not found.
	ErrUserNotFound int = iota + 130001

	// ErrUserAlreadyExists - 400: User already exists.
	ErrUserAlreadyExists

	// ErrCreateUserFailed - 400: Create user failed.
	ErrCreateUserFailed

	// ErrMatchNotFound - 404: Match not found.
	ErrMatchNotFound

	// ErrChatNotFound - 404: Chat not found.
	ErrChatNotFound
)
