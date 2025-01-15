package valueobject

type Position struct {
	latitude  float64
	longitude float64
}

func NewPosition(latitude float64, longitude float64) Position {
	return Position{
		latitude:  latitude,
		longitude: longitude,
	}
}
