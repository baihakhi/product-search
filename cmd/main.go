package main

import (
	"log"
	"net/http"

	"github.com/baihakhi/product-search/internal/config"
	"github.com/baihakhi/product-search/internal/handler"
	"github.com/baihakhi/product-search/internal/repository"
	"github.com/baihakhi/product-search/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading .env file")
	}
	esConfig := config.LoadESConfig()
	esClient, err := config.InitESClient(esConfig)
	if err != nil {
		log.Fatalf("Error initializing Elasticsearch client: %v", err)
	}

	esRepo := repository.NewProductRepo(config.NewESClient(esClient))
	productService := service.NewProductService(esRepo)
	productHandler := handler.NewProductHandler(productService)
	http.HandleFunc("/search", productHandler.SearchProducts)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
