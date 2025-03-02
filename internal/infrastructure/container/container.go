package container

import (
	"log"

	"go-cqrs/internal/adapters/cqrs/commands"
	"go-cqrs/internal/adapters/cqrs/queries"
	"go-cqrs/internal/adapters/http/controllers"
	"go-cqrs/internal/adapters/http/router"
	"go-cqrs/internal/application/ports"
	"go-cqrs/internal/application/services"
	"go-cqrs/internal/infrastructure/config"
	event_store "go-cqrs/internal/infrastructure/messaging/events"
	"go-cqrs/internal/infrastructure/persistence"
)

// Container holds all application dependencies
type Container struct {
	Config *config.Config
	DB     *persistence.Database

	// Repositories
	OrderRepository    ports.OrderRepository
	CustomerRepository ports.CustomerRepository

	// Use Cases
	OrderUseCase    ports.OrderUseCase
	CustomerUseCase ports.CustomerUseCase

	// Event Stores
	OrderEventStore    event_store.EventStore
	CustomerEventStore event_store.EventStore

	// Command Handlers
	OrderCommandHandler    *commands.OrderCommandHandler
	CustomerCommandHandler *commands.CustomerCommandHandler

	// Query Handlers
	OrderQueryHandler    *queries.OrderQueryHandler
	CustomerQueryHandler *queries.CustomerQueryHandler

	// Controllers
	OrderController    controllers.OrderController
	CustomerController controllers.CustomerController

	// Router
	Router router.Router
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
	container.DB, err = persistence.NewDatabase(container.Config)
	if err != nil {
		return nil, err
	}

	// Initialize tables
	if err = container.DB.SetupDatabaseTables(); err != nil {
		return nil, err
	}

	// Initialize repositories
	container.OrderRepository = persistence.NewOrderRepository(container.DB.DB)
	container.CustomerRepository = persistence.NewCustomerRepository(container.DB.DB)

	// Initialize use cases
	container.OrderUseCase = services.NewOrderService(
		container.OrderRepository,
		container.CustomerRepository,
	)
	container.CustomerUseCase = services.NewCustomerService(
		container.CustomerRepository,
	)

	// Initialize event stores
	container.OrderEventStore = event_store.NewEventStore("order")
	container.CustomerEventStore = event_store.NewEventStore("customer")

	// Initialize command handlers
	container.OrderCommandHandler = commands.NewOrderCommandHandler(
		container.OrderEventStore,
		container.OrderUseCase,
	)
	container.CustomerCommandHandler = commands.NewCustomerCommandHandler(
		container.CustomerEventStore,
		container.CustomerUseCase,
	)

	// Initialize query handlers
	container.OrderQueryHandler = queries.NewOrderQueryHandler(
		container.OrderRepository,
	)
	container.CustomerQueryHandler = queries.NewCustomerQueryHandler(
		container.CustomerRepository,
	)

	// Initialize controllers
	container.OrderController = *controllers.NewOrderController(
		container.OrderCommandHandler,
		container.OrderQueryHandler,
	)
	container.CustomerController = *controllers.NewCustomerController(
		container.CustomerCommandHandler,
		container.CustomerQueryHandler,
	)

	// Initialize router
	container.Router = router.NewRouter(
		container.CustomerController,
		container.OrderController,
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
