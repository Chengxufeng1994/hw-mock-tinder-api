package errors

const (
	// ErrOTPNotFound - 404: OTP not found.
	ErrOTPNotFound int = iota + 110001

	// ErrOTPInvalid - 400: OTP invalid.
	ErrOTPInvalid
)
