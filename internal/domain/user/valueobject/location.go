package valueobject

type Location struct {
	longitude float64
	latitude  float64
}

func NewLocation(longitude float64, latitude float64) Location {
	return Location{
		longitude: longitude,
		latitude:  latitude,
	}
}

func (l Location) Longitude() float64 {
	return l.longitude
}

func (l Location) Latitude() float64 {
	return l.latitude
}

func (l Location) IsEqual(other Location) bool {
	return l.Longitude() == other.Longitude() && l.Latitude() == other.Latitude()
}
