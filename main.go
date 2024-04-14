package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-cqrs/cmd/command_handlers"
	"go-cqrs/cmd/query_handlers"
	"go-cqrs/interfaces/web"
	"go-cqrs/internal/app/customer"
	"go-cqrs/internal/app/order"
	"go-cqrs/internal/app/order/repo"
	"go-cqrs/internal/infrastructure/db"
	"go-cqrs/internal/infrastructure/event_store"
	"log"
	"net/http"
)

func main() {
	database, err := db.SetupDatabase()
	if err != nil {
		log.Printf("Unable to setup database: %v", err)
		return
	}
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			log.Printf("Unable to close the connection to databaseL %v", err)
		}
	}(database)

	// Initialize repositories
	//customerRepo := customer.NewCustomerRepository(database)
	orderRepo := repo.NewOrderRepository(database)

	// Initialize the customer command handler and controller
	customerEventStore := event_store.NewEventStore("customer") //todo: impl eventstore
	customerCommandHandler := command_handlers.NewCustomerCommandHandler(customerEventStore, database)
	customerController := customer.NewCustomerController(customerCommandHandler)

	// Initialize the order command handler and controller
	orderEventStore := event_store.NewEventStore("order")
	orderCommandHandler := command_handlers.NewOrderCommandHandler(orderEventStore, *orderRepo)
	orderQueryHandler := query_handlers.NewOrderQueryHandler(*orderRepo)
	orderController := order.NewOrderController(orderCommandHandler, orderQueryHandler)

	router := web.SetupRouter(*customerController, *orderController)
	fmt.Println("Server is running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
