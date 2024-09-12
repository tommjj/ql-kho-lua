package ports

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

type IUserRepository interface {
	// Create insert a new user into the database
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	// GetByID select a user by id
	GetByID(ctx context.Context, id int) (*domain.User, error)
	// GetByEmail select a user by email
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	// GetList select a list users
	GetList(ctx context.Context, skip, limit int) ([]domain.User, error)
	// Update update a user, only update non-zero fields by default
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	// Delete delete a user
	Delete(ctx context.Context, id int) error
}
