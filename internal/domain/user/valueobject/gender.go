package valueobject

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Any    Gender = "any"
)

func NewGender(value string) Gender {
	switch value {
	case "male":
		return Male
	case "female":
		return Female
	case "any":
		return Any
	default:
		return ""
	}
}

func (g Gender) String() string {
	switch g {
	case Male, Female, Any:
		return string(g)
	default:
		return ""
	}
}

func (g Gender) IsEqual(other Gender) bool {
	return g.String() == other.String()
}
