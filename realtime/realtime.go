package main

import (
	"log"
	config "realtime/internals/Config"
	middleware "realtime/internals/MiddleWare"
	models "realtime/internals/Models"
	routes "realtime/internals/Routes"
	"realtime/GRPC"
	"github.com/gin-gonic/gin"
	"realtime/internals/WebSockets"
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
	notification_service:=websockets.NewNotificationService()
	// Apply gRPC Auth Middleware
	middleware.InitAuthClient()
	r.Use(middleware.Authorizied())
	go Gprc.StartGrpcServer()
	// Register Routes
	go notification_service.Run()
	routes.RegisterWebSocketRoutes(r)
	port := ":" + config.AppConfig.HTTP_PORT
	log.Println("Realtime Service running on :8081")
	r.Run(port)
}
