package web

import (
	"github.com/gorilla/mux"
	"go-cqrs/internal/interface/controller"
	"net/http"
)

func SetupRouter(customerController controller.CustomerController, orderController controller.OrderController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/orders", orderController.CreateOrderHandler).Methods(http.MethodPost)
	router.HandleFunc("/orders/{id}", orderController.GetOrderHandler).Methods(http.MethodGet)
	router.HandleFunc("/orders/{id}", orderController.UpdateOrderHandler).Methods(http.MethodPatch)
	router.HandleFunc("/orders/{id}", orderController.DeleteOrderHandler).Methods(http.MethodDelete)

	router.HandleFunc("/customers", customerController.CreateCustomerHandler).Methods(http.MethodPost)
	router.HandleFunc("/customers/{id}", customerController.CreateCustomerHandler).Methods(http.MethodGet)
	return router
}
