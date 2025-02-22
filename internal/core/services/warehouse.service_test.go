package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
	mockRepo "github.com/tommjj/ql-kho-lua/internal/core/services/mock"
)

func TestWarehouseServiceImplements(t *testing.T) {
	assert.Implements(t, (*ports.IWarehouseService)(nil), new(warehouseService))
}

func TestCreateWarehouse_Success(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)
	fileStorage := new(mockRepo.MockFileStorage)

	warehouseRepo.On("CreateWarehouse", mock.Anything, mock.Anything).Return(&domain.Warehouse{}, nil)
	fileStorage.On("SavePermanentFile", mock.Anything).Return(nil)
	fileStorage.On("DeleteTempFile", mock.Anything).Return(nil)

	service := NewWarehouseService(warehouseRepo, fileStorage)
	_, err := service.CreateWarehouse(context.TODO(), &domain.Warehouse{})
	assert.Nil(t, err)

	warehouseRepo.AssertExpectations(t)
}

func TestCreateWarehouse_FailConflicting(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)
	fileStorage := new(mockRepo.MockFileStorage)

	warehouseRepo.On("CreateWarehouse", mock.Anything, mock.Anything).Return(nil, domain.ErrConflictingData)
	fileStorage.On("SavePermanentFile", mock.Anything).Return(nil)
	fileStorage.On("DeleteTempFile", mock.Anything).Return(nil)
	fileStorage.On("DeleteFile", mock.Anything).Return(nil)

	service := NewWarehouseService(warehouseRepo, fileStorage)
	_, err := service.CreateWarehouse(context.TODO(), &domain.Warehouse{})
	assert.Equal(t, domain.ErrConflictingData, err)

	warehouseRepo.AssertExpectations(t)
}

func TestCreateWarehouse_FailFileIsNotExist(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)
	fileStorage := new(mockRepo.MockFileStorage)

	warehouseRepo.On("CreateWarehouse", mock.Anything, mock.Anything).Return(&domain.Warehouse{}, nil)
	fileStorage.On("SavePermanentFile", mock.Anything).Return(domain.ErrFileIsNotExist)
	fileStorage.On("DeleteTempFile", mock.Anything).Return(nil)
	fileStorage.On("DeleteFile", mock.Anything).Return(nil)

	service := NewWarehouseService(warehouseRepo, fileStorage)
	_, err := service.CreateWarehouse(context.TODO(), &domain.Warehouse{})
	assert.Equal(t, domain.ErrFileIsNotExist, err)
}

func TestCreateWarehouse_FailSaveFileUnknownErr(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)
	fileStorage := new(mockRepo.MockFileStorage)

	warehouseRepo.On("CreateWarehouse", mock.Anything, mock.Anything).Return(nil, domain.ErrInternal)
	fileStorage.On("SavePermanentFile", mock.Anything).Return(domain.ErrInternal)
	fileStorage.On("DeleteTempFile", mock.Anything).Return(nil)
	fileStorage.On("DeleteFile", mock.Anything).Return(nil)

	service := NewWarehouseService(warehouseRepo, fileStorage)
	_, err := service.CreateWarehouse(context.TODO(), &domain.Warehouse{})
	assert.Equal(t, domain.ErrInternal, err)
}

func TestCreateWarehouse_FailUnknownErr(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)
	fileStorage := new(mockRepo.MockFileStorage)

	warehouseRepo.On("CreateWarehouse", mock.Anything, mock.Anything).Return(nil, domain.ErrInternal)
	fileStorage.On("SavePermanentFile", mock.Anything).Return(nil)
	fileStorage.On("DeleteTempFile", mock.Anything).Return(nil)
	fileStorage.On("DeleteFile", mock.Anything).Return(nil)

	service := NewWarehouseService(warehouseRepo, fileStorage)
	_, err := service.CreateWarehouse(context.TODO(), &domain.Warehouse{})
	assert.Equal(t, domain.ErrInternal, err)
}

