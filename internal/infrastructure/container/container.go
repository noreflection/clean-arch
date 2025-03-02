package container

import (
	"go-cqrs/internal/adapters/cqrs/commands"
	"go-cqrs/internal/adapters/cqrs/queries"
	"go-cqrs/internal/adapters/http/controllers"
	"go-cqrs/internal/adapters/http/router"
	"go-cqrs/internal/application/ports"
	"go-cqrs/internal/application/services"
	"go-cqrs/internal/infrastructure/config"
	"go-cqrs/internal/infrastructure/database"
	"go-cqrs/internal/infrastructure/logger"
	event_store "go-cqrs/internal/infrastructure/messaging/events"
	"go-cqrs/internal/infrastructure/repositories"
)

// Container holds all application dependencies
type Container struct {
	Config *config.Config
	DB     *database.Database
	Logger logger.Logger

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
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// Initialize logger
	log := logger.NewZapLogger(logger.LogLevel(cfg.LogLevel), cfg.Environment == "production")
	log.Info("Initializing application container",
		logger.String("environment", cfg.Environment),
		logger.String("server_address", cfg.ServerAddress()))

	// Create container
	c := &Container{
		Config: cfg,
		Logger: log,
	}

	// Initialize database
	db, err := database.NewDatabase(cfg.DatabaseURL())
	if err != nil {
		log.Error("Failed to connect to database", logger.Error(err))
		return nil, err
	}
	c.DB = db
	log.Info("Connected to database", logger.String("host", cfg.DBHost), logger.String("database", cfg.DBName))

	// Initialize tables
	if err = c.DB.SetupDatabaseTables(); err != nil {
		return nil, err
	}

	// Initialize repositories
	c.OrderRepository = repositories.NewOrderRepository(c.DB.DB)
	c.CustomerRepository = repositories.NewCustomerRepository(c.DB.DB)

	// Initialize use cases
	c.OrderUseCase = services.NewOrderService(
		c.OrderRepository,
		c.CustomerRepository,
	)
	c.CustomerUseCase = services.NewCustomerService(
		c.CustomerRepository,
	)

	// Initialize event stores
	c.OrderEventStore = event_store.NewPostgresEventStore(c.DB.DB, "order", c.Logger)
	c.CustomerEventStore = event_store.NewPostgresEventStore(c.DB.DB, "customer", c.Logger)

	// Initialize command handlers
	c.OrderCommandHandler = commands.NewOrderCommandHandler(
		c.OrderEventStore,
		c.OrderUseCase,
	)
	c.CustomerCommandHandler = commands.NewCustomerCommandHandler(
		c.CustomerEventStore,
		c.CustomerUseCase,
	)

	// Initialize query handlers
	c.OrderQueryHandler = queries.NewOrderQueryHandler(
		c.OrderRepository,
	)
	c.CustomerQueryHandler = queries.NewCustomerQueryHandler(
		c.CustomerRepository,
	)

	// Initialize controllers
	c.OrderController = *controllers.NewOrderController(
		c.OrderCommandHandler,
		c.OrderQueryHandler,
	)
	c.CustomerController = *controllers.NewCustomerController(
		c.CustomerCommandHandler,
		c.CustomerQueryHandler,
	)

	// Initialize router
	c.Router = router.NewRouter(
		c.CustomerController,
		c.OrderController,
	)

	return c, nil
}

// Close closes all resources
func (c *Container) Close() {
	if c.DB != nil {
		if err := c.DB.Close(); err != nil {
			c.Logger.Error("error closing database", logger.Error(err))
		}
	}
}
