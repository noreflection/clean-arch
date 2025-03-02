package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Database represents a database connection
type Database struct {
	*sql.DB
}

// NewDatabase creates a new database connection
func NewDatabase(connString string) (*Database, error) {
	// Open database connection
	db, err := sql.Open("postgres", connString)
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
		return fmt.Errorf("failed to create customers table: %w", err)
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
		return fmt.Errorf("failed to create orders table: %w", err)
	}

	// Create events table for event sourcing
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS events (
			id SERIAL PRIMARY KEY,
			event_type TEXT NOT NULL,
			occurred_at TIMESTAMP NOT NULL,
			event_data JSONB NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create events table: %w", err)
	}

	return nil
}

// WithTransaction executes function within a database transaction
func (db *Database) WithTransaction(fn func(*sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Re-throw panic after rollback
		} else if err != nil {
			tx.Rollback() // err is not nil; rollback
		} else {
			err = tx.Commit() // err is nil; commit
		}
	}()

	err = fn(tx)
	return err
}

// Close closes the database connection
func (db *Database) Close() error {
	return db.DB.Close()
}
