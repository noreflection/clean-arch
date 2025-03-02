package persistence

import (
	"database/sql"

	"go-cqrs/internal/application/ports"
	"go-cqrs/internal/infrastructure/config"
	db "go-cqrs/internal/infrastructure/persistence/database"
	repos "go-cqrs/internal/infrastructure/persistence/repositories"
)

// Database is the main database connection
type Database struct {
	*db.Database
}

// NewDatabase creates a new database connection
func NewDatabase(cfg *config.Config) (*Database, error) {
	database, err := db.NewDatabase(cfg)
	if err != nil {
		return nil, err
	}

	return &Database{Database: database}, nil
}

// NewCustomerRepository creates a new customer repository
func NewCustomerRepository(db *sql.DB) ports.CustomerRepository {
	return repos.NewCustomerRepository(db)
}

// NewOrderRepository creates a new order repository
func NewOrderRepository(db *sql.DB) ports.OrderRepository {
	return repos.NewOrderRepository(db)
}
