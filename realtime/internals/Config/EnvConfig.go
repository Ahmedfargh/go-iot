package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DBUser         string
	DBPassword     string
	DBHost         string
	DBPort         string
	DBName         string
	HTTP_PORT      string
	MAIN_DASHBOARD string
}

var AppConfig EnvConfig

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, falling back to system environment variables")
	}

	AppConfig = EnvConfig{
		DBUser:         getEnv("DB_USER", "root"),
		DBPassword:     getEnv("DB_PASSWORD", ""),
		DBHost:         getEnv("DB_HOST", "127.0.0.1"),
		DBPort:         getEnv("DB_PORT", "3306"),
		DBName:         getEnv("REALTIME_DB_NAME", "realtime"),
		HTTP_PORT:      getEnv("REALTIME_HTTP_PORT", "8081"),
		MAIN_DASHBOARD: getEnv("MAIN_GO_DASHBOARD", "localhost:50051"),
	}

	log.Println("Realtime Configuration successfully loaded!")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
