package container

import (
	"log"

	"go-cqrs/internal/app"
	"go-cqrs/internal/app/usecases"
	"go-cqrs/internal/infra/config"
	"go-cqrs/internal/infra/db"
	"go-cqrs/internal/infra/event_store"
	"go-cqrs/internal/infra/repository"
	"go-cqrs/internal/interface/command_handlers"
	"go-cqrs/internal/interface/controller"
	"go-cqrs/internal/interface/query_handlers"
	"go-cqrs/internal/interface/web"
)

// Container holds all application dependencies
type Container struct {
	Config *config.Config
	DB     *db.Database

	// Repositories
	OrderRepository    app.OrderRepository
	CustomerRepository app.CustomerRepository

	// Use Cases
	OrderUseCase    app.OrderUseCase
	CustomerUseCase app.CustomerUseCase

	// Event Stores
	OrderEventStore    event_store.EventStore
	CustomerEventStore event_store.EventStore

	// Command Handlers
	OrderCommandHandler    *command_handlers.OrderCommandHandler
	CustomerCommandHandler *command_handlers.CustomerCommandHandler

	// Query Handlers
	OrderQueryHandler    *query_handlers.OrderQueryHandler
	CustomerQueryHandler *query_handlers.CustomerQueryHandler

	// Controllers
	OrderController    *controller.OrderController
	CustomerController *controller.CustomerController

	// Router
	Router web.Router
}

// NewContainer creates a new dependency injection container
func NewContainer() (*Container, error) {
	container := &Container{}

	// Load configuration
	var err error
	container.Config, err = config.Load()
	if err != nil {
		return nil, err
	}

	// Setup database
	container.DB, err = db.NewDatabase(container.Config)
	if err != nil {
		return nil, err
	}

	// Initialize tables
	if err = container.DB.SetupDatabaseTables(); err != nil {
		return nil, err
	}

	// Initialize repositories
	container.OrderRepository = repository.NewOrderRepository(container.DB.DB)
	container.CustomerRepository = repository.NewCustomerRepository(container.DB.DB)

	// Initialize use cases
	container.OrderUseCase = usecases.NewOrderUseCases(
		container.OrderRepository.(app.OrderRepository),
		container.CustomerRepository.(app.CustomerRepository),
	)
	container.CustomerUseCase = usecases.NewCustomerUseCases(
		container.CustomerRepository.(app.CustomerRepository),
	)

	// Initialize event stores
	container.OrderEventStore = event_store.NewEventStore("order")
	container.CustomerEventStore = event_store.NewEventStore("customer")

	// Initialize command handlers
	container.OrderCommandHandler = command_handlers.NewOrderCommandHandler(
		container.OrderEventStore,
		container.OrderUseCase.(app.OrderUseCase),
	)
	container.CustomerCommandHandler = command_handlers.NewCustomerCommandHandler(
		container.CustomerEventStore,
		container.CustomerUseCase.(app.CustomerUseCase),
	)

	// Initialize query handlers
	container.OrderQueryHandler = query_handlers.NewOrderQueryHandler(
		container.OrderRepository.(app.OrderRepository),
	)
	container.CustomerQueryHandler = query_handlers.NewCustomerQueryHandler(
		container.CustomerRepository.(app.CustomerRepository),
	)

	// Initialize controllers
	container.OrderController = controller.NewOrderController(
		container.OrderCommandHandler,
		container.OrderQueryHandler,
	)
	container.CustomerController = controller.NewCustomerController(
		container.CustomerCommandHandler,
		container.CustomerQueryHandler,
	)

	// Initialize router
	container.Router = web.NewRouter(
		*container.CustomerController,
		*container.OrderController,
	)

	return container, nil
}

// Close closes all resources
func (c *Container) Close() {
	if c.DB != nil {
		if err := c.DB.Close(); err != nil {
			log.Printf("error closing database: %v", err)
		}
	}
}
