package order

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-cqrs/internal/domain"
	"net/http"
)

// Import the OrderService interface from the correct package

// OrderController represents the controller for orders.
type OrderController struct {
	service OrderService // Use the imported OrderService interface
}

// NewOrderController creates a new order controller.
func NewOrderController(service OrderService) *OrderController {
	return &OrderController{service}
}

// CreateOrder handles the creation of a new order.
func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder domain.Order
	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	createdOrder, err := c.service.CreateOrder(newOrder)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdOrder)
}

// GetOrder handles retrieving an order by ID.
func (c *OrderController) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	foundOrder, err := c.service.GetOrder(orderID)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundOrder)
}

// Implement other order-related controller methods here
