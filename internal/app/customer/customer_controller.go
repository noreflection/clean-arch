// customer_controller.go

package customer

import (
	"encoding/json"
	"go-cqrs/cmd/command_handlers"

	"net/http"
)

// CustomerController handles HTTP requests related to customers.
type CustomerController struct {
	commandHandler *command_handlers.CustomerCommandHandler
}

// NewCustomerController creates a new instance of CustomerController.
func NewCustomerController(commandHandler *command_handlers.CustomerCommandHandler) *CustomerController {
	return &CustomerController{commandHandler: commandHandler}
}

// CreateCustomerHandler is an HTTP handler for creating a new customer.
func (c *CustomerController) CreateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request data and validate it
	// For simplicity, let's assume the data is in JSON format
	// You can use a JSON parsing library (e.g., encoding/json) to parse the request body
	// Ensure you handle errors properly in a production-ready code

	// Sample request body: {"ID": "1", "Name": "John Doe", "Email": "john@example.com"}

	// Parse the request body
	var createCmd command_handlers.CreateCustomerCommand
	// Use your JSON parsing library here to decode the request body into createCmd

	// Handle the create customer command
	if err := c.commandHandler.HandleCreateCustomerCommand(r.Context(), createCmd); err != nil {
		// Handle command execution error (e.g., return a JSON response with an error message)
		// HandleErrorResponse is a fictional function you should implement
		HandleErrorResponse(w, err)
		return
	}

	// Customer created successfully, return a success response
	// You can define a success response format and implement it here
	// HandleSuccessResponse is a fictional function you should implement
	HandleSuccessResponse(w, "Customer created successfully")
}

// ErrorResponse Add other customer-related HTTP handlers here
// Define a struct for error responses
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse Define a struct for success responses
type SuccessResponse struct {
	Message string `json:"message"`
}

// HandleErrorResponse sends an error response with the specified error message.
func HandleErrorResponse(w http.ResponseWriter, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest) // You can set the appropriate HTTP status code for errors

	// Create and marshal the error response
	response := ErrorResponse{Error: errorMessage}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		// Handle JSON marshaling error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send the JSON response
	_, err = w.Write(jsonResponse)
	if err != nil {
		// Handle write error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// HandleSuccessResponse sends a success response with the specified message.
func HandleSuccessResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // You can set the appropriate HTTP status code for success

	// Create and marshal the success response
	response := SuccessResponse{Message: message}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		// Handle JSON marshaling error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send the JSON response
	_, err = w.Write(jsonResponse)
	if err != nil {
		// Handle write error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
