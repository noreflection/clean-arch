package web

import (
	"net/http"
	
	"github.com/gorilla/mux"
	"go-cqrs/internal/interface/controller"
)

// Router is an interface for the HTTP router
type Router interface {
	http.Handler
	SetupRoutes()
}

// MuxRouter implements Router using gorilla/mux
type MuxRouter struct {
	*mux.Router
	customerController controller.CustomerController
	orderController    controller.OrderController
}

// NewRouter creates a new router with the given controllers
func NewRouter(customerController controller.CustomerController, orderController controller.OrderController) Router {
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
	r.Use(loggingMiddleware)
	r.Use(corsMiddleware)
	
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
	orders.HandleFunc("/{id:[0-9]+}/customer/{customerId:[0-9]+}", r.orderController.AssignCustomer).Methods(http.MethodPut)
}

// loggingMiddleware logs information about each request
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request details here
		next.ServeHTTP(w, r)
	})
}

// corsMiddleware handles CORS headers
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
} 