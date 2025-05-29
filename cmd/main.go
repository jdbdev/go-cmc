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
	if err := db.Connect(app); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Database connection successful")
}
