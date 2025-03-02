package domain

type Order struct {
	ID         int
	CustomerID *int // Using pointer instead of sql.NullInt64 to represent optional value
	Product    string
	Quantity   int
	// Could add other domain-related fields like:
	// Status    OrderStatus
	// CreatedAt time.Time
}

// OrderStatus represents the current state of an order
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusConfirmed OrderStatus = "CONFIRMED"
	OrderStatusShipped   OrderStatus = "SHIPPED"
	OrderStatusDelivered OrderStatus = "DELIVERED"
	OrderStatusCancelled OrderStatus = "CANCELLED"
)

// NewOrder creates a new Order entity
func NewOrder(product string, quantity int) *Order {
	return &Order{
		Product:  product,
		Quantity: quantity,
	}
}

// AssignCustomer assigns a customer to this order
func (o *Order) AssignCustomer(customerID int) {
	o.CustomerID = &customerID
}
