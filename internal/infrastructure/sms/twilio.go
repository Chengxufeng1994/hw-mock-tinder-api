package sms

import "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/port/out"

type TwilioProvider struct{}

var _ out.SmsProvider = (*TwilioProvider)(nil)

func NewTwilioProvider() *TwilioProvider {
	return &TwilioProvider{}
}

func (t *TwilioProvider) SendOTP(phoneNumber string, message string) error {
	panic("unimplemented")
}
