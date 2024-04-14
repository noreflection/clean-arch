package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-cqrs/cmd/command_handlers"
	"go-cqrs/interfaces/web"
	"go-cqrs/internal/app/customer"
	"go-cqrs/internal/app/order"
	"go-cqrs/internal/app/order/repo"
	"go-cqrs/internal/infrastructure/db"
	"go-cqrs/internal/infrastructure/event_store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	database, err := db.SetupDatabase()
	if err != nil {
		log.Fatal("Unable to setup database:", err)
	}

	// Initialize repositories
	//customerRepo := customer.NewCustomerRepository(database)
	gormdb, _ := convertToGormDB(database)
	orderRepo := repo.NewOrderRepository(gormdb)

	// Initialize the customer command handler and controller
	customerEventStore := event_store.NewEventStore("customer")
	customerCommandHandler := command_handlers.NewCustomerCommandHandler(customerEventStore, database)
	customerController := customer.NewCustomerController(customerCommandHandler)

	// Initialize the order command handler and controller
	orderEventStore := event_store.NewEventStore("order")
	orderCommandHandler := command_handlers.NewOrderCommandHandler(orderEventStore, *orderRepo)
	orderController := order.NewOrderController(orderCommandHandler)

	router := web.SetupRouter(*customerController, *orderController)
	fmt.Println("Server is running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

}

// remove
func convertToGormDB(db *sql.DB) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}
