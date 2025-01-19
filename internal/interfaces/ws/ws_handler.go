package ws

import (
	"github.com/gin-gonic/gin"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/ws"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/interfaces/api/middleware"
)

func WebsocketHandler(hub *ws.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		authDetail := middleware.GetAuthDetailFromContext(c)
		hub.ConnectWebSocket(c.Writer, c.Request, authDetail.UserID)
	}
}
