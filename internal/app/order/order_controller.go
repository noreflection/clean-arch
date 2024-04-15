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

type Controller struct {
	commandHandler *command_handlers.OrderCommandHandler
	queryHandler   *query_handlers.OrderQueryHandler
}

func NewOrderController(commandHandler *command_handlers.OrderCommandHandler, queryHandler *query_handlers.OrderQueryHandler) *Controller {
	return &Controller{commandHandler: commandHandler, queryHandler: queryHandler}
}

func (c *Controller) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var createCmd command_handlers.CreateOrderCommand
	err := json.NewDecoder(r.Body).Decode(&createCmd)
	if err != nil {
		HandleErrorResponse(w, err)
		return
	}
	defer r.Body.Close()

	orderId, err := c.commandHandler.HandleCreateOrderCommand(r.Context(), createCmd)
	if err != nil {
		HandleErrorResponse(w, err)
		return
	}
	HandleSuccessResponse(w, fmt.Sprintf("order with id: %d created", orderId))
}

func (c *Controller) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderIDStr := mux.Vars(r)["id"]
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
	HandleSuccessResponse(w, order)
}

func (c *Controller) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
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

func (c *Controller) UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
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

	var updateCmd command_handlers.UpdateOrderCommand
	err = json.NewDecoder(r.Body).Decode(&updateCmd)
	if err != nil {
		HandleErrorResponse(w, fmt.Errorf("failed to parse request body: %v", err))
		return
	}
	defer r.Body.Close()

	updateCmd.ID = orderID
	err = c.commandHandler.HandleUpdateOrderCommand(r.Context(), updateCmd)
	if err != nil {
		HandleErrorResponse(w, err)
		return
	}
	HandleSuccessResponse(w, fmt.Sprintf("Order ID:%d has been successfully updated", orderID))
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func HandleErrorResponse(w http.ResponseWriter, errorMessage error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

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

func HandleSuccessResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(data)
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
