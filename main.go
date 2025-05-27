package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rivioletz/go-clean-ecommerce/config"
	"github.com/rivioletz/go-clean-ecommerce/product/handler"
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
	uc := usecase.NewProductUseCase(repo)

	r := gin.Default()
	productHandler := handler.NewProductHandler(uc)
	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products/:id", productHandler.GetProductByID)

	r.Run(":8080")
	log.Println("Server running on port 8080")
}
