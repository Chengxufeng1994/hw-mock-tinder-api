package aggregate

import (
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/entity"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/valueobject"
)

type User struct {
	ID        string
	Name      string
	Photo     []valueobject.Photo
	Interests []entity.Interest
	Position  valueobject.Position
	Distance  uint
}