func TestGetWarehouseByID_Success(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("GetWarehouseByID", mock.Anything, 1).Return(&domain.Warehouse{
		ID:       1,
		Name:     "Warehouse",
		Location: "1, 1",
		Capacity: 100,
		Image:    "image.jpg",
	}, nil)

	service := NewWarehouseService(warehouseRepo, nil)
	warehouse, err := service.GetWarehouseByID(context.TODO(), 1)
	assert.Nil(t, err)
	assert.Equal(t, warehouse.ID, 1)
	assert.Equal(t, warehouse.Name, "Warehouse")
	assert.Equal(t, warehouse.Location, "1, 1")
	assert.Equal(t, warehouse.Capacity, 100)
	assert.Equal(t, warehouse.Image, "image.jpg")

	warehouseRepo.AssertExpectations(t)
}

func TestGetWarehouseByID_FailNotFound(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("GetWarehouseByID", mock.Anything, 1).Return(nil, domain.ErrDataNotFound)

	service := NewWarehouseService(warehouseRepo, nil)
	_, err := service.GetWarehouseByID(context.TODO(), 1)
	assert.Equal(t, domain.ErrDataNotFound, err)

	warehouseRepo.AssertExpectations(t)
}

func TestGetWarehouseByID_FailUnknown(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("GetWarehouseByID", mock.Anything, 1).Return(nil, domain.ErrInternal)

	service := NewWarehouseService(warehouseRepo, nil)
	_, err := service.GetWarehouseByID(context.TODO(), 1)
	assert.Equal(t, domain.ErrInternal, err)

	warehouseRepo.AssertExpectations(t)
}

func TestCountWarehouses_Success(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("CountWarehouses", mock.Anything, "query").Return(int64(10), nil)

	service := NewWarehouseService(warehouseRepo, nil)
	count, err := service.CountWarehouses(context.TODO(), "query")
	assert.Nil(t, err)
	assert.Equal(t, count, int64(10))

	warehouseRepo.AssertExpectations(t)
}

func TestCountWarehouses_SuccessNoQuery(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("CountWarehouses", mock.Anything, "").Return(int64(10), nil)

	service := NewWarehouseService(warehouseRepo, nil)
	count, err := service.CountWarehouses(context.TODO(), "")
	assert.Nil(t, err)
	assert.Equal(t, count, int64(10))

	warehouseRepo.AssertExpectations(t)
}

func TestCountWarehouses_FailUnknown(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("CountWarehouses", mock.Anything, "query").Return(int64(0), errors.New("unknown error"))

	service := NewWarehouseService(warehouseRepo, nil)
	_, err := service.CountWarehouses(context.TODO(), "query")
	assert.Equal(t, domain.ErrInternal, err)

	warehouseRepo.AssertExpectations(t)
}

func TestGetListWarehouses_Success(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("GetListWarehouses", mock.Anything, "query", 10, 0).Return([]domain.Warehouse{
		{
			ID:       1,
			Name:     "Warehouse",
			Location: "1, 1",
			Capacity: 100,
			Image:    "image.jpg",
		},
	}, nil)

	service := NewWarehouseService(warehouseRepo, nil)
	warehouses, err := service.GetListWarehouses(context.TODO(), "query", 10, 0)
	assert.Nil(t, err)
	assert.Len(t, warehouses, 1)
	assert.Equal(t, warehouses[0].ID, 1)
	assert.Equal(t, warehouses[0].Name, "Warehouse")
	assert.Equal(t, warehouses[0].Location, "1, 1")
	assert.Equal(t, warehouses[0].Capacity, 100)
	assert.Equal(t, warehouses[0].Image, "image.jpg")

	warehouseRepo.AssertExpectations(t)
}

func TestGetListWarehouses_FailNotFound(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("GetListWarehouses", mock.Anything, "query", 10, 0).Return(nil, domain.ErrDataNotFound)

	service := NewWarehouseService(warehouseRepo, nil)
	_, err := service.GetListWarehouses(context.TODO(), "query", 10, 0)
	assert.Equal(t, domain.ErrDataNotFound, err)

	warehouseRepo.AssertExpectations(t)
}

