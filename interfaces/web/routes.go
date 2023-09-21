package web

import (
	"cqrs-web-api/internal/app/customer"
	"cqrs-web-api/internal/app/order"
	"github.com/gorilla/mux"
	"net/http"
)

// SetupRouter configures API routes.
func SetupRouter(orderService order.OrderService, customerService customer.CustomerService) *mux.Router {
	router := mux.NewRouter()

	// Order routes
	orderController := order.NewOrderController(orderService)
	router.HandleFunc("/orders", orderController.CreateOrder).Methods(http.MethodPost)
	router.HandleFunc("/orders/{id}", orderController.GetOrder).Methods(http.MethodGet)
	// Add more order routes as needed

	// Customer routes
	customerController := customer.NewCustomerController(customerService)
	router.HandleFunc("/customers", customerController.CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{id}", customerController.GetCustomer).Methods(http.MethodGet)
	// Add more customer routes as needed

	// Add middleware and other routes here

	return router
}