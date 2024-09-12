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
