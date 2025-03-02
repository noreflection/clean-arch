package controllers

import (
	"encoding/json"
	"fmt"
	"go-cqrs/internal/adapters/cqrs/commands"
	"go-cqrs/internal/adapters/cqrs/queries"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OrderController struct {
	commandHandler *commands.OrderCommandHandler
	queryHandler   *queries.OrderQueryHandler
}

func NewOrderController(commandHandler *commands.OrderCommandHandler, queryHandler *queries.OrderQueryHandler) *OrderController {
	return &OrderController{commandHandler: commandHandler, queryHandler: queryHandler}
}

// CreateOrder handles the creation of a new order
func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var createCmd commands.CreateOrderCommand
	err := json.NewDecoder(r.Body).Decode(&createCmd)
	if err != nil {
		HandleOrderErrorResponse(w, err)
		return
	}

	orderID, err := c.commandHandler.HandleCreateOrderCommand(r.Context(), createCmd)
	if err != nil {
		HandleOrderErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      orderID,
		"message": "Order created successfully",
	})
}

// GetOrder handles retrieving an order by ID
func (c *OrderController) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		HandleOrderErrorResponse(w, fmt.Errorf("invalid order ID: %w", err))
		return
	}

	getQuery := queries.GetOrderQuery{ID: id}
	order, err := c.queryHandler.HandleGetOrderQuery(r.Context(), getQuery)
	if err != nil {
		HandleOrderErrorResponse(w, err)
		return
	}

	if order == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Order not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// UpdateOrder handles updating an existing order
func (c *OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		HandleOrderErrorResponse(w, fmt.Errorf("invalid order ID: %w", err))
		return
	}

	var updateCmd commands.UpdateOrderCommand
	err = json.NewDecoder(r.Body).Decode(&updateCmd)
	if err != nil {
		HandleOrderErrorResponse(w, err)
		return
	}
	updateCmd.ID = id

	err = c.commandHandler.HandleUpdateOrderCommand(r.Context(), updateCmd)
	if err != nil {
		HandleOrderErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order updated successfully"})
}

// DeleteOrder handles deleting an order
func (c *OrderController) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		HandleOrderErrorResponse(w, fmt.Errorf("invalid order ID: %w", err))
		return
	}

	deleteCmd := commands.DeleteOrderCommand{ID: id}
	err = c.commandHandler.HandleDeleteOrderCommand(r.Context(), deleteCmd)
	if err != nil {
		HandleOrderErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order deleted successfully"})
}

// ListOrders handles retrieving a list of orders
func (c *OrderController) ListOrders(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list orders functionality
	json.NewEncoder(w).Encode(map[string]string{"message": "List orders not implemented yet"})
}

// AssignCustomer handles assigning a customer to an order
func (c *OrderController) AssignCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		HandleOrderErrorResponse(w, fmt.Errorf("invalid order ID: %w", err))
		return
	}

	customerID, err := strconv.Atoi(vars["customerId"])
	if err != nil {
		HandleOrderErrorResponse(w, fmt.Errorf("invalid customer ID: %w", err))
		return
	}

	assignCmd := commands.AssignCustomerCommand{
		OrderID:    orderID,
		CustomerID: customerID,
	}

	err = c.commandHandler.HandleAssignCustomerCommand(r.Context(), assignCmd)
	if err != nil {
		HandleOrderErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Customer %d assigned to order %d successfully", customerID, orderID),
	})
}

// HandleOrderErrorResponse handles error responses for order endpoints
func HandleOrderErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
