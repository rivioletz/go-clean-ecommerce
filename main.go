package main

import (
	"fmt"
	"log"

	"github.com/rivioletz/go-clean-ecommerce/config"
)

func main() {
	// Load configuration
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("DB Connection error: %v", err)
	}

	// Print the loaded configuration
	fmt.Println("✅ Connected to DB:", db)
}
