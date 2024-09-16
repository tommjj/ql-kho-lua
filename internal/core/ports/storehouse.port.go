package ports

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

type IStorehouseRepository interface {
	// CreateStorehouse insert a new storehouse into the database
	CreateStorehouse(ctx context.Context, storehouses *domain.Storehouse) (*domain.Storehouse, error)
	// GetStorehouseByID select a storehouse by id
	GetStorehouseByID(ctx context.Context, id int) (*domain.Storehouse, error)
	// GetListStorehouses select a list storehouses
	GetListStorehouses(ctx context.Context, query string, limit, skip int) ([]domain.Storehouse, error)
	// UpdateStorehouse update a storehouse, only update non-zero fields by default
	UpdateStorehouse(ctx context.Context, storehouses *domain.Storehouse) (*domain.Storehouse, error)
	// DeleteStorehouse delete a storehouse
	DeleteStorehouse(ctx context.Context, id int) error
}
