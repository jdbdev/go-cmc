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

// Collector service (go-cmc)requires three services to run: internal/mapper, internal/coins and internal/ticker.
// Mapper service generates an ID map based on coin lookups (symbols) using Coinmarketcap API for ID mapping.
// ticker service fetches up to date data for each token/coin in the DB.
// Services run concurrently at set intervals found in config/config.go file. Adjust based on API credit expenditure.
// Services update the database with up to date data.
// All configuration settings are stored in .env and loaded by config/config.go file.

// Services holds the interfaces for the mapper, ticker and coins services.
type Services struct {
	Mapper mapper.IDMapInterface
	Ticker ticker.TickerInterface
	Coins  coins.CoinInterface
}

func main() {

	//==========================================================================
	// Configuration & Initialization/Setup
	//==========================================================================

	// Initialize Logger first (required for Init() and rest of app)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.Info("CMC API application starting - Version 0.1")
	// Initialize applicaiton configuration
	app := InitConfig(logger)
	// Initialize http client
	client := &http.Client{}
	// Initialize services Mapper, Ticker and Coins. Inject dependencies required.
	services := InitServices(app, logger, client)

	//==========================================================================
	// Database Setup
	//==========================================================================

	// Create connection to database
	database, err := InitDatabase(app, logger)
	if err != nil {
		logger.Error("failed to initialize database", "error", err)
	}
	if database != nil {
		defer database.Close()
		logger.Info("Database connection successful")
	}

	//==========================================================================
	// Service Calls with Context
	//==========================================================================

	// mapperService calls with context timeout
	ctx, cancel := context.WithTimeout(context.Background(), app.CMC.RequestTimeout)
	defer cancel()

	initialCoins, err := services.Mapper.GetCMCTopCoins(ctx, 2)
	if err != nil {
		logger.Error("Failed getting topcoins", "error", err)
	} else {
		fmt.Println(string(initialCoins))
	}

	// tickerService calls with context timeout
	// coinService calls with context timeout

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

			if err := services.Ticker.FetchAndDecodeData(); err != nil {
				log.Printf("Error fetching CMC data: %v", err)
				continue
			}

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
func InitConfig(logger *slog.Logger) *config.AppConfig {
	// Load .env file from root directory (monorepo structure)
	if err := godotenv.Load("../../.env"); err != nil {
		logger.Warn("Error loading .env file", "error", err)
	}
	app := config.NewAppConfig()
	PrintSettings(app)
	return app
}

func InitServices(app *config.AppConfig, logger *slog.Logger, client *http.Client) *Services {
	mapperService := mapper.NewIDMapService(app, logger, client)
	tickerService := ticker.NewTickerService(app, mapperService, logger)
	coinService := coins.NewCoinService(logger)

	return &Services{
		Mapper: mapperService,
		Ticker: tickerService,
		Coins:  coinService,
	}
}

func InitDatabase(app *config.AppConfig, logger *slog.Logger) (*db.Database, error) {
	if !app.AppCfg.UseDB {
		logger.Info("Database disabled in settings - not in use")
		return nil, nil
	}
	database, err := db.NewDatabase(app)
	if err != nil {
		log.Fatal(err)
	}
	db.SetDatabase(database)
	return database, nil
}

// TEMP HELPERS ONLY. REMOVE WHEN DONE.
func PrintSettings(app *config.AppConfig) {
	fmt.Printf("App in production: %v\n", app.AppCfg.InProduciton)
	fmt.Printf("Use DB: %v\n", app.AppCfg.UseDB)
	fmt.Printf("Base URL: %v\n", app.CMC.BaseURL)
	fmt.Printf("Request Timeout: %v\n", app.CMC.RequestTimeout)
}
