package ports

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

type IUserRepository interface {
	// CreateUser insert a new user into the database
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// GetUserByID select a user by id
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	// GetUserByEmail select a user by email
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	// CountUsers count user
	CountUsers(ctx context.Context, query string) (int64, error)
	// GetListUsers select a list users
	GetListUsers(ctx context.Context, query string, limit, skip int) ([]domain.User, error)
	// UpdateUser update a user, only update non-zero fields by default
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// DeleteUser delete a user
	DeleteUser(ctx context.Context, id int) error
}

type IUserService interface {
	// CreateUser create  new user
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// GetUserByID get a user by id
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	// GetUserByEmail get a user by email
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	// CountUsers count user
	CountUsers(ctx context.Context, query string) (int64, error)
	// GetListUsers get a list users
	GetListUsers(ctx context.Context, query string, limit, skip int) ([]domain.User, error)
	// UpdateUser update a user, only update non-zero fields by default
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// DeleteUser delete a user
	DeleteUser(ctx context.Context, id int) error
}
