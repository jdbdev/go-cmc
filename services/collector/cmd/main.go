package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
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

// collector service (go-cmc)requires two services to run: internal/mapper and internal/ticker.
// mapper service generates an ID map based on coin lookups (symbols) using Coinmarketcap API for ID mapping.
// ticker service fetches up to date data for each token/coin in the DB.
// both services run concurrently at set intervals found in config/config.go file. Adjust based on API credit expenditure.
// both services update the database with up to date data.

func main() {

	//==========================================================================
	// Configuration & Initialization/Setup
	//==========================================================================

	// Initialize config
	app := Init()

	// Initialize Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.Info("CMC API application starting - Version 0.1")

	// Initialize http client
	var client = &http.Client{}

	// Initialize Mapper Service (with dependency injection of app, logger and client)
	var IDMapSrvc mapper.IDMapInterface = mapper.NewIDMapService(app, logger, client)

	// Initialize ticker service (with dependency injection of app, logger and client)
	tickerService := ticker.NewTickerService(app, IDMapSrvc, logger)

	//==========================================================================
	// Database Setup
	//==========================================================================

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

	//==========================================================================
	// Go Routine: Data Update Service
	//==========================================================================

	// Call CMC API every x seconds and update DB
	// Loop will continue even with errors
	timeInterval := app.Interval.TickerInterval
	updater := time.NewTicker(timeInterval)
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
