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
	// CountStorehouse count storehouse
	CountStorehouse(ctx context.Context, query string) (int64, error)
	// GetListStorehouses select a list storehouses
	GetListStorehouses(ctx context.Context, query string, limit, skip int) ([]domain.Storehouse, error)
	// CountAuthorizedStorehouses count authorized storehouses
	CountAuthorizedStorehouses(ctx context.Context, userID int, query string) (int64, error)
	// GetAuthorizedStorehouses
	GetAuthorizedStorehouses(ctx context.Context, userID int, query string, limit, skip int) ([]domain.Storehouse, error)
	// GetStorehouseUsedCapacityByID get used capacity of storehouse
	GetStorehouseUsedCapacityByID(ctx context.Context, id int) (float64, error)
	// UpdateStorehouse update a storehouse, only update non-zero fields by default
	UpdateStorehouse(ctx context.Context, storehouses *domain.Storehouse) (*domain.Storehouse, error)
	// DeleteStorehouse delete a storehouse
	DeleteStorehouse(ctx context.Context, id int) error
}

type IStorehouseService interface {
	// CreateStorehouse insert a new storehouse into the database
	CreateStorehouse(ctx context.Context, storehouses *domain.Storehouse) (*domain.Storehouse, error)
	// GetStorehouseByID select a storehouse by id
	GetStorehouseByID(ctx context.Context, id int) (*domain.Storehouse, error)
	// GetListStorehouses select a list storehouses
	GetListStorehouses(ctx context.Context, query string, limit, skip int) ([]domain.Storehouse, error)
	// GetAuthorizedStorehouses
	GetAuthorizedStorehouses(ctx context.Context, userID int, query string, limit, skip int) ([]domain.Storehouse, error)
	// GetStorehouseUsedCapacityByID get used capacity of storehouse
	GetStorehouseUsedCapacityByID(ctx context.Context, id int) (float64, error)
	// UpdateStorehouse update a storehouse, only update non-zero fields by default
	UpdateStorehouse(ctx context.Context, storehouses *domain.Storehouse) (*domain.Storehouse, error)
	// DeleteStorehouse delete a storehouse
	DeleteStorehouse(ctx context.Context, id int) error
}
