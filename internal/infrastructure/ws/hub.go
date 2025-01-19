package ws

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/gorilla/websocket"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/errors"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

const (
	socketMaxMessageSizeKb = 8 * 1024
	broadcastQueueSize     = 4096
)

type Hub struct {
	mutex   *sync.Mutex
	clients []*WsConn

	logger         logging.Logger
	allowedOrigins []string

	register   chan *WsConn
	unregister chan *WsConn

	broadcast chan Message
}

func NewHub(logger logging.Logger) *Hub {
	return &Hub{
		mutex:   &sync.Mutex{},
		clients: make([]*WsConn, 0),

		logger:         logger.WithName("Hub"),
		allowedOrigins: []string{"*"},

		register:   make(chan *WsConn),
		unregister: make(chan *WsConn),
		broadcast:  make(chan Message, broadcastQueueSize),
	}
}

func (hub *Hub) Run(ctx context.Context) error {
	hub.logger.Info("starting websocket server")
	for {
		select {
		case <-ctx.Done():
			return nil
		case client := <-hub.register:
			hub.onConnect(client)
		case client := <-hub.unregister:
			hub.onDisconnect(client)
		case message := <-hub.broadcast:
			hub.onBroadcast(message)
		}
	}
}

func (hub *Hub) onBroadcast(msg Message) {
	receiverID := msg.ReceiverID
	for _, client := range hub.clients {
		if client.UserID == receiverID {
			client.send <- msg
		}
	}
}

func (hub *Hub) onConnect(client *WsConn) {
	hub.logger.WithFields(map[string]interface{}{"remote_addr": client.ws.RemoteAddr().String()}).Info("client connected")

	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	hub.clients = append(hub.clients, client)
}

func (hub *Hub) onDisconnect(client *WsConn) {
	hub.logger.WithFields(map[string]interface{}{"remote_addr": client.ws.RemoteAddr().String()}).Info("client disconnected")

	client.Close()
	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	i := -1
	for j, c := range hub.clients {
		if c.UserID == client.UserID {
			i = j
			break
		}
	}
	copy(hub.clients[i:], hub.clients[i+1:])
	hub.clients[len(hub.clients)-1] = nil
	hub.clients = hub.clients[:len(hub.clients)-1]
}

func (hub *Hub) checkOrigin(r *http.Request) bool {
	origin := r.Header["Origin"]
	if len(origin) == 0 {
		return true
	}
	u, err := url.Parse(origin[0])
	if err != nil {
		return false
	}
	var allow bool
	for _, o := range hub.allowedOrigins {
		if o == u.Host {
			allow = true
			break
		}
		if o == "*" {
			allow = true
			break
		}
	}
	if !allow {
		hub.logger.Infof("none of allowed origins: %s matched: %s\n", strings.Join(hub.allowedOrigins, ", "), u.Host)
	}
	return allow
}

func (hub *Hub) ConnectWebSocket(w http.ResponseWriter, r *http.Request, userID string) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  socketMaxMessageSizeKb,
		WriteBufferSize: socketMaxMessageSizeKb,
		CheckOrigin:     hub.checkOrigin,
	}

	if userID == "" {
		appErr := errors.NewAppError("HandleWebSocket", "app.websocket.upgrade.failed", nil, "", http.StatusInternalServerError, errors.ErrWebSocketUpgradeFailed)
		if appErr != nil {
			log.Println(appErr)
		}
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		appErr := errors.NewAppError("HandleWebSocket", "app.websocket.upgrade.failed", nil, "", http.StatusInternalServerError, errors.ErrWebSocketUpgradeFailed)
		if appErr != nil {
			log.Println(appErr)
		}
		return
	}

	client := NewWsConn(hub, ws, userID)
	hub.register <- client

	client.Pump()
}

func (hub *Hub) SendMessage(message Message) {
	hub.broadcast <- message
}
