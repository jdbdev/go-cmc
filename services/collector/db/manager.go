package db

import "sync"

var (
	instance *Database
	once     sync.Once
)

// SetDatabase sets the global database instance
func SetDatabase(db *Database) {
	once.Do(func() {
		instance = db
	})
}

// GetDatabase returns the global database instance
// Use db.GetDatabase() from other packages
func GetDatabase() *Database {
	return instance
}

// IsConnected returns true if the database is connected
func IsConnected() bool {
	return instance != nil && instance.db != nil
}
