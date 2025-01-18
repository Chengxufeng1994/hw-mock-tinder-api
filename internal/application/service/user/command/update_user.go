package command

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/entity"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/valueobject"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/transaction"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

type UpdateUserHandler interface {
	UpdateUser(context.Context, UpdateUserCommand) (UpdateUserCommandResult, error)
}

type UpdateUserCommand struct {
	ID           string
	Name         string
	Age          uint
	Gender       string
	Photos       []string
	Interests    []int
	Longitude    float64
	Latitude     float64
	AgeMin       uint
	AgeMax       uint
	GenderFilter string
	Distance     uint
}

type UpdateUserCommandResult struct{}

type UpdateUserCommandHandler struct {
	logger    logging.Logger
	tm        *transaction.TransactionManager
	users     repository.UserRepository
	interests repository.InterestRepository
}

var _ UpdateUserHandler = (*UpdateUserCommandHandler)(nil)

func NewUpdateUserCommandHandler(
	logger logging.Logger, tm *transaction.TransactionManager, users repository.UserRepository, interests repository.InterestRepository,
) *UpdateUserCommandHandler {
	return &UpdateUserCommandHandler{
		logger:    logger.WithName("UpdateUserCommandHandler"),
		tm:        tm,
		users:     users,
		interests: interests,
	}
}

func (h UpdateUserCommandHandler) UpdateUser(ctx context.Context, cmd UpdateUserCommand) (UpdateUserCommandResult, error) {
	err := h.tm.Execute(ctx, func(ctx context.Context) error {
		user, err := h.users.GetUserByID(ctx, cmd.ID)
		if err != nil {
			return err
		}

		interests := make([]valueobject.Interest, 0, len(cmd.Interests))
		for _, interestID := range cmd.Interests {
			interest, err := h.interests.GetInterestByID(ctx, interestID)
			if err != nil {
				return err
			}
			interests = append(interests, interest)
		}

		photos := make([]valueobject.Photo, 0, len(cmd.Photos))
		for _, photo := range cmd.Photos {
			// 計算 SHA256 雜湊值
			hash := sha256.Sum256([]byte(photo))
			id := hex.EncodeToString(hash[:]) // 將雜湊值轉為十六進制字串
			photos = append(photos, valueobject.NewPhoto(id, photo))
		}
		user.UpdateProfile(cmd.Name, cmd.Age, valueobject.NewGender(cmd.Gender), photos, interests, valueobject.NewLocation(cmd.Longitude, cmd.Latitude))

		preference := entity.NewUserPreference(cmd.AgeMin, cmd.AgeMax, valueobject.NewGender(cmd.GenderFilter), cmd.Distance)
		user.UpdatePreference(preference)

		_, err = h.users.Save(ctx, user)
		if err != nil {
			return err
		}
		return err
	})

	if err != nil {
		return UpdateUserCommandResult{}, err
	}

	return UpdateUserCommandResult{}, nil
}
