package user

import (
	"context"

	"gorm.io/gorm"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/aggregate"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
)

type UserRepository struct {
	db *gorm.DB
}

var _ repository.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*aggregate.User, error) {
	panic("unimplemented")
}

func (u *UserRepository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*aggregate.User, error) {
	panic("unimplemented")
}
