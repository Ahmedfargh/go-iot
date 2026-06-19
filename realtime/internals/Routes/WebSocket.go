package routes

import (
	"net/http"

	websockets "realtime/internals/WebSockets"
	"realtime/internals/commen"

	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"fmt"
	"realtime/internals/MiddleWare"
	"reflect"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RegisterWebSocketRoutes(r *gin.Engine) {
		r.GET("/ws/notification",middleware.Authorizied(), func(c *gin.Context) {
			val, exists := c.Get("user_id")
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
				return
			}
			fmt.Println(reflect.TypeOf(val),"-",val)
			

			

			conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
			if err != nil {
				fmt.Println("error line 48:-",err)
				log.Printf("Upgrade error: %v", err)
				return
			}
			var clientID int32

			switch v := val.(type) {
				case uint64:
					clientID = int32(v)
				case int64:
					clientID = int32(v)
				case int:
					clientID = int32(v)
				default:
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Unsupported User ID type"})
					return
			}
			client := commen.NewClient(clientID, conn)
			websockets.Notification_service.RegisterClient(client)
	})
}
