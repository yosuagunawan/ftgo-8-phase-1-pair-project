// internal/config/config.go
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Name     string
	User     string
	Password string
	Host     string
}

func LoadConfig() (*DBConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config := &DBConfig{
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
	}

	if config.Name == "" || config.User == "" || config.Password == "" || config.Host == "" {
		return nil, fmt.Errorf("all database configuration fields are required")
	}

	return config, nil
}
