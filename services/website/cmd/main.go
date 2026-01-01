package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// Moonramp web application queries the DB to gather up to date data for coins in portfolios.
// Coin/Token data is not updated by the web application.
// Rellies on collector service to update DB with up to date data from Coincmarketcap API.

func main() {

	//==========================================================================
	// Configuration & Initialization/Setup
	//==========================================================================

	// Initialize and configer server:
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      nil, //replace with real handler later
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Initialize logger:
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.Info("Moonramp application starting - Version 0.1")

	//==========================================================================
	// Web Server Setup and Run
	//==========================================================================

	slog.Info("Server starting on port 8080")
	http.HandleFunc("/", homeHandler)

	srv.ListenAndServe()
}

// homHandler handles the home page request - TEMP ONLY.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "staging for moonbags web application")
}
