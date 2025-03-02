package db

import (
	"database/sql"
	"errors"
	"fmt"
	"go-cqrs/internal/infra/config"
	
	_ "github.com/lib/pq"
)

// Database represents a database connection
type Database struct {
	*sql.DB
}

// NewDatabase creates a new database connection
func NewDatabase(cfg *config.Config) (*Database, error) {
	// Open database connection
	db, err := sql.Open("postgres", cfg.DatabaseURL())
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}
	
	// Verify connection is working
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	
	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	
	return &Database{db}, nil
}

// SetupDatabaseTables creates database tables if they don't exist
func (db *Database) SetupDatabaseTables() error {
	// Create customers table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS customers (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return errors.New("failed to create customers table: " + err.Error())
	}
	
	// Create orders table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			customer_id INTEGER REFERENCES customers(id) ON DELETE SET NULL,
			product TEXT NOT NULL,
			quantity INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return errors.New("failed to create orders table: " + err.Error())
	}
	
	return nil
}

// Close closes the database connection
func (db *Database) Close() error {
	return db.DB.Close()
} 