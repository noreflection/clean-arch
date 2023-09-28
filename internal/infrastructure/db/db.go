package db

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	db  *gorm.DB
	err error
)

// SetupDatabase configures database.
func SetupDatabase() (*sql.DB, error) {
	// Initialize the database connection
	//dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	//dsn := "host='vultr-prod-5f785376-9e78-4398-86ef-5bd59e46afa8-vultr-prod-5c15.vultrdb.com' user=vultradmin password='AVNS_4ijKKcYd-4-mdo65XBT' dbname=defaultdb port=16751 sslmode=disable"
	connectionString := "postgres://vultradmin:AVNS_4ijKKcYd-4-mdo65XBT@vultr-prod-5f785376-9e78-4398-86ef-5bd59e46afa8-vultr-prod-5c15.vultrdb.com:16751/defaultdb"
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open database connection:", err)
	}

	// Ping the database to check connectivity
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get the underlying SQL database", err)
	}

	// If Ping() succeeds, you have a valid database connection
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	fmt.Println("Database connection established successfully")

	print("here")
	defer sqlDB.Close()

	dbname := "core-service-db"
	// Check if the database already exists
	var dbExists bool
	err = sqlDB.QueryRow("SELECT EXISTS(SELECT datname FROM pg_database WHERE datname=$1)", dbname).Scan(&dbExists)
	if err != nil {
		log.Fatal(err)
	}

	if !dbExists {
		// Database does not exist, so create it
		err := db.Exec("CREATE DATABASE " + pq.QuoteIdentifier(dbname))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Database '%s' created.\n", dbname)

	} else {
		fmt.Printf("Database '%s' already exists.\n", dbname)
	}

	// create table
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS my_table (
            id serial PRIMARY KEY,
            name VARCHAR (255) NOT NULL,
            age INT
        );
    `

	// Execute the SQL statement to create the table
	err1 := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("shitttttttttttttttt", err1)
	}

	fmt.Println("Table 'my_table' created successfully.")

	return sqlDB, err
}

func createTable() error {
	// Create the SQL statement to create the table
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS my_table (
            id serial PRIMARY KEY,
            name VARCHAR (255) NOT NULL,
            age INT
        );
    `

	// Execute the SQL statement to create the table
	err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("shitttttttttttttttt", err)
	}

	fmt.Println("Table 'my_table' created successfully.")
	return nil
}
