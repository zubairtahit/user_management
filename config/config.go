package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DBConfig holds the configuration for the database.
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// LoadDBConfig loads the database configuration from the .env file and validates the required keys.
func LoadDBConfig() *DBConfig {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Required environment variables
	requiredEnvVars := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_SSLMODE",
	}

	// Validate that all required environment variable keys are present
	for _, key := range requiredEnvVars {
		if _, exists := os.LookupEnv(key); !exists {
			log.Fatalf("Missing required environment variable: %s", key)
		}
	}

	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
}
