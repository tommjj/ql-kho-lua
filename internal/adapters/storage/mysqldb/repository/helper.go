package repository

import (
	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb/schema"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

// convertToUser is a helper to convert schema user to domain user type
func convertToUser(u *schema.User) *domain.User {
	return &domain.User{
		ID:       u.ID,
		Name:     u.Name,
		Phone:    u.Phone,
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
	}
}

// convertToWarehouse is a helper to convert schema warehouse to domain warehouse type
func convertToWarehouse(s *schema.Warehouse) *domain.Warehouse {
	return &domain.Warehouse{
		ID:       s.ID,
		Name:     s.Name,
		Location: s.Location,
		Capacity: s.Capacity,
		Image:    s.Image,
	}
}

// convertToRice is a helper to convert schema rice to domain rice type
func convertToRice(r *schema.Rice) *domain.Rice {
	return &domain.Rice{
		ID:   r.ID,
		Name: r.Name,
	}
}

// convertToCustomer is a helper to convert schema customer to domain customer type
func convertToCustomer(c *schema.Customer) *domain.Customer {
	return &domain.Customer{
		ID:      c.ID,
		Name:    c.Name,
		Email:   c.Email,
		Phone:   c.Phone,
		Address: c.Address,
	}
}
