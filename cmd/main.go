package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jdbdev/go-cmc/config"
	"github.com/jdbdev/go-cmc/db"
	"github.com/jdbdev/go-cmc/internal/mapper"
	"github.com/jdbdev/go-cmc/internal/ticker"
	"github.com/joho/godotenv"
)

func main() {

	//==========================================================================
	// Configuration & Setup
	//==========================================================================

	// Initialize config
	app := Init()

	// Initialize ID Mapper Service
	var IDMapSrvc mapper.IDMapInterface = mapper.NewIDMapService(app) // declare var as interface type and assign concrete implementation
	if err := IDMapSrvc.Initialize(); err != nil {
		log.Fatal("Failed to initialize ID map to fetch data from Coinmarketcap", err)
	}
	// Initialize ticker service
	tickerService := ticker.NewTickerService(app, IDMapSrvc)

	// Create connection to database
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

	// Check and/or Initialize DB tables
	// db.Initialize()

	//==========================================================================
	// Go Routine: Data Update Service
	//==========================================================================

	// Call CMC API every x seconds and update DB
	// Loop will continue even with errors
	updater := time.NewTicker(5 * time.Second)
	go func() {
		for range updater.C {
			fmt.Println("Updating CMC Data...")

			if err := tickerService.FetchAndDecodeData(); err != nil {
				log.Printf("Error fetching CMC data: %v", err)
				continue
			}

			// if err := tickerService.UpdateDB(); err != nil {
			// 	log.Printf("Error updating database: %v", err)
			// 	continue
			// }

			// fmt.Println("CMC Data Update Complete")
		}
	}()
	defer updater.Stop()

	//==========================================================================
	// Application Shutdown
	//==========================================================================

	// Wait for interrupt signal to gracefully shut down
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down gracefully...")
}

// Init initializes the application configuration and prints to stdout basic information
func Init() *config.AppConfig {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	app := config.NewAppConfig()
	PrintSettings(app)
	return app
}

// TEMP ONLY
func PrintSettings(app *config.AppConfig) {
	fmt.Printf("App in production: %v\n", app.AppCfg.InProduciton)
	fmt.Printf("Use DB: %v\n", app.AppCfg.UseDB)
	fmt.Printf("Base URL: %v\n", app.CMC.BaseURL)
}
