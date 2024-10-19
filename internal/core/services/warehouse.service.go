package services

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type warehouseService struct {
	repo ports.IWarehouseRepository
	file ports.IFileStorage
}

func NewWarehouseService(repo ports.IWarehouseRepository, fileStorage ports.IFileStorage) ports.IWarehouseService {
	return &warehouseService{
		repo: repo,
		file: fileStorage,
	}
}

func (ss *warehouseService) CreateWarehouse(ctx context.Context, warehouse *domain.Warehouse) (*domain.Warehouse, error) {
	err := ss.file.SavePermanentFile(warehouse.Image)
	if err != nil {
		if err == domain.ErrFileIsNotExist {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	created, err := ss.repo.CreateWarehouse(ctx, warehouse)
	if err != nil {
		_ = ss.file.DeleteFile(warehouse.Image)

		switch err {
		case domain.ErrConflictingData:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}
	ss.file.DeleteTempFile(warehouse.Image)

	return created, nil
}

func (ss *warehouseService) GetWarehouseByID(ctx context.Context, id int) (*domain.Warehouse, error) {
	store, err := ss.repo.GetWarehouseByID(ctx, id)
	if err != nil {
		switch err {
		case domain.ErrDataNotFound:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}

	return store, nil
}

func (ss *warehouseService) CountWarehouses(ctx context.Context, query string) (int64, error) {
	count, err := ss.repo.CountWarehouses(ctx, query)
	if err != nil {
		return 0, domain.ErrInternal
	}

	return count, nil
}

func (ss *warehouseService) GetListWarehouses(ctx context.Context, query string, limit, skip int) ([]domain.Warehouse, error) {
	list, err := ss.repo.GetListWarehouses(ctx, query, limit, skip)
	if err != nil {
		switch err {
		case domain.ErrDataNotFound:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}

	return list, nil
}

func (ss *warehouseService) CountAuthorizedWarehouses(ctx context.Context, userID int, query string) (int64, error) {
	count, err := ss.repo.CountAuthorizedWarehouses(ctx, userID, query)
	if err != nil {
		return 0, domain.ErrInternal
	}

	return count, nil
}

func (ss *warehouseService) GetAuthorizedWarehouses(ctx context.Context, userID int, query string, limit int, skip int) ([]domain.Warehouse, error) {
	list, err := ss.repo.GetAuthorizedWarehouses(ctx, userID, query, limit, skip)
	if err != nil {
		switch err {
		case domain.ErrDataNotFound:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}

	return list, nil
}

func (ss *warehouseService) GetUsedCapacityByID(ctx context.Context, id int) (int64, error) {
	usedCapacity, err := ss.repo.GetUsedCapacityByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return 0, err
		}
		return 0, domain.ErrInternal
	}

	return usedCapacity, nil
}

func (ss *warehouseService) UpdateWarehouse(ctx context.Context, warehouse *domain.Warehouse) (*domain.Warehouse, error) {
	current, err := ss.repo.GetWarehouseByID(ctx, warehouse.ID)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	isChangeImage := warehouse.Image != "" && warehouse.Image != current.Image
	if isChangeImage {
		err := ss.file.SavePermanentFile(warehouse.Image)
		if err != nil {
			if err == domain.ErrFileIsNotExist {
				return nil, err
			}
			return nil, domain.ErrInternal
		}
	}

	updated, err := ss.repo.UpdateWarehouse(ctx, warehouse)
	if err != nil {
		ss.file.DeleteFile(warehouse.Image)

		switch err {
		case domain.ErrNoUpdatedData:
			return nil, err
		case domain.ErrConflictingData:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}

	if isChangeImage {
		ss.file.DeleteFile(current.Image)
		ss.file.DeleteTempFile(updated.Image)
	}

	return updated, nil
}

func (ss *warehouseService) DeleteWarehouse(ctx context.Context, id int) error {
	err := ss.repo.DeleteWarehouse(ctx, id)
	if err != nil {
		switch err {
		case domain.ErrDataNotFound:
			return err
		default:
			return domain.ErrInternal
		}
	}
	return nil
}
