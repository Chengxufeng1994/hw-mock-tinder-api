package ws

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const pingMultiplier = 6

var (
	// Time allowed to write a message to the peer.
	writeWaitTime = 30 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWaitTime = 100 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingInterval = (pongWaitTime * pingMultiplier) / 10
	// Maximum message size allowed from peer.
	SocketMaxMessageSizeKb int64 = 64 * 1024
	// Maximum size of the write buffer per connection.
	sendQueueSize = 256
)

type WsConnOptions struct{}

type WsConn struct {
	hub                *Hub
	ws                 *websocket.Conn
	UserID             string
	lastUserActivityAt int64

	send         chan Message
	endWritePump chan struct{}
	pumpFinished chan struct{}
}

func NewWsConn(hub *Hub, wc *websocket.Conn, userID string) *WsConn {
	client := &WsConn{
		hub:                hub,
		ws:                 wc,
		UserID:             userID,
		lastUserActivityAt: time.Now().UnixMilli(),
		send:               make(chan Message, sendQueueSize),
		endWritePump:       make(chan struct{}),
		pumpFinished:       make(chan struct{}),
	}
	return client
}

func (c *WsConn) Close() {
	c.ws.Close()
	<-c.pumpFinished
}

// Pump starts the client instance. Aft er this, the websocket
// is ready to send/receive messages.
func (c *WsConn) Pump() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.WritePump()
	}()

	c.ReadPump()
	close(c.endWritePump)
	wg.Wait()
	close(c.pumpFinished)
}

func (c *WsConn) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.ws.Close()
	}()
	c.ws.SetReadLimit(SocketMaxMessageSizeKb)
	err := c.ws.SetReadDeadline(time.Now().Add(pongWaitTime))
	if err != nil {
		fmt.Println("websocket.SetReadDeadline")
		return
	}
	c.ws.SetPongHandler(func(string) error {
		if err := c.ws.SetReadDeadline(time.Now().Add(pongWaitTime)); err != nil {
			return err
		}
		return nil
	})
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		var msg Message
		_ = json.Unmarshal(message, &msg)
		c.hub.broadcast <- msg
	}
}

func (c *WsConn) WritePump() {
	ticker := time.NewTicker(pingInterval)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				if err := c.writeMessageBuf(websocket.CloseMessage, []byte{}); err != nil {
					fmt.Println("websocket.send", err)
				}
				return
			}
			_ = c.ws.WriteMessage(websocket.TextMessage, []byte(message.Message))

		case <-ticker.C:
			if err := c.ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				fmt.Println("websocket.ticker", err)
				return
			}

		case <-c.endWritePump:
			return
		}
	}
}

func (c *WsConn) writeMessageBuf(msgType int, data []byte) error {
	if err := c.ws.SetWriteDeadline(time.Now().Add(writeWaitTime)); err != nil {
		return err
	}
	return c.ws.WriteMessage(msgType, data)
}
