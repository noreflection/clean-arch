package customer

import (
	"encoding/json"
	"go-cqrs/cmd/command_handlers"
	"net/http"
)

type CustomerController struct {
	commandHandler *command_handlers.CustomerCommandHandler
}

func NewCustomerController(commandHandler *command_handlers.CustomerCommandHandler) *CustomerController {
	return &CustomerController{commandHandler: commandHandler}
}

func (c *CustomerController) CreateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request data and validate it
	// For simplicity, let's assume the data is in JSON format
	// You can use a JSON parsing library (e.g., encoding/json) to parse the request body
	// Ensure you handle errors properly in a production-ready code

	// Sample request body: {"ID": "1", "Name": "John Doe", "Email": "john@example.com"}

	var createCmd command_handlers.CreateCustomerCommand
	// Use your JSON parsing library here to decode the request body into createCmd
	if err := c.commandHandler.HandleCreateCustomerCommand(createCmd); err != nil {
		HandleErrorResponse(w, err)
		return
	}
	HandleSuccessResponse(w, "Customer created successfully")
}

type ErrorResponse struct {
	Error string `json:"error"`
}

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
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func HandleSuccessResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // You can set the appropriate HTTP status code for success
	response := SuccessResponse{Message: message}
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
