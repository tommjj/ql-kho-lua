package ports

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

type IWarehouseRepository interface {
	// CreateWarehouse insert a new warehouse into the database
	CreateWarehouse(ctx context.Context, warehouses *domain.Warehouse) (*domain.Warehouse, error)
	// GetWarehouseByID select a warehouse by id
	GetWarehouseByID(ctx context.Context, id int) (*domain.Warehouse, error)
	// CountWarehouses count warehouse
	CountWarehouses(ctx context.Context, query string) (int64, error)
	// GetListWarehouses select a list warehouse
	GetListWarehouses(ctx context.Context, query string, limit, skip int) ([]domain.Warehouse, error)
	// CountAuthorizedWarehouses count authorized warehouse
	CountAuthorizedWarehouses(ctx context.Context, userID int, query string) (int64, error)
	// GetAuthorizedWarehouses
	GetAuthorizedWarehouses(ctx context.Context, userID int, query string, limit, skip int) ([]domain.Warehouse, error)
	// GetUsedCapacityByID get used capacity of warehouse
	GetUsedCapacityByID(ctx context.Context, id int) (int64, error)
	// GetInventory get warehouse inventory by warehouse id
	GetInventory(ctx context.Context, id int) ([]domain.WarehouseItem, error)
	// UpdateWarehouse update a warehouse, only update non-zero fields by default
	UpdateWarehouse(ctx context.Context, warehouses *domain.Warehouse) (*domain.Warehouse, error)
	// DeleteWarehouse delete a warehouse
	DeleteWarehouse(ctx context.Context, id int) error
}

type IWarehouseService interface {
	// CreateWarehouse insert a new warehouse into the database
	CreateWarehouse(ctx context.Context, warehouses *domain.Warehouse) (*domain.Warehouse, error)
	// GetWarehouseByID select a warehouse by id
	GetWarehouseByID(ctx context.Context, id int) (*domain.Warehouse, error)
	// CountWarehouse count warehouse
	CountWarehouses(ctx context.Context, query string) (int64, error)
	// GetListWarehouses select a list warehouses
	GetListWarehouses(ctx context.Context, query string, limit, skip int) ([]domain.Warehouse, error)
	// CountAuthorizedWarehouses count authorized warehouse
	CountAuthorizedWarehouses(ctx context.Context, userID int, query string) (int64, error)
	// GetAuthorizedWarehouses
	GetAuthorizedWarehouses(ctx context.Context, userID int, query string, limit, skip int) ([]domain.Warehouse, error)
	// GetUsedCapacityByID get used capacity of warehouse
	GetUsedCapacityByID(ctx context.Context, id int) (int64, error)
	// GetInventory get warehouse inventory by warehouse id
	GetInventory(ctx context.Context, id int) ([]domain.WarehouseItem, error)
	// UpdateWarehouse update a warehouse, only update non-zero fields by default
	UpdateWarehouse(ctx context.Context, warehouses *domain.Warehouse) (*domain.Warehouse, error)
	// DeleteWarehouse delete a warehouse
	DeleteWarehouse(ctx context.Context, id int) error
}
