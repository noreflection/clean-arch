package main

import (
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"go-cqrs/interfaces/web"
	"go-cqrs/internal/app/customer"
	"go-cqrs/internal/app/order"
	"go-cqrs/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var order1 = domain.Order{
	ID:          123,
	CustomerID:  "21313",
	Title:       "adwa",
	Description: "dwadwa",
	Price:       021,
}

var (
	db  *gorm.DB
	err error
)

func main() {
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
	fmt.Println("Table created successfully")

	//
	// Auto-migrate to create tables for the models
	//db.AutoMigrate(&domain.Order{}, &domain.Customer{})

	//// Create a gormigrate instance using the new "gorm.io/gorm" DB
	//m := gormigrate.New(db, &gormigrate.Options{
	//	TableName:      "migrations",
	//	IDColumnName:   "id",
	//	IDColumnSize:   255,
	//	UseTransaction: true,
	//}, []*gormigrate.Migration{
	//	{
	//		ID: "202209201200", // Unique ID for the migration
	//		Migrate: func(tx *gorm.DB) error {
	//			// Your migration logic here
	//			return nil
	//		},
	//		Rollback: func(tx *gorm.DB) error {
	//			// Your rollback logic here
	//			return nil
	//		},
	//	},
	//})
	//
	//// Apply migrations
	//if err := m.Migrate(); err != nil {
	//	log.Fatalf("Could not migrate: %v", err)
	//}

	log.Println("Migration completed")

	log.Println("Migration completed")

	// Initialize the order and customer services with the database connection
	orderService := order.NewService(sqlDB)
	customerService := customer.NewService(sqlDB)

	// Configure routes
	router := web.SetupRouter(*orderService, *customerService)

	// Start the HTTP server
	fmt.Println("Server is running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

}
