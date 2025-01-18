package aggregate

import (
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/entity"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/valueobject"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/ddd"
)

const UserAggregate = "users.UserAggregate"

type User struct {
	ddd.Aggregate
	AccountID string
	Name      string
	Age       uint
	Gender    valueobject.Gender
	Photos    []valueobject.Photo
	Interests []valueobject.Interest
	Location  valueobject.Location
	Status    valueobject.UserStatus

	Preference *entity.UserPreference

	userSnapshot *UserSnapshot
}

var _ ddd.Snapshotter = (*User)(nil)

func NewUser(id string, accountID string) *User {
	return &User{
		Aggregate: ddd.NewAggregate(id, UserAggregate),
		AccountID: accountID,
		Status:    valueobject.UserStatusNew,
	}
}

func (u *User) UpdateProfile(
	name string,
	age uint,
	gender valueobject.Gender,
	photos []valueobject.Photo,
	interests []valueobject.Interest,
	location valueobject.Location,
) {
	if u.Status == valueobject.UserStatusNew {
		u.Status = valueobject.UserStatusActive
	}
	if name != "" {
		u.Name = name
	}
	if age != 0 {
		u.Age = age
	}
	if gender != "" {
		u.Gender = gender
	}
	if location != (valueobject.Location{}) {
		u.Location = location
	}

	u.Photos = photos
	u.Interests = interests
}

func (u *User) UpdatePreference(preference *entity.UserPreference) {
	u.Preference = u.Preference.Update(preference)
}

func HydrateUser(
	id, accountID string,
	name string,
	age uint,
	gender valueobject.Gender,
	photos []valueobject.Photo,
	interests []valueobject.Interest,
	location valueobject.Location,
	status valueobject.UserStatus,
	preference *entity.UserPreference,
) *User {
	return &User{
		Aggregate:  ddd.NewAggregate(id, UserAggregate),
		AccountID:  accountID,
		Name:       name,
		Age:        age,
		Gender:     gender,
		Photos:     photos,
		Interests:  interests,
		Location:   location,
		Status:     status,
		Preference: preference,
	}
}

func (u *User) ApplySnapshot(snapshot ddd.Snapshot) error {
	if snapshot == nil {
		return nil
	}
	switch snapshot.(type) {
	case *UserSnapshot:
	default:
		return nil
	}
	return nil
}

func (u *User) ToSnapshot() {
	u.userSnapshot = &UserSnapshot{
		ID:         u.ID(),
		AccountID:  u.AccountID,
		Name:       u.Name,
		Age:        u.Age,
		Gender:     valueobject.NewGender(u.Gender.String()),
		Photos:     u.Photos,
		Interests:  u.Interests,
		Location:   valueobject.NewLocation(u.Location.Longitude(), u.Location.Latitude()),
		Status:     valueobject.NewUserStatus(u.Status.String()),
		Preference: entity.NewUserPreference(u.Preference.MinAge, u.Preference.MaxAge, u.Preference.Gender, u.Preference.Distance),
	}
}

func (u *User) CheckChange() bool {
	if u.userSnapshot == nil {
		return true
	}

	haveChange := false
	if u.userSnapshot.Name != u.Name {
		haveChange = true
	}
	if u.userSnapshot.Age != u.Age {
		haveChange = true
	}
	if !u.userSnapshot.Gender.IsEqual(u.Gender) {
		haveChange = true
	}
	if !u.userSnapshot.Location.IsEqual(u.Location) {
		haveChange = true
	}

	photosMap := make(map[string]struct{})
	for _, photo := range u.Photos {
		photosMap[photo.ID()] = struct{}{}
	}
	for _, photo := range u.userSnapshot.Photos {
		if _, exists := photosMap[photo.ID()]; !exists {
			haveChange = true
		}
		delete(photosMap, photo.ID())
	}
	if len(photosMap) > 0 {
		haveChange = true
	}

	// 建立當前興趣的映射
	interestsMap := make(map[int]struct{})
	for _, interest := range u.Interests {
		interestsMap[interest.ID()] = struct{}{}
	}

	// 快照中是否有多出的興趣
	for _, interest := range u.userSnapshot.Interests {
		if _, exists := interestsMap[interest.ID()]; !exists {
			// 快照中的興趣在當前興趣中不存在，新增興趣
			haveChange = true
		}
		// 快照興趣已處理，從映射中刪除
		delete(interestsMap, interest.ID())
	}

	// 當前興趣中是否有快照中不存在的興趣
	if len(interestsMap) > 0 {
		// 映射中剩餘的項目表示快照中缺少的興趣
		haveChange = true
	}

	if !u.userSnapshot.Preference.IsEqual(u.Preference) {
		haveChange = true
	}

	return haveChange
}
