package main

import (
	"cqrs-web-api/interfaces/web"
	"cqrs-web-api/internal/app/customer"
	"cqrs-web-api/internal/app/order"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"log"
	"net/http"
)

func main() {
	// Connect to PostgreSQL database
	db, err := sql.Open("postgres", "user=your_user dbname=your_db_name sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the order and customer services with the database connection
	orderService := order.NewService(db)
	customerService := customer.NewService(db)

	// Configure routes
	router := web.SetupRouter(*orderService, *customerService)

	// Start the HTTP server
	fmt.Println("Server is running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
