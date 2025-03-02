package persistence

import (
	"database/sql"

	"go-cqrs/internal/application/ports"
	db "go-cqrs/internal/infrastructure/persistence/database"
	repos "go-cqrs/internal/infrastructure/persistence/repositories"
)

// Database is the main database connection
type Database struct {
	*db.Database
}

// NewDatabase creates a new database connection
func NewDatabase(connString string) (*Database, error) {
	database, err := db.NewDatabase(connString)
	if err != nil {
		return nil, err
	}

	return &Database{Database: database}, nil
}

// WithTransaction executes function within a database transaction
func (d *Database) WithTransaction(fn func(*sql.Tx) error) error {
	return d.Database.WithTransaction(fn)
}

// NewCustomerRepository creates a new customer repository
func NewCustomerRepository(db *sql.DB) ports.CustomerRepository {
	return repos.NewCustomerRepository(db)
}

// NewOrderRepository creates a new order repository
func NewOrderRepository(db *sql.DB) ports.OrderRepository {
	return repos.NewOrderRepository(db)
}
