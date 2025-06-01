package config

import (
	"net/http"
	"os"
)

// AppConfig holds all configuration settings for the application
type AppConfig struct {
	DB     DBSettings
	CMC    CMCSettings
	AppCfg AppSettings
	Srv    *http.Server
}

// AppCofig holds general application settings
type AppSettings struct {
	InProduciton bool
	UseDB        bool
}

// DBConfig holds database configuration settings
type DBSettings struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// CMCCOnfig holds Coinmarketcap API configuration
type CMCSettings struct {
	APIKey  string
	BaseURL string
}

// NewConfig creates and returns a new AppConfig instance
func NewAppConfig() *AppConfig {
	return &AppConfig{
		DB: DBSettings{
			Host:     getEnv("DB_HOST", "postgres"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "postgres"),
		},
		CMC: CMCSettings{
			APIKey:  getEnv("CMC_API_KEY", "123"),
			BaseURL: getEnv("CMC_BASE_URL", ""),
		},

		AppCfg: AppSettings{
			InProduciton: getEnv("IN_PRODUCTION", "false") == "true",
			UseDB:        getEnv("USE_DB", "false") == "true",
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
