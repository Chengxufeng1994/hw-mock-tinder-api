package aggregate

import "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/domain/user/valueobject"

type Chat struct {
	ID       string
	MatchID  string
	Messages []valueobject.Message
}

func HydrateChat(id string, matchID string, messages []valueobject.Message) *Chat {
	return &Chat{
		ID:       id,
		MatchID:  matchID,
		Messages: messages,
	}
}

func (c *Chat) AddMessage(message valueobject.Message) {
	c.Messages = append(c.Messages, message)
}
