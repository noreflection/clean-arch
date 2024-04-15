package order

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-cqrs/cmd/command_handlers"
	"go-cqrs/cmd/query_handlers"
	"net/http"
	"strconv"
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
	HandleSuccessResponse(w, "order created successfully") //todo: add returned id
}

func (c *OrderController) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderIDStr, ok := vars["id"]
	if !ok {
		HandleErrorResponse(w, fmt.Errorf("order ID not found in URL"))
		return
	}

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		HandleErrorResponse(w, fmt.Errorf("invalid order ID: %s", orderIDStr))
		return
	}

	getQuery := query_handlers.GetOrderQuery{ID: orderID}
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

func (c *OrderController) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderIDStr, ok := vars["id"]
	if !ok {
		HandleErrorResponse(w, fmt.Errorf("order ID not found in URL"))
		return
	}

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		HandleErrorResponse(w, fmt.Errorf("invalid order ID: %s", orderIDStr))
		return
	}

	deleteCmd := command_handlers.DeleteOrderCommand{ID: orderID}
	err = c.commandHandler.HandleDeleteOrderCommand(r.Context(), deleteCmd)
	if err != nil {
		HandleErrorResponse(w, err)
		return
	}
	HandleSuccessResponse(w, fmt.Sprintf("order ID:%d has been successfully deleted", orderID))
}

func (c *OrderController) UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderIDStr, ok := vars["id"]
	if !ok {
		HandleErrorResponse(w, fmt.Errorf("order ID not found in URL"))
		return
	}

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		HandleErrorResponse(w, fmt.Errorf("invalid order ID: %s", orderIDStr))
		return
	}

	// Parse the request body to get the updated order data
	var updateCmd command_handlers.UpdateOrderCommand
	err = json.NewDecoder(r.Body).Decode(&updateCmd)
	if err != nil {
		HandleErrorResponse(w, fmt.Errorf("failed to parse request body: %v", err))
		return
	}

	// Set the ID in the update command
	updateCmd.ID = strconv.Itoa(orderID)

	// Call the command handler to update the order
	err = c.commandHandler.HandleUpdateOrderCommand(r.Context(), updateCmd)
	if err != nil {
		HandleErrorResponse(w, err)
		return
	}

	// Respond with a success message
	HandleSuccessResponse(w, fmt.Sprintf("Order ID:%d has been successfully updated", orderID))
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
