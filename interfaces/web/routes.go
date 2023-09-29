// routes.go

package web

import (
	"github.com/gorilla/mux"
	"go-cqrs/internal/app/customer"
	"net/http"

	//"go-cqrs/internal/app/customer"

	//"go-cqrs/internal/app/customer"
	"go-cqrs/internal/app/order"
)

// SetupRouter configures API routes.
func SetupRouter(
	orderService order.OrderService,
	customerService customer.CustomerService,
	customerController customer.CustomerController,
	orderController order.OrderController) *mux.Router {
	router := mux.NewRouter()

	// Order routes
	//orderController := order.NewOrderController(orderService)
	router.HandleFunc("/orders", orderController.CreateOrderHandler).Methods(http.MethodPost)
	router.HandleFunc("/orders/{id}", orderController.CreateOrderHandler).Methods(http.MethodGet)
	// Add more order routes as needed

	// Customer routes
	//customerController := customer.NewCustomerController(customerService)
	router.HandleFunc("/customers", customerController.CreateCustomerHandler).Methods(http.MethodPost)
	router.HandleFunc("/customers/{id}", customerController.CreateCustomerHandler).Methods(http.MethodGet)
	// Add more customer routes as needed

	// Add middleware and other routes here

	return router
}
