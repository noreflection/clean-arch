package domain

type CustomerService interface {
	CreateCustomer(id, name, email string) error
}
