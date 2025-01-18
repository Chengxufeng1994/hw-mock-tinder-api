package user

import (
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/aggregate"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/entity"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/valueobject"
)

type UserMapper struct{}

func NewUserMapper() UserMapper {
	return UserMapper{}
}

func (m UserMapper) ToDto(user *aggregate.User) User {
	photos := make([]Photo, 0, len(user.Photos))
	for _, photo := range user.Photos {
		photos = append(photos, Photo{
			ID:  photo.ID(),
			URL: photo.URL(),
		})
	}

	interests := make([]Interest, len(user.Interests))
	for i, interest := range user.Interests {
		interests[i] = Interest{
			ID:   interest.ID(),
			Name: interest.Name(),
		}
	}

	return User{
		ID:        user.ID(),
		AccountID: user.AccountID,
		Name:      user.Name,
		Age:       user.Age,
		Gender:    user.Gender.String(),
		Photos:    photos,
		Interests: interests,
		Longitude: user.Location.Longitude(),
		Latitude:  user.Location.Latitude(),
		Preference: Preference{
			AgeMin:   user.Preference.MinAge,
			AgeMax:   user.Preference.MaxAge,
			Gender:   user.Preference.Gender.String(),
			Distance: user.Preference.Distance,
		},
	}
}

func (m UserMapper) ToDomain(user User) *aggregate.User {
	photos := make([]valueobject.Photo, len(user.Photos))
	for i, photo := range user.Photos {
		photos[i] = valueobject.NewPhoto(photo.ID, photo.URL)
	}

	interests := make([]valueobject.Interest, len(user.Interests))
	for i, interest := range user.Interests {
		interests[i] = valueobject.NewInterest(interest.ID, interest.Name)
	}

	preference := entity.NewUserPreference(
		user.Preference.AgeMin,
		user.Preference.AgeMax,
		valueobject.NewGender(user.Preference.Gender),
		user.Preference.Distance,
	)

	return aggregate.HydrateUser(
		user.ID,
		user.AccountID,
		user.Name,
		user.Age,
		valueobject.NewGender(user.Gender),
		photos,
		interests,
		valueobject.NewLocation(user.Latitude, user.Longitude),
		valueobject.UserStatusActive,
		preference,
	)
}
