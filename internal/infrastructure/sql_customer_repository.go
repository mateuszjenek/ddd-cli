package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mateuszjenek/ddd-cli/internal/domain/model"
	"github.com/mateuszjenek/ddd-cli/internal/domain/model/valueobject"
	"github.com/mateuszjenek/ddd-cli/internal/domain/service"
)

type sqlCustomerRepository struct {
	db *sql.DB
}

func NewSqlCustomerRepository(db *sql.DB) service.CustomerRepository {
	return &sqlCustomerRepository{db}
}

func (r *sqlCustomerRepository) Create(ctx context.Context, customer model.Customer) (model.Customer, error) {
	result, err := r.db.ExecContext(
		ctx, 
		"INSERT INTO customers (uuid, first_name, last_name, email) VALUES ($1, $2, $3, $4)", 
		customer.Id.GetId(), 
		customer.FullName.GetFirstName(), 
		customer.FullName.GetLastName(), 
		customer.Email.GetEmail(),
	)
	if err != nil {
		return customer, fmt.Errorf("failed to insert customer: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return customer, fmt.Errorf("failed to get number of affected rows after customer insert: %v", err)
	}
	if rows != 1 {
		return customer, fmt.Errorf("insert customer expected to affect exactly one row, affected %d", rows)
	}

	return customer, nil
}

func (r *sqlCustomerRepository) Get(ctx context.Context, uuid valueobject.CustomerId) (model.Customer, error) {
	r.db.QueryRowContext(ctx, "SELECT * FROM customers WHERE uuid = $1", uuid.GetId())
	panic("unimplemented")
}

func (*sqlCustomerRepository) List(context.Context) ([]model.Customer, error) {
	panic("unimplemented")
}

func (*sqlCustomerRepository) Remove(context.Context, valueobject.CustomerId) error {
	panic("unimplemented")
}

func (*sqlCustomerRepository) Update(context.Context, model.Customer) (model.Customer, error) {
	panic("unimplemented")
}
