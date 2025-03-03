package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb/schema"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
	"gorm.io/gorm"
)

type customerRepository struct {
	db *mysqldb.MysqlDB
}

func NewCustomerRepository(db *mysqldb.MysqlDB) ports.ICustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (cr *customerRepository) CreateCustomer(ctx context.Context, customer *domain.Customer) (*domain.Customer, error) {
	createData := &schema.Customer{
		Name:    customer.Name,
		Email:   customer.Email,
		Phone:   customer.Phone,
		Address: customer.Address,
	}

	err := cr.db.WithContext(ctx).Create(createData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return convertToCustomer(createData), nil
}

func (cr *customerRepository) GetCustomerByID(ctx context.Context, id int) (*domain.Customer, error) {
	customer := &schema.Customer{}

	err := cr.db.WithContext(ctx).Where("id = ?", id).First(customer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return convertToCustomer(customer), nil
}

func (cr *customerRepository) CountCustomers(ctx context.Context, query string) (int64, error) {
	var count int64
	var err error

	q := cr.db.WithContext(ctx).Table("customers")

	trimQuery := strings.TrimSpace(query)
	if trimQuery != "" {
		q.Where("name LIKE ?", fmt.Sprintf("%%%v%%", trimQuery))
	}

	err = q.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (cr *customerRepository) GetListCustomers(ctx context.Context, query string, limit, skip int) ([]domain.Customer, error) {
	customers := []domain.Customer{}
	var err error

	q := cr.db.WithContext(ctx).Table("customers").
		Limit(limit).Offset((skip - 1) * limit).Order("name DESC").Where("deleted_at is NULL")

	trimQuery := strings.TrimSpace(query)
	if trimQuery != "" {
		q.Where("name LIKE ?", fmt.Sprintf("%%%v%%", trimQuery))
	}

	err = q.Scan(&customers).Error
	if err != nil {
		return nil, err
	}
	if len(customers) == 0 {
		return nil, domain.ErrDataNotFound
	}

	return customers, nil
}

func (cr *customerRepository) UpdateCustomer(ctx context.Context, customer *domain.Customer) (*domain.Customer, error) {
	updatedData := &schema.Customer{}

	result := cr.db.WithContext(ctx).Model(updatedData).Where("id = ?", customer.ID).
		Updates(&schema.Customer{
			Name:    customer.Name,
			Email:   customer.Email,
			Phone:   customer.Phone,
			Address: customer.Address,
		})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrConflictingData
		}
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, domain.ErrNoUpdatedData
	}

	return cr.GetCustomerByID(ctx, customer.ID)
}

func (cr *customerRepository) DeleteCustomer(ctx context.Context, id int) error {
	result := cr.db.WithContext(ctx).Where("id = ?", id).Delete(&schema.Customer{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrDataNotFound
	}

	return nil
}
