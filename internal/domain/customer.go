package domain

type Customer struct {
	ID      string
	Name    string
	Surname string
	Email   string
}

func NewCustomer(name string, email string) *Customer {
	var customerID string

	return &Customer{
		ID:    customerID,
		Name:  name,
		Email: email,
	}
}
