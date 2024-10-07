package services

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type customerService struct {
	repo ports.ICustomerRepository
}

func NewCustomerService(repo ports.ICustomerRepository) ports.ICustomerService {
	return &customerService{
		repo: repo,
	}
}

func (c *customerService) CreateCustomer(ctx context.Context, customer *domain.Customer) (*domain.Customer, error) {
	created, err := c.repo.CreateCustomer(ctx, customer)
	if err != nil {
		switch err {
		case domain.ErrConflictingData:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}

	return created, nil
}

func (c *customerService) GetCustomerByID(ctx context.Context, id int) (*domain.Customer, error) {
	customer, err := c.repo.GetCustomerByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return customer, nil
}

func (c *customerService) GetListCustomers(ctx context.Context, query string, limit, skip int) ([]domain.Customer, error) {
	customers, err := c.repo.GetListCustomers(ctx, query, limit, skip)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return customers, nil
}

func (c *customerService) UpdateCustomer(ctx context.Context, customer *domain.Customer) (*domain.Customer, error) {
	_, err := c.repo.GetCustomerByID(ctx, customer.ID)
	if err != nil {
		switch err {
		case domain.ErrDataNotFound:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}

	updated, err := c.repo.UpdateCustomer(ctx, customer)
	if err != nil {
		switch err {
		case domain.ErrConflictingData:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}

	return updated, nil
}

func (c *customerService) DeleteCustomer(ctx context.Context, id int) error {
	err := c.repo.DeleteCustomer(ctx, id)
	if err != nil {
		switch err {
		case domain.ErrDataNotFound:
			return err
		default:
			return domain.ErrInternal
		}
	}
	return nil
}