func TestGetListWarehouses_FailUnknown(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("GetListWarehouses", mock.Anything, "query", 10, 0).Return(nil, errors.New("unknown error"))

	service := NewWarehouseService(warehouseRepo, nil)
	_, err := service.GetListWarehouses(context.TODO(), "query", 10, 0)
	assert.Equal(t, domain.ErrInternal, err)

	warehouseRepo.AssertExpectations(t)
}

func TestGetAuthorizedWarehouses_Success(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("GetAuthorizedWarehouses", mock.Anything, 1, "query", 10, 0).Return([]domain.Warehouse{
		{
			ID:       1,
			Name:     "Warehouse",
			Location: "1, 1",
			Capacity: 100,
			Image:    "image.jpg",
		},
	}, nil)

	service := NewWarehouseService(warehouseRepo, nil)
	warehouses, err := service.GetAuthorizedWarehouses(context.TODO(), 1, "query", 10, 0)
	assert.Nil(t, err)
	assert.Len(t, warehouses, 1)
	assert.Equal(t, warehouses[0].ID, 1)
	assert.Equal(t, warehouses[0].Name, "Warehouse")
	assert.Equal(t, warehouses[0].Location, "1, 1")
	assert.Equal(t, warehouses[0].Capacity, 100)
	assert.Equal(t, warehouses[0].Image, "image.jpg")

	warehouseRepo.AssertExpectations(t)
}

func TestGetAuthorizedWarehouses_FailNotFound(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("GetAuthorizedWarehouses", mock.Anything, 1, "query", 10, 0).Return(nil, domain.ErrDataNotFound)

	service := NewWarehouseService(warehouseRepo, nil)
	_, err := service.GetAuthorizedWarehouses(context.TODO(), 1, "query", 10, 0)
	assert.Equal(t, domain.ErrDataNotFound, err)

	warehouseRepo.AssertExpectations(t)
}

func TestGetAuthorizedWarehouses_FailUnknown(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("GetAuthorizedWarehouses", mock.Anything, 1, "query", 10, 0).Return(nil, errors.New("unknown error"))

	service := NewWarehouseService(warehouseRepo, nil)
	_, err := service.GetAuthorizedWarehouses(context.TODO(), 1, "query", 10, 0)
	assert.Equal(t, domain.ErrInternal, err)

	warehouseRepo.AssertExpectations(t)
}

func TestCountAuthorizedWarehouses_Success(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("CountAuthorizedWarehouses", mock.Anything, 1, "query").Return(int64(10), nil)

	service := NewWarehouseService(warehouseRepo, nil)
	count, err := service.CountAuthorizedWarehouses(context.TODO(), 1, "query")
	assert.Nil(t, err)
	assert.Equal(t, count, int64(10))

	warehouseRepo.AssertExpectations(t)
}

func TestCountAuthorizedWarehouses_SuccessNoQuery(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("CountAuthorizedWarehouses", mock.Anything, 1, "").Return(int64(10), nil)

	service := NewWarehouseService(warehouseRepo, nil)
	count, err := service.CountAuthorizedWarehouses(context.TODO(), 1, "")
	assert.Nil(t, err)
	assert.Equal(t, count, int64(10))

	warehouseRepo.AssertExpectations(t)
}

func TestCountAuthorizedWarehouses_FailUnknown(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("CountAuthorizedWarehouses", mock.Anything, 1, "query").Return(int64(0), errors.New("unknown error"))

	service := NewWarehouseService(warehouseRepo, nil)
	_, err := service.CountAuthorizedWarehouses(context.TODO(), 1, "query")
	assert.Equal(t, domain.ErrInternal, err)

	warehouseRepo.AssertExpectations(t)
}

func TestUpdateWarehouse_Success(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)
	fileStorage := new(mockRepo.MockFileStorage)

	warehouseRepo.On("GetWarehouseByID", mock.Anything, 1).Return(&domain.Warehouse{}, nil)
	warehouseRepo.On("UpdateWarehouse", mock.Anything, mock.Anything).Return(&domain.Warehouse{}, nil)
	fileStorage.On("SavePermanentFile", mock.Anything).Return(nil)
	fileStorage.On("DeleteTempFile", mock.Anything).Return(nil)

	service := NewWarehouseService(warehouseRepo, fileStorage)
	_, err := service.UpdateWarehouse(context.TODO(), &domain.Warehouse{ID: 1})
	assert.Nil(t, err)

	warehouseRepo.AssertExpectations(t)
}

