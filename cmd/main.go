package main

import (
	"fmt"

	"github.com/jdbdev/go-cmc/config"
)

func main() {
	// Initialize config
	var app = config.NewConfig()
	fmt.Println(app.APIKey)
	fmt.Println(app.DB.Host)

}
