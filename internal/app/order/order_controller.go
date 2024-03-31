// order_controller.go

package order

import (
	"encoding/json"
	"go-cqrs/cmd/command_handlers"
	"net/http"
)

// OrderController handles HTTP requests related to orders.
type OrderController struct {
	commandHandler *command_handlers.OrderCommandHandler
}

// NewOrderController creates a new instance of OrderController.
func NewOrderController(commandHandler *command_handlers.OrderCommandHandler) *OrderController {
	return &OrderController{commandHandler: commandHandler}
}

// CreateOrderHandler is an HTTP handler for creating a new order.
func (c *OrderController) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request data and validate it
	// For simplicity, let's assume the data is in JSON format
	// You can use a JSON parsing library (e.g., encoding/json) to parse the request body
	// Ensure you handle errors properly in a production-ready code

	// Sample request body: {"ID": "1", "Product": "Product A", "Quantity": 5}

	// Parse the request body
	var createCmd command_handlers.CreateOrderCommand
	// Use your JSON parsing library here to decode the request body into createCmd
	//var createCmd command_handlers.CreateOrderCommand
	err := json.NewDecoder(r.Body).Decode(&createCmd)
	if err != nil {
		// Handle parsing error
		HandleErrorResponse(w, err)
		return
	}

	// Handle the create order command
	if err := c.commandHandler.HandleCreateOrderCommand(r.Context(), createCmd); err != nil {
		// Handle command execution error (e.g., return a JSON response with an error message)
		// HandleErrorResponse is a fictional function you should implement
		HandleErrorResponse(w, err)
		return
	}

	// Order created successfully, return a success response
	// You can define a success response format and implement it here
	// HandleSuccessResponse is a fictional function you should implement
	HandleSuccessResponse(w, "Order created successfully")
}

// ErrorResponse Define a struct for error responses
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse Define a struct for success responses
type SuccessResponse struct {
	Message string `json:"message"`
}

// HandleErrorResponse sends an error response with the specified error message.
func HandleErrorResponse(w http.ResponseWriter, errorMessage error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest) // You can set the appropriate HTTP status code for errors

	// Create and marshal the error response
	response := ErrorResponse{Error: errorMessage.Error()}
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
