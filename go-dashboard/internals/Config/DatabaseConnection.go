package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func GetDbConnection() {
	// 1. Safely pull credentials from the environment (loaded via godotenv)
	LoadConfig()
	configStruct := AppConfig
	dbUser := configStruct.DBUser
	dbPass := configStruct.DBPassword
	dbHost := configStruct.DBHost
	dbPort := configStruct.DBPort
	dbName := configStruct.DBName

	// 2. Build the MySQL Data Source Name (DSN) string
	// parseTime=True is crucial for letting GORM translate MySQL timestamps into Go time.Time fields
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	// 3. Establish the GORM connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Logs executed SQL queries into terminal
	})
	if err != nil {
		log.Fatalf("Critical Error: Could not connect to MySQL database: %v", err)
	}

	// 4. Fine-tune performance parameters via the underlying connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to configure underlying SQL driver: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)           // Keep up to 10 idle connections ready
	sqlDB.SetMaxOpenConns(100)          // Allow up to 100 maximum concurrent connections
	sqlDB.SetConnMaxLifetime(time.Hour) // Recycle connections older than 1 hour

	log.Printf("Successfully connected to MySQL database: %s", dbName)
	DB = db
}
