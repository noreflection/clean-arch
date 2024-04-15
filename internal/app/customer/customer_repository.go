package customer

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"go-cqrs/internal/domain"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, customer domain.Customer) (int, error) {
	var orderID int
	err := r.db.QueryRow("INSERT INTO customers (name, email) VALUES ($1, $2) RETURNING id", customer.Name, customer.Email).Scan(&orderID)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to create order")
	}
	return orderID, nil
}

func (r *Repository) Get(ctx context.Context, orderID int) (domain.Customer, error) {
	var customer domain.Customer
	err := r.db.QueryRow("SELECT * FROM cutomers WHERE id = $1", orderID).Scan(&customer.ID, &customer.Name, &customer.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Customer{}, errors.Wrap(err, "failed to find customer")
		}
		return domain.Customer{}, errors.Wrap(err, "failed to get customer")
	}
	return customer, nil
}

func (r *Repository) Update(ctx context.Context, customer domain.Customer) error {
	query := "UPDATE customers SET name = $1, email = $2 WHERE id = $3"
	_, err := r.db.Exec(query, customer.Name, customer.Email)
	if err != nil {
		return errors.Wrap(err, "failed to update customer")
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare("DELETE FROM customers WHERE id = $1")
	if err != nil {
		return errors.Wrap(err, "failed to prepare statement for deleting customer")
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.Wrap(err, "failed to delete customer")
	}
	return nil
}
