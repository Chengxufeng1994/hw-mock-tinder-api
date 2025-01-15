package valueobject

type Photo struct {
	url string
}

func NewPhoto(url string) Photo {
	return Photo{
		url: url,
	}
}
