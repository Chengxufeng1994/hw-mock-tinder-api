package out

type SmsProvider interface {
	SendOTP(phoneNumber string, otp string) error
}
