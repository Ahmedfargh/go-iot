package main

import (
	"log"
	"realtime/internals/Config"
	"realtime/internals/Models"
	"realtime/internals/Routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	config.GetDbConnection()

	// Auto-migrate the notification table (Laravel-style)
	err := config.DB.AutoMigrate(&models.Notification{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	r := gin.Default()

	// Apply gRPC Auth Middleware
	r.Use(middleware.Authorizied())

	// Register Routes
	routes.RegisterWebSocketRoutes(r)

	log.Println("Realtime Service running on :8081")
	r.Run(":8081")
}
