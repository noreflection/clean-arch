package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"go-cqrs/internal/app"
	"go-cqrs/internal/infrastructure/db"
	"go-cqrs/internal/infrastructure/event_store"
	"go-cqrs/internal/infrastructure/repository"
	"go-cqrs/internal/interface/command_handlers"
	"go-cqrs/internal/interface/controller"
	"go-cqrs/internal/interface/query_handlers"
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
	orderService := app.NewOrderService(orderRepo)
	orderEventStore := event_store.NewEventStore("order")
	orderCommandHandler := command_handlers.NewOrderCommandHandler(orderEventStore, orderService)
	orderQueryHandler := query_handlers.NewOrderQueryHandler(orderRepo)
	orderController := controller.NewOrderController(orderCommandHandler, orderQueryHandler)

	customerRepo := repository.NewCustomerRepository(database)
	customerService := app.NewCustomerService(customerRepo)
	customerEventStore := event_store.NewEventStore("customer")
	customerCommandHandler := command_handlers.NewCustomerCommandHandler(customerEventStore, customerService)
	customerQueryHandler := query_handlers.NewCustomerQueryHandler(customerRepo)
	customerController := controller.NewCustomerController(customerCommandHandler, customerQueryHandler)

	router := web.SetupRouter(*customerController, *orderController)
	fmt.Println("Server is running on localhost:8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
