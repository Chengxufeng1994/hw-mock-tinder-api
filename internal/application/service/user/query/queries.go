package query

import "context"

type Queries interface {
	GetUserByAccountID(context.Context, GetUserByAccountIDQuery) (GetUserByAccountIDQueryResult, error)
	GetUserByID(context.Context, GetUserByIDQuery) (GetUserByIDQueryResult, error)
	GetRecommendations(context.Context, GetRecommendationsQuery) (GetRecommendationsQueryResult, error)
	GetPreferenceRecommendations(context.Context, GetPreferenceRecommendationsQuery) (GetPreferenceRecommendationsQueryResult, error)
}
