package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID          int
	CustomerID  string
	Title       string
	Description string
	Price       int

	// Add other fields specific to orders
}

// You can also include methods related to orders here, if needed.
