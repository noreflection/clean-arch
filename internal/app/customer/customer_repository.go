package customer

import (
	"database/sql"
	"go-cqrs/internal/infrastructure/command_database"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db}
}

// Create inserts a new customer record into the database.
func (r *CustomerRepository) Create(name, email string) (int64, error) {
	cmdDB := command_database.NewCustomerCommandDB(r.db)
	return cmdDB.Create(name, email)
	// Validate input, perform any necessary business logic, etc.
	//customer := domain.Customer{ID: id, Name: name, Email: email}
	//// Execute SQL insert statement to create a new customer
	//_, err := r.db.Exec("INSERT INTO customers (name) VALUES (?)", customer.Name)
	//if err != nil {
	//	// Handle the error, such as rolling back a transaction or logging
	//	return err
	//}
	//
	//// If the insert was successful, you can update the customer ID field
	//// with the generated ID if needed (e.g., for auto-incrementing primary keys).
	//
	//return nil
}

// GetCustomerByID retrieves a customer record from the database by ID.
//func (r *CustomerRepository) GetCustomerByID(customerID int) (*Customer, error) {
//	// Execute SQL query to retrieve customer by ID
//	row := r.db.QueryRow("SELECT id, name FROM customers WHERE id = ?", customerID)
//
//	var customer Customer
//	err := row.Scan(&customer.ID, &customer.Name)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			// Handle the case where no customer with the given ID was found
//			return nil, nil
//		}
//		// Handle other errors
//		return nil, err
//	}
//
//	return &customer, nil
//}
