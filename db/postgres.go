package db

import (
	"database/sql"
	"fmt"

	"github.com/jdbdev/go-cmc/config"
	_ "github.com/lib/pq"
)

// DB is a global database connection pool
var DB *sql.DB

// Connect establishes a connection to the database using the provided configuration
func Connect(cfg *config.AppConfig) error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	// Verify connection
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("error pinging the database: %v", err)
	}

	return nil
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
