package config

import (
	"net/http"
	"os"
)

// AppConfig holds all configuration settings for the application
type AppConfig struct {
	DB     DBConfig
	APIKey string
	AppCfg AppSettings
	Srv    *http.Server
}

// AppSettings holds general application settings
type AppSettings struct {
	InProduciton bool
}

// DBConfig holds database  configuration settings
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// NewConfig creates and returns a new AppConfig instance
func NewConfig() *AppConfig {
	return &AppConfig{
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "postgres"),
		},
		APIKey: getEnv("API_KEY", "123"),
		AppCfg: AppSettings{
			InProduciton: getEnv("IN_PRODUCTION", "false") == "true",
		},
	}
}

// getEnv() function to get env variables from .env file
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
