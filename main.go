package main

import (
	"fmt"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"go-cqrs/cmd/command_handlers"
	"go-cqrs/interfaces/web"
	"go-cqrs/internal/app/customer"
	//"go-cqrs/internal/app/customer"
	"go-cqrs/internal/app/order"
	"go-cqrs/internal/domain"
	"go-cqrs/internal/infrastructure/db"
	"go-cqrs/internal/infrastructure/event_store"

	//"go-cqrs/internal/infrastructure/db"

	//"gorm.io/gorm"
	"log"
	"net/http"
)

var order1 = domain.Order{
	ID:          123,
	CustomerID:  "21313",
	Title:       "adwa",
	Description: "dwadwa",
	Price:       021,
}

//var (
//	db  *gorm.DB
//	err error
//)

func main() {
	database, err := db.SetupDatabase()
	if err != nil {
		log.Fatal("Unable to setup database:", err)
	}

	// Initialize dependencies
	//customerEventStore := event_store.NewEventStore("customer") // Assume NewEventStore is a function that returns an event store

	// Initialize the customer command handler and controller
	customerCommandHandler := command_handlers.NewCustomerCommandHandler( /*customerEventStore*/ database)
	customerController := customer.NewCustomerController(customerCommandHandler)

	// Initialize the order command handler and controller
	orderEventStore := event_store.NewEventStore("order")
	orderCommandHandler := command_handlers.NewOrderCommandHandler(orderEventStore)
	orderController := order.NewOrderController(orderCommandHandler)

	// Initialize the order and customer services with the database connection
	orderService := order.NewService(database)
	customerService := customer.NewService(database)

	// Configure routes
	router := web.SetupRouter(*orderService, *customerService, *customerController, *orderController)

	// Start the HTTP server
	fmt.Println("Server is running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

}
