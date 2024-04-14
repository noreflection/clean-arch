package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

const (
	host     = "localhost"
	port     = 5432
	user     = "your_username"
	password = "your_password"
	dbname   = "app_database"
)

func SetupDatabase() (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}
	fmt.Println("Database connection established successfully")

	var dbExists bool
	if err := db.QueryRow("SELECT EXISTS(SELECT datname FROM pg_database WHERE datname=$1)", dbname).Scan(&dbExists); err != nil {
		return nil, fmt.Errorf("failed to check database existence: %w", err)
	}

	return db, nil
}
