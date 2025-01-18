package errors

const (
	// ErrAccountNotFound - 404: Account not found.
	ErrAccountNotFound int = iota + 120001

	// ErrAccountAlreadyExists - 400: Account already exists.
	ErrAccountAlreadyExists

	// ErrCreateAccountFailed - 400: Create account failed.
	ErrCreateAccountFailed
)