func TestUpdateWarehouse_FailNoUpdatedData(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)
	fileStorage := new(mockRepo.MockFileStorage)

	warehouseRepo.On("GetWarehouseByID", mock.Anything, 1).Return(&domain.Warehouse{
		ID:       1,
		Name:     "Warehouse",
		Location: "1, 1",
		Capacity: 100,
		Image:    "image.jpg",
	}, nil)
	warehouseRepo.On("UpdateWarehouse", mock.Anything, mock.Anything).Return(nil, domain.ErrNoUpdatedData)

	service := NewWarehouseService(warehouseRepo, fileStorage)
	_, err := service.UpdateWarehouse(context.TODO(), &domain.Warehouse{
		ID:       1,
		Name:     "Warehouse",
		Location: "1, 1",
		Capacity: 100,
		Image:    "image.jpg",
	})
	assert.Equal(t, domain.ErrNoUpdatedData, err)

	warehouseRepo.AssertExpectations(t)
}

func TestUpdateWarehouse_FailFileIsNotExist(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)
	fileStorage := new(mockRepo.MockFileStorage)

	warehouseRepo.On("GetWarehouseByID", mock.Anything, 1).Return(&domain.Warehouse{
		ID:    1,
		Image: "image.jpg",
	}, nil)

	fileStorage.On("SavePermanentFile", mock.Anything).Return(domain.ErrFileIsNotExist)

	service := NewWarehouseService(warehouseRepo, fileStorage)
	_, err := service.UpdateWarehouse(context.TODO(), &domain.Warehouse{ID: 1, Image: "image02.jpg"})
	assert.Equal(t, domain.ErrFileIsNotExist, err)

	warehouseRepo.AssertExpectations(t)
}

func TestUpdateWarehouse_FailNotFound(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)
	fileStorage := new(mockRepo.MockFileStorage)

	warehouseRepo.On("GetWarehouseByID", mock.Anything, 1).Return(nil, domain.ErrDataNotFound)

	service := NewWarehouseService(warehouseRepo, fileStorage)
	_, err := service.UpdateWarehouse(context.TODO(), &domain.Warehouse{ID: 1})
	assert.Equal(t, domain.ErrDataNotFound, err)

	warehouseRepo.AssertExpectations(t)
}

func TestUpdateWarehouse_FailConflicting(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)
	fileStorage := new(mockRepo.MockFileStorage)

	warehouseRepo.On("GetWarehouseByID", mock.Anything, 1).Return(&domain.Warehouse{}, nil)
	warehouseRepo.On("UpdateWarehouse", mock.Anything, mock.Anything).Return(nil, domain.ErrConflictingData)

	service := NewWarehouseService(warehouseRepo, fileStorage)
	_, err := service.UpdateWarehouse(context.TODO(), &domain.Warehouse{ID: 1})
	assert.Equal(t, domain.ErrConflictingData, err)

	warehouseRepo.AssertExpectations(t)
}

func TestUpdateWarehouse_FailUnknown(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)
	fileStorage := new(mockRepo.MockFileStorage)

	warehouseRepo.On("GetWarehouseByID", mock.Anything, 1).Return(&domain.Warehouse{}, nil)
	warehouseRepo.On("UpdateWarehouse", mock.Anything, mock.Anything).Return(nil, errors.New("unknown error"))

	service := NewWarehouseService(warehouseRepo, fileStorage)
	_, err := service.UpdateWarehouse(context.TODO(), &domain.Warehouse{ID: 1})
	assert.Equal(t, domain.ErrInternal, err)

	warehouseRepo.AssertExpectations(t)
}

