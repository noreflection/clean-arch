package main

import (
	"cqrs-web-api/config"
	"cqrs-web-api/infrastructure/database"
	"cqrs-web-api/interfaces/web"
	"cqrs-web-api/internal/app/customer"
	"cqrs-web-api/internal/app/order"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the database connection
	db, err := database.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize the order and customer services
	orderService := order.NewService(db)
	customerService := customer.NewService(db)

	// Initialize the web server
	router := web.SetupRouter(orderService, customerService)
	port := cfg.ServerPort
	serverAddr := fmt.Sprintf(":%d", port)

	// Start the web server
	fmt.Printf("Server is running on port %d...\n", port)
	err = http.ListenAndServe(serverAddr, router)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
