package config

import (
	"net/http"
	"os"
	"time"
)

// AppConfig holds all configuration settings for the application
type AppConfig struct {
	DB       DBSettings
	CMC      CMCSettings
	AppCfg   AppSettings
	Srv      *http.Server
	Interval IntervalSettings
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
	APIKey         string
	BaseURL        string
	QuotesURL      string
	IDMapURL       string
	RequestTimeout time.Duration
}

// IntervalSettings holds the time settings in seconds for the ticker and mapper services
type IntervalSettings struct {
	TickerInterval time.Duration
	MapperInterval time.Duration
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
			APIKey:         getEnv("CMC_API_KEY", "123"),
			BaseURL:        getEnv("CMC_BASE_URL", ""),
			QuotesURL:      getEnv("CMC_QUOTES_URL", ""),
			IDMapURL:       getEnv("CMC_ID_MAP_URL", ""),
			RequestTimeout: getEnvAsDuration("CMC_REQUEST_TIMEOUT", "30s"),
		},

		AppCfg: AppSettings{
			InProduciton: getEnv("IN_PRODUCTION", "false") == "true",
			UseDB:        getEnv("USE_DB", "false") == "true",
		},

		Interval: IntervalSettings{
			TickerInterval: getEnvAsDuration("TICKER_INTERVAL", "2m"),
			MapperInterval: getEnvAsDuration("MAPPER_INTERVAL", "24h"),
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

// getEnvAsDuration() function to get env variables as duration from .env file
func getEnvAsDuration(key, defaultValue string) time.Duration {
	value := getEnv(key, defaultValue)
	duration, err := time.ParseDuration(value)
	if err != nil {
		// If parsing fails, return default
		defaultDuration, _ := time.ParseDuration(defaultValue)
		return defaultDuration
	}
	return duration
}
