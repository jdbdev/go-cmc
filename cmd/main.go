package main

import (
	"fmt"
	"os"

	"github.com/jdbdev/go-cmc/types"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("test package 'main' build")
	// test package types integration:
	a := types.TestTypes("types")
	fmt.Printf("loading package: %s", a)

	// load variables from .env file
	godotenv.Load()
	// test access to .env:
	fmt.Println(".env variable for TEST is:", os.Getenv("ENV1"))
	
}