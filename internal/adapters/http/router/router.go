package router

import (
	"net/http"

	"go-cqrs/internal/adapters/http/controllers"
	"go-cqrs/internal/adapters/http/middleware"

	"github.com/gorilla/mux"
)

// Router is an interface for the HTTP router
type Router interface {
	http.Handler
	SetupRoutes()
}

// MuxRouter implements Router using gorilla/mux
type MuxRouter struct {
	*mux.Router
	customerController controllers.CustomerController
	orderController    controllers.OrderController
}

// NewRouter creates a new router with the given controllers
func NewRouter(customerController controllers.CustomerController, orderController controllers.OrderController) Router {
	r := &MuxRouter{
		Router:             mux.NewRouter(),
		customerController: customerController,
		orderController:    orderController,
	}
	r.SetupRoutes()
	return r
}

// SetupRoutes configures all the routes for the application
func (r *MuxRouter) SetupRoutes() {
	// Middleware
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.CorsMiddleware)

	// API Routes
	api := r.PathPrefix("/api").Subrouter()

	// Health check
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods(http.MethodGet)

	// Customer routes
	customers := api.PathPrefix("/customers").Subrouter()
	customers.HandleFunc("", r.customerController.CreateCustomer).Methods(http.MethodPost)
	customers.HandleFunc("", r.customerController.ListCustomers).Methods(http.MethodGet)
	customers.HandleFunc("/{id:[0-9]+}", r.customerController.GetCustomer).Methods(http.MethodGet)
	customers.HandleFunc("/{id:[0-9]+}", r.customerController.UpdateCustomer).Methods(http.MethodPut)
	customers.HandleFunc("/{id:[0-9]+}", r.customerController.DeleteCustomer).Methods(http.MethodDelete)

	// Order routes
	orders := api.PathPrefix("/orders").Subrouter()
	orders.HandleFunc("", r.orderController.CreateOrder).Methods(http.MethodPost)
	orders.HandleFunc("", r.orderController.ListOrders).Methods(http.MethodGet)
	orders.HandleFunc("/{id:[0-9]+}", r.orderController.GetOrder).Methods(http.MethodGet)
	orders.HandleFunc("/{id:[0-9]+}", r.orderController.UpdateOrder).Methods(http.MethodPut)
	orders.HandleFunc("/{id:[0-9]+}", r.orderController.DeleteOrder).Methods(http.MethodDelete)
	orders.HandleFunc("/{id:[0-9]+}/customers/{customerId:[0-9]+}", r.orderController.AssignCustomer).Methods(http.MethodPost)
}
