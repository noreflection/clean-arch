package customer

import (
	"encoding/json"
	"fmt"
	"go-cqrs/cmd/command_handlers"
	"go-cqrs/cmd/query_handlers"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller struct {
	commandHandler *command_handlers.CustomerCommandHandler
	queryHandler   *query_handlers.CustomerQueryHandler
}

func NewCustomerController(commandHandler *command_handlers.CustomerCommandHandler, queryHandler *query_handlers.CustomerQueryHandler) *Controller {
	return &Controller{commandHandler: commandHandler, queryHandler: queryHandler}
}

func (c *Controller) CreateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var createCmd command_handlers.CreateCustomerCommand
	err := json.NewDecoder(r.Body).Decode(&createCmd)
	if err != nil {
		HandleErrorResponse(w, err)
		return
	}
	defer r.Body.Close()

	CustomerId, err := c.commandHandler.HandleCreateCustomerCommand(r.Context(), createCmd)
	if err != nil {
		HandleErrorResponse(w, err)
		return
	}
	HandleSuccessResponse(w, fmt.Sprintf("Customer with id: %d created", CustomerId))
}

func (c *Controller) GetCustomerHandler(w http.ResponseWriter, r *http.Request) {
	CustomerIDStr := mux.Vars(r)["id"]
	CustomerID, err := strconv.Atoi(CustomerIDStr)
	if err != nil {
		HandleErrorResponse(w, fmt.Errorf("invalid Customer ID: %s", CustomerIDStr))
		return
	}

	getQuery := query_handlers.GetCustomerQuery{ID: CustomerID}
	Customer, err := c.queryHandler.HandleGetCustomerQuery(r.Context(), getQuery)
	if err != nil {
		HandleErrorResponse(w, err)
		return
	}
	HandleSuccessResponse(w, Customer)
}

func (c *Controller) DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	CustomerIDStr, ok := vars["id"]
	if !ok {
		HandleErrorResponse(w, fmt.Errorf("Customer ID not found in URL"))
		return
	}

	CustomerID, err := strconv.Atoi(CustomerIDStr)
	if err != nil {
		HandleErrorResponse(w, fmt.Errorf("invalid Customer ID: %s", CustomerIDStr))
		return
	}

	deleteCmd := command_handlers.DeleteCustomerCommand{ID: CustomerID}
	err = c.commandHandler.HandleDeleteCustomerCommand(r.Context(), deleteCmd)
	if err != nil {
		HandleErrorResponse(w, err)
		return
	}
	HandleSuccessResponse(w, fmt.Sprintf("Customer ID:%d has been successfully deleted", CustomerID))
}

func (c *Controller) UpdateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	CustomerIDStr, ok := vars["id"]
	if !ok {
		HandleErrorResponse(w, fmt.Errorf("Customer ID not found in URL"))
		return
	}

	CustomerID, err := strconv.Atoi(CustomerIDStr)
	if err != nil {
		HandleErrorResponse(w, fmt.Errorf("invalid Customer ID: %s", CustomerIDStr))
		return
	}

	var updateCmd command_handlers.UpdateCustomerCommand
	err = json.NewDecoder(r.Body).Decode(&updateCmd)
	if err != nil {
		HandleErrorResponse(w, fmt.Errorf("failed to parse request body: %v", err))
		return
	}
	defer r.Body.Close()

	updateCmd.ID = CustomerID
	err = c.commandHandler.HandleUpdateCustomerCommand(r.Context(), updateCmd)
	if err != nil {
		HandleErrorResponse(w, err)
		return
	}
	HandleSuccessResponse(w, fmt.Sprintf("Customer ID:%d has been successfully updated", CustomerID))
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func HandleErrorResponse(w http.ResponseWriter, errorMessage error) {
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
