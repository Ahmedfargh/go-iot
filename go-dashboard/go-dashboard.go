package main

import (
	AutoLoader "go-dashboard/internals/AutoLoader"
	config "go-dashboard/internals/Config"
	routes "go-dashboard/internals/Routes"
	seeders "go-dashboard/internals/Seeders"
)

func InitConfiguration() {
	config.GetDbConnection()
	config.InitRabbitMQ()
	config.MigrateModels()
	AutoLoader.LoadMainServices()
	
	// Seed the database
	seeders.SeedAll()
}

func main() {
	InitConfiguration()

	r := routes.InitializeRoutes()
	
	// Start the server on port 8080
	r.Run(":8080")
}
