package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

// SetupDatabase configures database.
func SetupDatabase() (*sql.DB, error) {
	// Initialize the database connection
	connectionString := "postgres://vultradmin:AVNS_4ijKKcYd-4-mdo65XBT@vultr-prod-5f785376-9e78-4398-86ef-5bd59e46afa8-vultr-prod-5c15.vultrdb.com:16751/defaultdb"
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Failed to open database connection:", err)
	}

	// If Ping() succeeds, you have a valid database connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	fmt.Println("Database connection established successfully")

	dbname := "core-service-db"

	// Check if the database already exists
	var dbExists bool
	err = db.QueryRow("SELECT EXISTS(SELECT datname FROM pg_database WHERE datname=$1)", dbname).Scan(&dbExists)
	if err != nil {
		log.Fatal(err)
	}

	if !dbExists {
		// Database does not exist, so create it
		_, err := db.Exec("CREATE DATABASE " + dbname) // Capture the result and ignore it
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Database '%s' created.\n", dbname)
	} else {
		fmt.Printf("Database '%s' already exists.\n", dbname)
	}

	// Note: Do not close the database connection here.
	// The caller of this function should be responsible for closing it.

	return db, err
}
