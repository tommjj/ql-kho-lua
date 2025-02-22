package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

type MockWarehouseRepository struct {
	mock.Mock
}

func (m *MockWarehouseRepository) CreateWarehouse(ctx context.Context, warehouse *domain.Warehouse) (*domain.Warehouse, error) {
	args := m.Called(ctx, warehouse)
	return args.Get(0).(*domain.Warehouse), args.Error(1)
}

func (m *MockWarehouseRepository) GetWarehouseByID(ctx context.Context, id int) (*domain.Warehouse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Warehouse), args.Error(1)
}

func (m *MockWarehouseRepository) CountWarehouses(ctx context.Context, query string) (int64, error) {
	args := m.Called(ctx, query)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockWarehouseRepository) GetListWarehouses(ctx context.Context, query string, limit, skip int) ([]domain.Warehouse, error) {
	args := m.Called(ctx, query, limit, skip)
	return args.Get(0).([]domain.Warehouse), args.Error(1)
}

func (m *MockWarehouseRepository) CountAuthorizedWarehouses(ctx context.Context, userID int, query string) (int64, error) {
	args := m.Called(ctx, userID, query)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockWarehouseRepository) GetAuthorizedWarehouses(ctx context.Context, userID int, query string, limit, skip int) ([]domain.Warehouse, error) {
	args := m.Called(ctx, userID, query, limit, skip)
	return args.Get(0).([]domain.Warehouse), args.Error(1)
}

func (m *MockWarehouseRepository) GetUsedCapacityByID(ctx context.Context, id int) (int64, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockWarehouseRepository) GetInventory(ctx context.Context, id int) ([]domain.WarehouseItem, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]domain.WarehouseItem), args.Error(1)
}

func (m *MockWarehouseRepository) UpdateWarehouse(ctx context.Context, warehouse *domain.Warehouse) (*domain.Warehouse, error) {
	args := m.Called(ctx, warehouse)
	return args.Get(0).(*domain.Warehouse), args.Error(1)
}

func (m *MockWarehouseRepository) DeleteWarehouse(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
