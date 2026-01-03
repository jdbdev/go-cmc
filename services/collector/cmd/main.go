package main

import (
	"context"
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
	"github.com/jdbdev/go-cmc/internal/coins"
	"github.com/jdbdev/go-cmc/internal/mapper"
	"github.com/jdbdev/go-cmc/internal/ticker"
	"github.com/joho/godotenv"
)

// Collector service (go-cmc)requires two services to run: internal/mapper and internal/ticker.
// Mapper service generates an ID map based on coin lookups (symbols) using Coinmarketcap API for ID mapping.
// ticker service fetches up to date data for each token/coin in the DB.
// Services run concurrently at set intervals found in config/config.go file. Adjust based on API credit expenditure.
// Services update the database with up to date data.
// All configuration settings are stored in .env and loaded by config/config.go file.

func main() {

	//==========================================================================
	// Configuration & Initialization/Setup
	//==========================================================================

	// Initialize Logger first (required for Init() and rest of app)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.Info("CMC API application starting - Version 0.1")

	// Initialize applicaiton configuration
	app := Init(logger)

	// Initialize http client
	var client = &http.Client{}

	// Initialize Mapper Service (with dependency injection of app, logger and client)
	var mapperService mapper.IDMapInterface = mapper.NewIDMapService(app, logger, client)

	// Initialize ticker service (with dependency injection of app, logger and client)
	var tickerService ticker.TickerInterface = ticker.NewTickerService(app, mapperService, logger)

	// Initialize coins service (with dependency injection of logger)
	var coinService coins.CoinInterface = coins.NewCoinService(logger)

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
	// Service Calls with Context
	//==========================================================================

	// mapperService calls with context timeout
	ctx, cancel := context.WithTimeout(context.Background(), app.CMC.RequestTimeout)
	defer cancel()

	initialCoins, err := mapperService.GetCMCTopCoins(ctx, 2)
	if err != nil {
		logger.Error("Failed getting topcoins", "error", err)
	} else {
		fmt.Println(string(initialCoins))
	}

	// tickerService calls with context timeout
	// coinService calls with context timeout
	coinService.InitializeCoinTable()

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
func Init(logger *slog.Logger) *config.AppConfig {
	// Load .env file from root directory (monorepo structure)
	if err := godotenv.Load("../../.env"); err != nil {
		logger.Warn("Error loading .env file", "error", err)
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
	fmt.Printf("Request Timeout: %v\n", app.CMC.RequestTimeout)
}
