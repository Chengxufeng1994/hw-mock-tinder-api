package valueobject

type Photo struct {
	id  string
	url string
}

func NewPhoto(id string, url string) Photo {
	return Photo{
		id:  id,
		url: url,
	}
}

func (p Photo) ID() string {
	return p.id
}

func (p Photo) URL() string {
	return p.url
}
