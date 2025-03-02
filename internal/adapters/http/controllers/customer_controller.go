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

type CustomerController struct {
	commandHandler *commands.CustomerCommandHandler
	queryHandler   *queries.CustomerQueryHandler
}

func NewCustomerController(commandHandler *commands.CustomerCommandHandler, queryHandler *queries.CustomerQueryHandler) *CustomerController {
	return &CustomerController{commandHandler: commandHandler, queryHandler: queryHandler}
}

// CreateCustomer handles the creation of a new customer
func (c *CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var createCmd commands.CreateCustomerCommand
	err := json.NewDecoder(r.Body).Decode(&createCmd)
	if err != nil {
		HandleCustomerErrorResponse(w, err)
		return
	}

	customerID, err := c.commandHandler.HandleCreateCustomerCommand(r.Context(), createCmd)
	if err != nil {
		HandleCustomerErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      customerID,
		"message": "Customer created successfully",
	})
}

// GetCustomer handles retrieving a customer by ID
func (c *CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		HandleCustomerErrorResponse(w, fmt.Errorf("invalid customer ID: %w", err))
		return
	}

	getQuery := queries.GetCustomerQuery{ID: id}
	customer, err := c.queryHandler.HandleGetCustomerQuery(r.Context(), getQuery)
	if err != nil {
		HandleCustomerErrorResponse(w, err)
		return
	}

	if customer == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Customer not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// UpdateCustomer handles updating an existing customer
func (c *CustomerController) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		HandleCustomerErrorResponse(w, fmt.Errorf("invalid customer ID: %w", err))
		return
	}

	var updateCmd commands.UpdateCustomerCommand
	err = json.NewDecoder(r.Body).Decode(&updateCmd)
	if err != nil {
		HandleCustomerErrorResponse(w, err)
		return
	}
	updateCmd.ID = id

	err = c.commandHandler.HandleUpdateCustomerCommand(r.Context(), updateCmd)
	if err != nil {
		HandleCustomerErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer updated successfully"})
}

// DeleteCustomer handles deleting a customer
func (c *CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		HandleCustomerErrorResponse(w, fmt.Errorf("invalid customer ID: %w", err))
		return
	}

	deleteCmd := commands.DeleteCustomerCommand{ID: id}
	err = c.commandHandler.HandleDeleteCustomerCommand(r.Context(), deleteCmd)
	if err != nil {
		HandleCustomerErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer deleted successfully"})
}

// ListCustomers handles retrieving a list of customers
func (c *CustomerController) ListCustomers(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list customers functionality
	json.NewEncoder(w).Encode(map[string]string{"message": "List customers not implemented yet"})
}

// HandleCustomerErrorResponse handles error responses for customer endpoints
func HandleCustomerErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
