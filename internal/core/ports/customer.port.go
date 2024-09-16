package ports

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

type ICustomerRepository interface {
	// CreateCustomer insert a new customer into the database
	CreateCustomer(ctx context.Context, customer *domain.Customer) (*domain.Customer, error)
	// GetCustomerByID select a customer by id
	GetCustomerByID(ctx context.Context, id int) (*domain.Customer, error)
	// GetListCustomers select a customer
	GetListCustomers(ctx context.Context, query string, limit, skip int) ([]domain.Customer, error)
	// UpdateCustomer update a customer, only update non-zero fields by default
	UpdateCustomer(ctx context.Context, customer *domain.Customer) (*domain.Customer, error)
	// DeleteCustomer delete a customer
	DeleteCustomer(ctx context.Context, id int) error
}
