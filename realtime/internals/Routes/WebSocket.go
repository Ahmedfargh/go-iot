package routes

import (
	"net/http"

	websockets "realtime/internals/WebSockets"
	"realtime/internals/commen"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var notificationService = websockets.NewNotificationService()

func init() {
	go notificationService.Run()
}
func RegisterWebSocketRoutes(r *gin.Engine) {
	r.GET("/ws/notification", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		client_id_str:= c.Query("id")
		client_id, err := strconv.Atoi(client_id_str)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{	
				"error": "Invalid client ID",
			})
			return
		}

		client := commen.NewClient(int32(client_id), conn)
		notificationService.RegisterClient(client)
	})
}
