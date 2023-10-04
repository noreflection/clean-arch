package domain

type CustomerService interface {
	CreateCustomer(id, name, email string) error
	// ... other methods you might need
}
