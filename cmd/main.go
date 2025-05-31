package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jdbdev/go-cmc/config"
	"github.com/jdbdev/go-cmc/db"
)

func main() {
	// Initialize config
	app := Init()

	// Connect to database
	if app.AppCfg.UseDB {
		database, err := db.NewDatabase(app)
		if err != nil {
			log.Fatal(err)
		}
		defer database.Close()

		// Set the global database instance
		db.SetDatabase(database)

		fmt.Println("Database connection successful")

		// Simple Query to test connection
		if db.IsConnected() {
			if err := database.GetDB().Ping(); err != nil {
				log.Printf("Database ping failed: %v", err)
			} else {
				log.Println("Database ping successful")
			}
		}
	}

	// Wait for interrupt signal to gracefully shut down
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down gracefully...")
}

// Init initializes the application configuration and prints to stdout basic information
func Init() *config.AppConfig {
	app := config.NewConfig()
	fmt.Printf("App in production: %v\n", app.AppCfg.InProduciton)
	fmt.Printf("Use DB: %v\n", app.AppCfg.UseDB)
	return app
}
