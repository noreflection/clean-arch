package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"go-cqrs/cmd/command_handlers"
	"go-cqrs/cmd/query_handlers"
	"go-cqrs/interfaces/web"
	"go-cqrs/internal/app/customer"
	"go-cqrs/internal/app/order"
	"go-cqrs/internal/infrastructure/db"
	"go-cqrs/internal/infrastructure/event_store"
	"log"
	"net/http"
)

func main() {
	database, err := db.SetupDatabase()
	if err != nil {
		log.Printf("unable to setup database: %v", err)
		return
	}
	defer database.Close()

	orderRepo := order.NewRepository(database)
	orderService := order.NewService(orderRepo)
	orderEventStore := event_store.NewEventStore("order")
	orderCommandHandler := command_handlers.NewOrderCommandHandler(orderEventStore, orderService)
	orderQueryHandler := query_handlers.NewOrderQueryHandler(orderRepo)
	orderController := order.NewOrderController(orderCommandHandler, orderQueryHandler)

	customerRepo := customer.NewRepository(database)
	customerService := customer.NewService(customerRepo)
	customerEventStore := event_store.NewEventStore("customer")
	customerCommandHandler := command_handlers.NewCustomerCommandHandler(customerEventStore, customerService)
	customerQueryHandler := query_handlers.NewCustomerQueryHandler(customerRepo)
	customerController := customer.NewCustomerController(customerCommandHandler, customerQueryHandler)

	router := web.SetupRouter(*customerController, *orderController)
	fmt.Println("Server is running on localhost:8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
