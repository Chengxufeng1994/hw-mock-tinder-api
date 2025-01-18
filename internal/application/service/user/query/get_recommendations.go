package query

import (
	"context"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/aggregate"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/repository"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

type GetRecommendationsHandler interface {
	GetRecommendations(context.Context, GetRecommendationsQuery) (GetRecommendationsQueryResult, error)
}

type GetRecommendationsQuery struct {
	Latitude  float64
	Longitude float64
	AgeMin    uint
	AgeMax    uint
	Gender    string
	Distance  uint
	Count     int
}

type GetRecommendationsQueryResult struct {
	Users []*aggregate.User
}

type GetRecommendationsQueryHandler struct {
	logger logging.Logger
	users  repository.UserRepository
}

func NewGetRecommendationsQueryHandler(logger logging.Logger, users repository.UserRepository) *GetRecommendationsQueryHandler {
	return &GetRecommendationsQueryHandler{
		logger: logger.WithName("GetRecommendationsQueryHandler"),
		users:  users,
	}
}

func (h GetRecommendationsQueryHandler) GetRecommendations(ctx context.Context, query GetRecommendationsQuery) (GetRecommendationsQueryResult, error) {
	h.logger.WithFields(
		map[string]interface{}{
			"latitude":  query.Latitude,
			"longitude": query.Longitude,
			"ageMin":    query.AgeMin,
			"ageMax":    query.AgeMax,
			"gender":    query.Gender,
			"distance":  query.Distance,
			"count":     query.Count,
		}).
		Info("GetRecommendations")

	users, err := h.users.GetRecommendations(ctx, repository.SearchParams{
		Latitude:  query.Latitude,
		Longitude: query.Longitude,
		AgeMin:    query.AgeMin,
		AgeMax:    query.AgeMax,
		Gender:    query.Gender,
		Distance:  query.Distance,
		Count:     query.Count,
	})
	if err != nil {
		return GetRecommendationsQueryResult{}, err
	}

	return GetRecommendationsQueryResult{Users: users}, nil
}
