package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-cqrs/internal/interface/command_handlers"
	"go-cqrs/internal/interface/query_handlers"
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
		HandleOrderErrorResponse(w, err)
		return
	}
	defer r.Body.Close()

	orderId, err := c.commandHandler.HandleCreateOrderCommand(r.Context(), createCmd)
	if err != nil {
		HandleOrderErrorResponse(w, err)
		return
	}
	HandleOrderSuccessResponse(w, fmt.Sprintf("order with id: %d created", orderId))
}

func (c *OrderController) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderIDStr := mux.Vars(r)["id"]
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		HandleOrderErrorResponse(w, fmt.Errorf("invalid order ID: %s", orderIDStr))
		return
	}

	getQuery := query_handlers.GetOrderQuery{ID: orderID}
	order, err := c.queryHandler.HandleGetOrderQuery(r.Context(), getQuery)
	if err != nil {
		HandleOrderErrorResponse(w, err)
		return
	}
	HandleOrderSuccessResponse(w, order)
}

func (c *OrderController) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderIDStr, ok := vars["id"]
	if !ok {
		HandleOrderErrorResponse(w, fmt.Errorf("order ID not found in URL"))
		return
	}

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		HandleOrderErrorResponse(w, fmt.Errorf("invalid order ID: %s", orderIDStr))
		return
	}

	deleteCmd := command_handlers.DeleteOrderCommand{ID: orderID}
	err = c.commandHandler.HandleDeleteOrderCommand(r.Context(), deleteCmd)
	if err != nil {
		HandleOrderErrorResponse(w, err)
		return
	}
	HandleOrderSuccessResponse(w, fmt.Sprintf("order ID:%d has been successfully deleted", orderID))
}

func (c *OrderController) UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderIDStr, ok := vars["id"]
	if !ok {
		HandleOrderErrorResponse(w, fmt.Errorf("order ID not found in URL"))
		return
	}

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		HandleOrderErrorResponse(w, fmt.Errorf("invalid order ID: %s", orderIDStr))
		return
	}

	var updateCmd command_handlers.UpdateOrderCommand
	err = json.NewDecoder(r.Body).Decode(&updateCmd)
	if err != nil {
		HandleOrderErrorResponse(w, fmt.Errorf("failed to parse request body: %v", err))
		return
	}
	defer r.Body.Close()

	updateCmd.ID = orderID
	err = c.commandHandler.HandleUpdateOrderCommand(r.Context(), updateCmd)
	if err != nil {
		HandleOrderErrorResponse(w, err)
		return
	}
	HandleOrderSuccessResponse(w, fmt.Sprintf("Order ID:%d has been successfully updated", orderID))
}

type OrderErrorResponse struct {
	Error string `json:"error"`
}

func HandleOrderErrorResponse(w http.ResponseWriter, errorMessage error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	response := map[string]string{"error": errorMessage.Error()}
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

func HandleOrderSuccessResponse(w http.ResponseWriter, data interface{}) {
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
