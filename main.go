package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"go-cqrs/internal/app/customer"
	"go-cqrs/internal/app/order"
	"go-cqrs/internal/infrastructure/db"
	"go-cqrs/internal/infrastructure/event_store"
	"go-cqrs/internal/infrastructure/repository"
	command_handlers2 "go-cqrs/internal/interface/command_handlers"
	"go-cqrs/internal/interface/controller"
	query_handlers2 "go-cqrs/internal/interface/query_handlers"
	"go-cqrs/internal/interface/web"
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

	orderRepo := repository.NewOrderRepository(database)
	orderService := order.NewService(orderRepo)
	orderEventStore := event_store.NewEventStore("order")
	orderCommandHandler := command_handlers2.NewOrderCommandHandler(orderEventStore, orderService)
	orderQueryHandler := query_handlers2.NewOrderQueryHandler(orderRepo)
	orderController := controller.NewOrderController(orderCommandHandler, orderQueryHandler)

	customerRepo := repository.NewCustomerRepository(database)
	customerService := customer.NewService(customerRepo)
	customerEventStore := event_store.NewEventStore("customer")
	customerCommandHandler := command_handlers2.NewCustomerCommandHandler(customerEventStore, customerService)
	customerQueryHandler := query_handlers2.NewCustomerQueryHandler(customerRepo)
	customerController := controller.NewCustomerController(customerCommandHandler, customerQueryHandler)

	router := web.SetupRouter(*customerController, *orderController)
	fmt.Println("Server is running on localhost:8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
