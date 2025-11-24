package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() (*Config, error) {
	if err := godotenv.Load("internal/auth-service/.env"); err != nil {
		log.Println("Warning: .env not found:", err)
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}, nil
}
