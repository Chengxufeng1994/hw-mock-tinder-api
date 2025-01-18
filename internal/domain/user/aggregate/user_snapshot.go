package aggregate

import (
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/entity"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/valueobject"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/ddd"
)

type UserSnapshot struct {
	ID        string
	AccountID string
	Name      string
	Age       uint
	Gender    valueobject.Gender
	Photos    []valueobject.Photo
	Interests []valueobject.Interest
	Location  valueobject.Location
	Status    valueobject.UserStatus

	Preference *entity.UserPreference
}

var _ ddd.Snapshot = (*UserSnapshot)(nil)

// SnapshotName implements ddd.Snapshot.
func (u *UserSnapshot) SnapshotName() string {
	return "users.UserSnapshot"
}
