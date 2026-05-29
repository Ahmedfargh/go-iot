package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// EnvConfig holds our loaded variables
type EnvConfig struct {
	DBUser               string
	DBPassword           string
	DBHost               string
	DBPort               string
	DBName               string
	RabbitMQURL          string
	JWTSecret            string
	JWTExpirationHours string
}

var AppConfig EnvConfig

func LoadConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, falling back to system environment variables")
	}

	AppConfig = EnvConfig{
		DBUser:               getEnv("DB_USER", "root"),
		DBPassword:           getEnv("DB_PASSWORD", ""),
		DBHost:               getEnv("DB_HOST", "127.0.0.1"),
		DBPort:               getEnv("DB_PORT", "3306"),
		DBName:               getEnv("DB_NAME", "my_app_db"),
		RabbitMQURL:          getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		JWTSecret:            getEnv("JWT_SECRET", "very_secret_key"),
		JWTExpirationHours: getEnv("JWT_EXPIRATION_HOURS", "24"),
	}

	log.Println("Configuration successfully loaded!")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
