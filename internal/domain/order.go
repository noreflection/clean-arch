package domain

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strconv"
)

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

func NewOrder(id string, product string, quantity int) (*Order, error) {
	orderIDStr := generateUniqueID()
	i, err := strconv.Atoi(orderIDStr)

	if err != nil {
		// Handle the error if the conversion fails.
		fmt.Println("Conversion error:", err)
		return nil, err
	}
	customerID := generateUniqueID()

	return &Order{
		ID:         i,
		CustomerID: customerID,
	}, nil
}

// generateUniqueID generates a unique identifier for the customer.
func generateUniqueID() string {
	// You can use a library like "github.com/google/uuid" to generate UUIDs.
	// Example:
	uniqueID := uuid.New().String()
	return uniqueID
}
