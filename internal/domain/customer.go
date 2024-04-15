package domain

type Customer struct {
	ID    string
	Name  string
	Email string
	//CreatedAt time.Time
	//UpdatedAt time.Time
}

func NewCustomer(name string, email string) *Customer {
	var customerID string
	return &Customer{
		ID:    customerID,
		Name:  name,
		Email: email,
	}
}
