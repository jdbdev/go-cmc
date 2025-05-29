package main

import (
	"fmt"
	"log"

	"github.com/jdbdev/go-cmc/config"
	"github.com/jdbdev/go-cmc/db"
)

func main() {
	// Initialize config
	app := config.NewConfig()

	// Connect to database
	database, err := db.NewDatabase(app)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	fmt.Println("Database connection successful")
}
