package repository

import (
	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb/schema"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

// convertToDomainUser is a helper to convert schema user to domain user type
func convertToDomainUser(u *schema.User) *domain.User {
	return &domain.User{
		ID:       u.ID,
		Name:     u.Name,
		Phone:    u.Phone,
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
	}
}

// convertToDomainStorehouse is a helper to convert schema storehouse to domain storehouse type
func convertToDomainStorehouse(s *schema.Storehouse) *domain.Storehouse {
	return &domain.Storehouse{
		ID:       s.ID,
		Name:     s.Name,
		Location: s.Location,
		Capacity: s.Capacity,
	}
}

// convertToDomainRice is a helper to convert schema rice to domain rice type
func convertToDomainRice(r *schema.Rice) *domain.Rice {
	return &domain.Rice{
		ID:   r.ID,
		Name: r.Name,
	}
}
