package aggregate

import "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/valueobject"

type Match struct {
	ID      uint
	UserAID string
	UserBID string
	Status  valueobject.MatchStatus
}

func NewMatch(id uint, userAID, userBID string, status valueobject.MatchStatus) *Match {
	return &Match{
		ID:      id,
		UserAID: userAID,
		UserBID: userBID,
		Status:  status,
	}
}
