package query

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/aggregate"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

const COUNT = 60

type GetPreferenceRecommendationsHandler interface {
	GetPreferenceRecommendations(context.Context, GetPreferenceRecommendationsQuery) (GetPreferenceRecommendationsQueryResult, error)
}

type GetPreferenceRecommendationsQuery struct {
	UserID string
}

type GetPreferenceRecommendationsQueryResult struct {
	Users []*aggregate.User
}

type GetPreferenceRecommendationsQueryHandler struct {
	logger logging.Logger
	users  repository.UserRepository
}

func NewGetPreferenceRecommendationsQueryHandler(logger logging.Logger, users repository.UserRepository) *GetPreferenceRecommendationsQueryHandler {
	return &GetPreferenceRecommendationsQueryHandler{
		logger: logger.WithName("GetPreferenceRecommendationsHandler"),
		users:  users,
	}
}

func (h GetPreferenceRecommendationsQueryHandler) GetPreferenceRecommendations(ctx context.Context, query GetPreferenceRecommendationsQuery) (GetPreferenceRecommendationsQueryResult, error) {
	user, err := h.users.GetUserByID(ctx, query.UserID)
	if err != nil {
		return GetPreferenceRecommendationsQueryResult{}, err
	}

	users, err := h.users.GetRecommendations(ctx, repository.SearchParams{
		Latitude:  user.Location.Latitude(),
		Longitude: user.Location.Longitude(),
		AgeMin:    user.Preference.MinAge,
		AgeMax:    user.Preference.MaxAge,
		Gender:    user.Preference.Gender.String(),
		Distance:  user.Preference.Distance,
		Count:     COUNT,
	})
	if err != nil {
		return GetPreferenceRecommendationsQueryResult{}, err
	}

	return GetPreferenceRecommendationsQueryResult{Users: users}, nil
}
