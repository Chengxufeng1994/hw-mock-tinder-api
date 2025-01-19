package repository

type UserSpecification interface {
	Query() string
	Value() []any
}

type UserSpecificationByID struct {
	id string
}

var _ UserSpecification = (*UserSpecificationByID)(nil)

func NewUserSpecificationByID(id string) UserSpecificationByID {
	return UserSpecificationByID{id: id}
}

func (UserSpecificationByID) Query() string {
	return "id = ?"
}

func (u UserSpecificationByID) Value() []any {
	return []any{u.id}
}
