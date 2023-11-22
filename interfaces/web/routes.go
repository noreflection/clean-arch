package web

import (
	"github.com/gorilla/mux"
	"go-cqrs/internal/app/customer"
	"go-cqrs/internal/app/order"
	"net/http"
)

func SetupRouter(
	customerController customer.CustomerController,
	orderController order.OrderController) *mux.Router {
	router := mux.NewRouter()

	// Order routes
	router.HandleFunc("/orders", orderController.CreateOrderHandler).Methods(http.MethodPost)
	router.HandleFunc("/orders/{id}", orderController.CreateOrderHandler).Methods(http.MethodGet)

	// Customer routes
	router.HandleFunc("/customers", customerController.CreateCustomerHandler).Methods(http.MethodPost)
	router.HandleFunc("/customers/{id}", customerController.CreateCustomerHandler).Methods(http.MethodGet)
	return router
}
