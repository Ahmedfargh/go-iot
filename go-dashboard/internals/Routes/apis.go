package routes

import (
	controllers "go-dashboard/internals/Controllers"
	middlewares "go-dashboard/internals/Middlewares"
	"time"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	r := gin.Default()

	// Serve static files
	r.Static("/uploads", "./uploads")

	adminCtrl := controllers.NewAdminCrudController()
	roleCtrl := controllers.NewRoleController()
	deviceCtrl := controllers.NewDeviceController()
	sectorCtrl := controllers.NewSectorController()

	// Apply global rate limiting
	r.Use(middlewares.RateLimitMiddleware(100, time.Minute))

	// Public routes
	r.POST("/login", adminCtrl.Login)

	// Protected routes
	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleware())
	{
		// Admin Management
		admins := api.Group("/admins")
		{
			admins.GET("/", middlewares.PermissionMiddleware("view-dashboard"), adminCtrl.GetAll)
			admins.POST("/", middlewares.PermissionMiddleware("manage-admins"), adminCtrl.Create)
			admins.GET("/:id", middlewares.PermissionMiddleware("view-dashboard"), adminCtrl.GetByID)
			admins.PUT("/:id", middlewares.PermissionMiddleware("manage-admins"), adminCtrl.Update)
			admins.DELETE("/:id", middlewares.PermissionMiddleware("manage-admins"), adminCtrl.Delete)
		}

		// Role Management
		roles := api.Group("/roles")
		{
			roles.GET("/", middlewares.PermissionMiddleware("manage-roles"), roleCtrl.GetAll)
			roles.POST("/", middlewares.PermissionMiddleware("manage-roles"), roleCtrl.Create)
			roles.POST("/:id/permissions", middlewares.PermissionMiddleware("manage-roles"), roleCtrl.AssignPermissions)
		}

		// Sector Management
		sectors := api.Group("/sectors")
		{
			sectors.GET("/", sectorCtrl.GetAll)
			sectors.POST("/", middlewares.PermissionMiddleware("manage-iot-devices"), sectorCtrl.Create)
			sectors.GET("/:id", sectorCtrl.GetByID)
			sectors.PUT("/:id", middlewares.PermissionMiddleware("manage-iot-devices"), sectorCtrl.Update)
			sectors.DELETE("/:id", middlewares.PermissionMiddleware("manage-iot-devices"), sectorCtrl.Delete)
		}

		// Device Management
		devices := api.Group("/devices")
		{
			devices.GET("/", deviceCtrl.GetAll)
			devices.POST("/", middlewares.PermissionMiddleware("manage-iot-devices"), deviceCtrl.Create)
			devices.GET("/:id", deviceCtrl.GetByID)
			devices.PUT("/:id", middlewares.PermissionMiddleware("manage-iot-devices"), deviceCtrl.Update)
			devices.DELETE("/:id", middlewares.PermissionMiddleware("manage-iot-devices"), deviceCtrl.Delete)
		}
	}

	return r
}
