package user

import (
	"context"
	"errors"

	"github.com/twpayne/go-geom"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/aggregate"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/entity"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/valueobject"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/shard"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/persistence/gorm/user/model"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/transaction"
)

type UserRepository struct {
	tm *transaction.TransactionManager
}

var _ repository.UserRepository = (*UserRepository)(nil)

func NewUserRepository(tm *transaction.TransactionManager) *UserRepository {
	return &UserRepository{tm: tm}
}

func (r UserRepository) Save(ctx context.Context, user *aggregate.User) (*aggregate.User, error) {
	tx := r.tm.GetGormTransaction(ctx)
	haveChange := user.CheckChange()
	if !haveChange {
		return user, nil
	}

	var userPO model.User
	userPO.ID = user.ID()
	userPO.AccountID = user.AccountID
	userPO.Name = user.Name
	userPO.Age = user.Age
	userPO.Gender = user.Gender.String()
	userPO.Photos = r.photoListFromDO(user.Photos)
	userPO.Interests = r.interestListFromDO(user.Interests)
	userPO.Location = shard.GeoPoint{
		Point: *geom.NewPointFlat(shard.SRID, []float64{
			user.Location.Longitude(),
			user.Location.Latitude(),
		}),
	}
	userPO.Status = user.Status.String()
	userPO.Preference = model.Preference{
		MinAge:   user.Preference.MinAge,
		MaxAge:   user.Preference.MaxAge,
		Gender:   user.Preference.Gender.String(),
		Distance: user.Preference.Distance,
	}

	err := tx.WithContext(ctx).
		Model(&model.User{
			ID: user.ID(),
		}).
		Association("Photos").
		Clear()
	if err != nil {
		return nil, err
	}

	err = tx.WithContext(ctx).
		Model(&model.User{
			ID: user.ID(),
		}).
		Association("Interests").
		Clear()
	if err != nil {
		return nil, err
	}

	err = tx.WithContext(ctx).
		Model(&model.User{}).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"name",
				"age",
				"gender",
				"location",
				"status",
			}),
		}).
		Create(&userPO).Error

	if err != nil {
		return nil, err
	}

	photos := r.photoListToDO(userPO.Photos)
	interests := r.interestListToDO(userPO.Interests)
	preference := entity.NewUserPreference(userPO.Preference.MinAge, userPO.Preference.MaxAge, valueobject.NewGender(userPO.Preference.Gender), userPO.Preference.Distance)

	newUser := aggregate.HydrateUser(
		userPO.ID,
		userPO.AccountID,
		userPO.Name,
		userPO.Age,
		valueobject.NewGender(userPO.Gender),
		photos,
		interests,
		valueobject.NewLocation(userPO.Location.X(), userPO.Location.Y()),
		valueobject.NewUserStatus(userPO.Status),
		preference,
	)
	newUser.ToSnapshot()

	return newUser, nil
}

func (r UserRepository) GetUserByAccountID(ctx context.Context, accountID string) (*aggregate.User, error) {
	tx := r.tm.GetGormTransaction(ctx)

	var userPO model.User
	err := tx.WithContext(ctx).
		Preload("Photos").
		Preload("Interests").
		Preload("Preference").
		Where("account_id = ?", accountID).
		First(&userPO).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	photos := r.photoListToDO(userPO.Photos)
	interests := r.interestListToDO(userPO.Interests)
	preference := entity.NewUserPreference(userPO.Preference.MinAge, userPO.Preference.MaxAge, valueobject.NewGender(userPO.Preference.Gender), userPO.Preference.Distance)

	newUser := aggregate.HydrateUser(
		userPO.ID,
		userPO.AccountID,
		userPO.Name,
		userPO.Age,
		valueobject.NewGender(userPO.Gender),
		photos,
		interests,
		valueobject.NewLocation(userPO.Location.X(), userPO.Location.Y()),
		valueobject.NewUserStatus(userPO.Status),
		preference,
	)
	newUser.ToSnapshot()

	return newUser, nil
}

