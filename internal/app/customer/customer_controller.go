package customer

import (
	"cqrs-web-api/internal/domain"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Import the CustomerService interface from the correct package

// CustomerController represents the controller for customers.
type CustomerController struct {
	service CustomerService // Use the imported CustomerService interface
}

// NewCustomerController creates a new customer controller.
func NewCustomerController(service CustomerService) *CustomerController {
	return &CustomerController{service}
}

// CreateCustomer handles the creation of a new customer.
func (c *CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var newCustomer domain.Customer
	if err := json.NewDecoder(r.Body).Decode(&newCustomer); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	createdCustomer, err := c.service.CreateCustomer(newCustomer)
	if err != nil {
		http.Error(w, "Failed to create customer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdCustomer)
}

// GetCustomer handles retrieving a customer by ID.
func (c *CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["id"]

	foundCustomer, err := c.service.GetCustomer(customerID)
	if err != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundCustomer)
}

// Implement other customer-related controller methods here
