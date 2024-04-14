// order_controller.go

package order

import (
	"encoding/json"
	"go-cqrs/cmd/command_handlers"
	"go-cqrs/cmd/query_handlers"
	"net/http"
)

type OrderController struct {
	commandHandler *command_handlers.OrderCommandHandler
	queryHandler   *query_handlers.OrderQueryHandler
}

func NewOrderController(commandHandler *command_handlers.OrderCommandHandler, queryHandler *query_handlers.OrderQueryHandler) *OrderController {
	return &OrderController{commandHandler: commandHandler, queryHandler: queryHandler}
}

func (c *OrderController) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var createCmd command_handlers.CreateOrderCommand
	err := json.NewDecoder(r.Body).Decode(&createCmd)
	if err != nil {
		HandleErrorResponse(w, err)
		return
	}

	if err := c.commandHandler.HandleCreateOrderCommand(r.Context(), createCmd); err != nil {
		HandleErrorResponse(w, err)
		return
	}
	HandleSuccessResponse(w, "Order created successfully") //todo: add returned id
}

func (c *OrderController) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	var getQuery query_handlers.GetOrderQuery
	err := json.NewDecoder(r.Body).Decode(&getQuery)
	if err != nil {
		HandleErrorResponse(w, err)
		return
	}

	order, err := c.queryHandler.HandleGetOrderQuery(r.Context(), getQuery)
	if err != nil {
		HandleErrorResponse(w, err)
		return
	}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		HandleErrorResponse(w, err)
		return
	}
	HandleSuccessResponse(w, string(orderJSON))
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

func HandleErrorResponse(w http.ResponseWriter, errorMessage error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

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

func HandleSuccessResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

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
