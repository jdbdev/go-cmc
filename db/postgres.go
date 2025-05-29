package db

import (
	"database/sql"
	"fmt"

	"github.com/jdbdev/go-cmc/config"
	_ "github.com/lib/pq"
)

// Database represents a database connection and its methods
type Database struct {
	db  *sql.DB
	cfg *config.AppConfig
}

// NewDatabase creates and returns a new Database instance
func NewDatabase(cfg *config.AppConfig) (*Database, error) {
	database := &Database{
		cfg: cfg,
	}

	if err := database.connect(); err != nil {
		return nil, err
	}

	return database, nil
}

// connect establishes a connection to the database using the stored configuration
func (d *Database) connect() error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		d.cfg.DB.Host, d.cfg.DB.Port, d.cfg.DB.User, d.cfg.DB.Password, d.cfg.DB.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		db.Close() // Clean up before returning error
		return fmt.Errorf("error pinging the database: %v", err)
	}

	d.db = db
	return nil
}

// Close closes the database connection
func (d *Database) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}

// GetDB returns the underlying *sql.DB instance
func (d *Database) GetDB() *sql.DB {
	return d.db
}
