package repositories

import (
	"context"
	"database/sql"
	"errors"
	"go-cqrs/internal/domain"
)

// CustomerRepository implements ports.CustomerRepository
type CustomerRepository struct {
	db *sql.DB
}

// NewCustomerRepository creates a new CustomerRepository
func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

// Create inserts a new customer into the database
func (r *CustomerRepository) Create(ctx context.Context, customer domain.Customer) (int, error) {
	var customerID int

	err := r.db.QueryRowContext(ctx,
		"INSERT INTO customers (name, email) VALUES ($1, $2) RETURNING id",
		customer.Name, customer.Email).Scan(&customerID)

	if err != nil {
		return 0, errors.New("failed to create customer: " + err.Error())
	}

	return customerID, nil
}

// GetByID retrieves a customer by their ID
func (r *CustomerRepository) GetByID(ctx context.Context, id int) (*domain.Customer, error) {
	var customer domain.Customer

	err := r.db.QueryRowContext(ctx,
		"SELECT id, name, email FROM customers WHERE id = $1",
		id).Scan(&customer.ID, &customer.Name, &customer.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found, return nil without error
		}
		return nil, errors.New("failed to get customer: " + err.Error())
	}

	return &customer, nil
}

// GetByEmail retrieves a customer by their email
func (r *CustomerRepository) GetByEmail(ctx context.Context, email string) (*domain.Customer, error) {
	var customer domain.Customer

	err := r.db.QueryRowContext(ctx,
		"SELECT id, name, email FROM customers WHERE email = $1",
		email).Scan(&customer.ID, &customer.Name, &customer.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found, return nil without error
		}
		return nil, errors.New("failed to get customer by email: " + err.Error())
	}

	return &customer, nil
}

// Update updates an existing customer
func (r *CustomerRepository) Update(ctx context.Context, customer domain.Customer) error {
	_, err := r.db.ExecContext(ctx,
		"UPDATE customers SET name = $1, email = $2 WHERE id = $3",
		customer.Name, customer.Email, customer.ID)

	if err != nil {
		return errors.New("failed to update customer: " + err.Error())
	}

	return nil
}

// Delete removes a customer
func (r *CustomerRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM customers WHERE id = $1", id)

	if err != nil {
		return errors.New("failed to delete customer: " + err.Error())
	}

	return nil
}

// List retrieves customers with pagination
func (r *CustomerRepository) List(ctx context.Context, limit, offset int) ([]domain.Customer, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, name, email FROM customers LIMIT $1 OFFSET $2",
		limit, offset)

	if err != nil {
		return nil, errors.New("failed to list customers: " + err.Error())
	}
	defer rows.Close()

	var customers []domain.Customer
	for rows.Next() {
		var customer domain.Customer

		err := rows.Scan(&customer.ID, &customer.Name, &customer.Email)
		if err != nil {
			return nil, errors.New("failed to scan customer row: " + err.Error())
		}

		customers = append(customers, customer)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("error iterating customer rows: " + err.Error())
	}

	return customers, nil
}
