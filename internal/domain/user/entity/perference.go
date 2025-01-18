package entity

import "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/valueobject"

type UserPreference struct {
	MinAge   uint
	MaxAge   uint
	Gender   valueobject.Gender
	Distance uint
}

func NewUserPreference(minAge uint, maxAge uint, gender valueobject.Gender, distance uint) *UserPreference {
	return &UserPreference{
		MinAge:   minAge,
		MaxAge:   maxAge,
		Gender:   gender,
		Distance: distance,
	}
}

func (u *UserPreference) Update(other *UserPreference) *UserPreference {
	return &UserPreference{
		MinAge:   other.MinAge,
		MaxAge:   other.MaxAge,
		Gender:   other.Gender,
		Distance: other.Distance,
	}
}

func (u *UserPreference) IsEqual(other *UserPreference) bool {
	return u.MinAge == other.MinAge &&
		u.MaxAge == other.MaxAge &&
		u.Gender.IsEqual(other.Gender) &&
		u.Distance == other.Distance
}
