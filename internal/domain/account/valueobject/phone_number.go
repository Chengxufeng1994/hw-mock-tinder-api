package valueobject

type PhoneNumber struct {
	value string
}

func NewPhoneNumber(value string) PhoneNumber {
	return PhoneNumber{value: value}
}
