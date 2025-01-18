package valueobject

type Token struct {
	value string
}

func NewToken(value string) Token {
	return Token{value: value}
}

func (t Token) Value() string {
	return t.value
}