func (r UserRepository) GetUserByID(ctx context.Context, id string) (*aggregate.User, error) {
	tx := r.tm.GetGormTransaction(ctx)

	var userPO model.User
	err := tx.WithContext(ctx).
		Preload("Photos").
		Preload("Interests").
		Preload("Preference").
		Where("id = ?", id).
		First(&userPO).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	photos := r.photoListToDO(userPO.Photos)
	interests := r.interestListToDO(userPO.Interests)
	preference := entity.NewUserPreference(userPO.Preference.MinAge, userPO.Preference.MaxAge, valueobject.NewGender(userPO.Preference.Gender), userPO.Preference.Distance)

	newUser := aggregate.HydrateUser(
		userPO.ID,
		userPO.AccountID,
		userPO.Name,
		userPO.Age,
		valueobject.NewGender(userPO.Gender),
		photos,
		interests,
		valueobject.NewLocation(userPO.Location.X(), userPO.Location.Y()),
		valueobject.NewUserStatus(userPO.Status),
		preference,
	)
	newUser.ToSnapshot()

	return newUser, nil
}

func (r *UserRepository) GetRecommendations(ctx context.Context, params repository.SearchParams) ([]*aggregate.User, error) {
	const kilometer = 1000
	tx := r.tm.GetGormTransaction(ctx)

	var users []model.User
	query := tx.WithContext(ctx).
		Model(&model.User{}).
		Preload("Photos").
		Preload("Interests").
		Preload("Preference")

	query = query.Where("ST_DWithin(location, ST_SetSRID(ST_MakePoint(?, ?), 4326), ?)", params.Longitude, params.Latitude, params.Distance*kilometer)
	query = query.Where("age BETWEEN ? AND ?", params.AgeMin, params.AgeMax)
	if params.Gender != "any" {
		query = query.Where("gender = ?", params.Gender)
	} else {
		query = query.Where("gender != ?", params.Gender)
	}
	query = query.Limit(params.Count)

	err := query.Find(&users).Error
	if err != nil {
		return nil, err
	}

	usersDO := make([]*aggregate.User, len(users))
	for i, user := range users {
		photos := r.photoListToDO(user.Photos)
		interests := r.interestListToDO(user.Interests)
		preference := entity.NewUserPreference(user.Preference.MinAge, user.Preference.MaxAge, valueobject.NewGender(user.Preference.Gender), user.Preference.Distance)
		newUser := aggregate.HydrateUser(
			user.ID,
			user.AccountID,
			user.Name,
			user.Age,
			valueobject.NewGender(user.Gender),
			photos,
			interests,
			valueobject.NewLocation(user.Location.X(), user.Location.Y()),
			valueobject.NewUserStatus(user.Status),
			preference,
		)
		usersDO[i] = newUser
	}

	return usersDO, nil
}

func (r UserRepository) interestToDO(interestPO model.Interest) valueobject.Interest {
	return valueobject.NewInterest(interestPO.ID, interestPO.Name)
}

func (r UserRepository) interestListToDO(interests []model.Interest) []valueobject.Interest {
	interestsDO := make([]valueobject.Interest, 0, len(interests))
	for _, interest := range interests {
		interestsDO = append(interestsDO, r.interestToDO(interest))
	}
	return interestsDO
}

func (r UserRepository) interestFromDO(interest valueobject.Interest) model.Interest {
	return model.Interest{
		ID:   interest.ID(),
		Name: interest.Name(),
	}
}

func (r UserRepository) interestListFromDO(interests []valueobject.Interest) []model.Interest {
	interestsDO := make([]model.Interest, 0, len(interests))
	for _, interest := range interests {
		interestsDO = append(interestsDO, r.interestFromDO(interest))
	}
	return interestsDO
}

func (r UserRepository) photoToDO(photoPO model.Photo) valueobject.Photo {
	return valueobject.NewPhoto(photoPO.ID, photoPO.URL)
}

func (r UserRepository) photoListToDO(photos []model.Photo) []valueobject.Photo {
	photosDO := make([]valueobject.Photo, 0, len(photos))
	for _, photo := range photos {
		photosDO = append(photosDO, r.photoToDO(photo))
	}
	return photosDO
}

func (r UserRepository) photoFromDO(photo valueobject.Photo) model.Photo {
	return model.Photo{
		ID:  photo.ID(),
		URL: photo.URL(),
	}
}

func (r UserRepository) photoListFromDO(photos []valueobject.Photo) []model.Photo {
	photosDO := make([]model.Photo, 0, len(photos))
	for _, photo := range photos {
		photosDO = append(photosDO, r.photoFromDO(photo))
	}
	return photosDO
}
