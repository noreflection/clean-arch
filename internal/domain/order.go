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
}

func NewOrder(id string, product string, quantity int) (*Order, error) {
	orderIDStr := generateUniqueID()
	i, err := strconv.Atoi(orderIDStr)

	if err != nil {
		fmt.Println("Conversion error:", err)
		return nil, err
	}
	customerID := generateUniqueID()

	return &Order{
		ID:         i,
		CustomerID: customerID,
	}, nil
}

func generateUniqueID() string {
	// You can use a library like "github.com/google/uuid" to generate UUIDs.
	uniqueID := uuid.New().String()
	return uniqueID
}
