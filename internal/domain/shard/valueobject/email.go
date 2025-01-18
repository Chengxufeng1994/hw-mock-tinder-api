package valueobject

type Email struct {
	value string
}

func NewEmail(value string) Email {
	return Email{value: value}
}

func (e Email) Value() string {
	return e.value
}
