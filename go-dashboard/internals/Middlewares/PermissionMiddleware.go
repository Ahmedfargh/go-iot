package middlewares

import (
	"net/http"

	Autoloader "go-dashboard/internals/AutoLoader"
	interfaces "go-dashboard/internals/Interfaces"
	models "go-dashboard/internals/Models"

	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, exists := c.Get("admin")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			c.Abort()
			return
		}

		admin := val.(*models.Admin)
		permService := Autoloader.Services_providers["permission"].(interfaces.PermissionInterface)

		if !permService.HasPermissionTo(admin, permission) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to perform this action"})
			c.Abort()
			return
		}

		c.Next()
	}
}
