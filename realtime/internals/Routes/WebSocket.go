package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterWebSocketRoutes(r *gin.Engine) {
	r.GET("/ws", func(c *gin.Context) {
		// Placeholder for WebSocket upgrade logic
		c.JSON(http.StatusOK, gin.H{
			"message": "WebSocket endpoint reached",
		})
	})
}
