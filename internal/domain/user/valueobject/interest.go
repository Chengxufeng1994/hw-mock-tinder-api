package valueobject

type Interest struct {
	id   int
	name string
}

func NewInterest(id int, name string) Interest {
	return Interest{
		id:   id,
		name: name,
	}
}

func (i Interest) ID() int {
	return i.id
}

func (i Interest) Name() string {
	return i.name
}
