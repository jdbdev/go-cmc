package main

import (
	"fmt"

	"github.com/jdbdev/go-cmc/types"
)

func main() {
	fmt.Println("test package main build")
	a := types.TestTypes("types")
	fmt.Printf("loading package: %s", a)
}