func TestDeleteWarehouse_Success(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)
	fileStorage := new(mockRepo.MockFileStorage)

	warehouseRepo.On("DeleteWarehouse", mock.Anything, 1).Return(nil)
	fileStorage.On("DeleteFile", mock.Anything).Return(nil)

	service := NewWarehouseService(warehouseRepo, fileStorage)
	err := service.DeleteWarehouse(context.TODO(), 1)
	assert.Nil(t, err)

	warehouseRepo.AssertExpectations(t)
}

func TestDeleteWarehouse_FailNotFound(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)
	fileStorage := new(mockRepo.MockFileStorage)

	warehouseRepo.On("DeleteWarehouse", mock.Anything, 1).Return(domain.ErrDataNotFound)

	service := NewWarehouseService(warehouseRepo, fileStorage)
	err := service.DeleteWarehouse(context.TODO(), 1)
	assert.Equal(t, domain.ErrDataNotFound, err)

	warehouseRepo.AssertExpectations(t)
}

func TestDeleteWarehouse_FailUnknown(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)
	fileStorage := new(mockRepo.MockFileStorage)

	warehouseRepo.On("DeleteWarehouse", mock.Anything, 1).Return(errors.New("unknown error"))

	service := NewWarehouseService(warehouseRepo, fileStorage)
	err := service.DeleteWarehouse(context.TODO(), 1)
	assert.Equal(t, domain.ErrInternal, err)

	warehouseRepo.AssertExpectations(t)
}

func TestGetInventory_Success(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("GetInventory", mock.Anything, 1).Return([]domain.WarehouseItem{
		{
			RiceID: 1,
			Rice: &domain.Rice{
				ID:   1,
				Name: "Rice",
			},
			Quantity: 100,
		},
	}, nil)

	service := NewWarehouseService(warehouseRepo, nil)
	inventory, err := service.GetInventory(context.TODO(), 1)
	assert.Nil(t, err)
	assert.Len(t, inventory, 1)
	assert.Equal(t, inventory[0].RiceID, 1)
	assert.Equal(t, inventory[0].Rice.Name, "Rice")
	assert.Equal(t, inventory[0].Quantity, 100)

	warehouseRepo.AssertExpectations(t)
}

func TestGetInventory_FailNotFound(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("GetInventory", mock.Anything, 1).Return(nil, domain.ErrDataNotFound)

	service := NewWarehouseService(warehouseRepo, nil)
	_, err := service.GetInventory(context.TODO(), 1)
	assert.Equal(t, domain.ErrDataNotFound, err)

	warehouseRepo.AssertExpectations(t)
}

func TestGetInventory_FailUnknown(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("GetInventory", mock.Anything, 1).Return(nil, errors.New("unknown error"))

	service := NewWarehouseService(warehouseRepo, nil)
	_, err := service.GetInventory(context.TODO(), 1)
	assert.Equal(t, domain.ErrInternal, err)

	warehouseRepo.AssertExpectations(t)
}

func TestGetUsedCapacityByID_Success(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("GetUsedCapacityByID", mock.Anything, 1).Return(int64(100), nil)

	service := NewWarehouseService(warehouseRepo, nil)
	capacity, err := service.GetUsedCapacityByID(context.TODO(), 1)
	assert.Nil(t, err)
	assert.Equal(t, capacity, int64(100))

	warehouseRepo.AssertExpectations(t)
}

func TestGetUsedCapacityByID_FailNotFound(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("GetUsedCapacityByID", mock.Anything, 1).Return(int64(0), domain.ErrDataNotFound)

	service := NewWarehouseService(warehouseRepo, nil)
	_, err := service.GetUsedCapacityByID(context.TODO(), 1)
	assert.Equal(t, domain.ErrDataNotFound, err)

	warehouseRepo.AssertExpectations(t)
}

func TestGetUsedCapacityByID_FailUnknown(t *testing.T) {
	warehouseRepo := new(mockRepo.MockWarehouseRepository)

	warehouseRepo.On("GetUsedCapacityByID", mock.Anything, 1).Return(int64(0), errors.New("unknown error"))

	service := NewWarehouseService(warehouseRepo, nil)
	_, err := service.GetUsedCapacityByID(context.TODO(), 1)
	assert.Equal(t, domain.ErrInternal, err)

	warehouseRepo.AssertExpectations(t)
}
