package ws

type Message struct {
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Message    string `json:"message"`
}

// type WebsocketEventType string

// const (
// 	WebsocketEventTypeMessage WebsocketEventType = "message"
// )

// type WebSocketMessage interface {
// 	ToJSON() ([]byte, error)
// 	IsValid() bool
// 	EventType() WebsocketEventType
// }
