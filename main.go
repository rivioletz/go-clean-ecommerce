package main

import (
	"fmt"
	"log"

	"github.com/rivioletz/go-clean-ecommerce/config"
	"github.com/rivioletz/go-clean-ecommerce/product/repository/postgres"
	"github.com/rivioletz/go-clean-ecommerce/product/usecase"
)

func main() {
	// Load configuration
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("DB Connection error: %v", err)
	}

	// Print the loaded configuration
	fmt.Println("✅ Connected to DB:", db)

	repo := postgres.NewProductRepository(db)
	uc := usecase.NewProductUsecase(repo)
}
