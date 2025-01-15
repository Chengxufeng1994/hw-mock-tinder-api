package entity

type Interest struct {
	id   string
	name string
}

func NewInterest(id string, name string) Interest {
	return Interest{
		id:   id,
		name: name,
	}
}
