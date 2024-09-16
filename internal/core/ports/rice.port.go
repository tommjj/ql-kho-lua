package ports

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

type IRiceRepository interface {
	// CreateRice insert a new rice into the database
	CreateRice(ctx context.Context, rice *domain.Rice) (*domain.Rice, error)
	// GetRiceByID select a rice by id
	GetRiceByID(ctx context.Context, id int) (*domain.Rice, error)
	// GetListRice select a rice
	GetListRice(ctx context.Context, query string, skip, limit int) ([]domain.Rice, error)
	// UpdateRice update a rice, only update non-zero fields by default
	UpdateRice(ctx context.Context, rice *domain.Rice) (*domain.Rice, error)
	// DeleteRice delete a rice
	DeleteRice(ctx context.Context, id int) error
}
