package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// Config represents the application configuration.
type Config struct {
	ServerPort  int
	DatabaseURL string
	// Add other configuration parameters as needed
}

// LoadConfig loads configuration from a .env file.
func LoadConfig(filePath string) (*Config, error) {
	if err := godotenv.Load(filePath); err != nil {
		log.Printf("Error loading .env file, using defaults: %v", err)
	}

	cfg := &Config{
		ServerPort:  getEnvAsInt("SERVER_PORT", 8080),
		DatabaseURL: getEnv("DATABASE_URL", "postgresql://user:password@localhost:5432/mydb"),
		// Add other configuration parameters here
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
