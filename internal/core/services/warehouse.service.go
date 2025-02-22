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

func (w *warehouseService) CreateWarehouse(ctx context.Context, warehouse *domain.Warehouse) (*domain.Warehouse, error) {
	err := w.file.SavePermanentFile(warehouse.Image)
	if err != nil {
		if err == domain.ErrFileIsNotExist {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	created, err := w.repo.CreateWarehouse(ctx, warehouse)
	if err != nil {
		_ = w.file.DeleteFile(warehouse.Image)

		switch err {
		case domain.ErrConflictingData:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}
	w.file.DeleteTempFile(warehouse.Image)

	return created, nil
}

func (w *warehouseService) GetWarehouseByID(ctx context.Context, id int) (*domain.Warehouse, error) {
	store, err := w.repo.GetWarehouseByID(ctx, id)
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

func (w *warehouseService) CountWarehouses(ctx context.Context, query string) (int64, error) {
	count, err := w.repo.CountWarehouses(ctx, query)
	if err != nil {
		return 0, domain.ErrInternal
	}

	return count, nil
}

func (w *warehouseService) GetListWarehouses(ctx context.Context, query string, limit, skip int) ([]domain.Warehouse, error) {
	list, err := w.repo.GetListWarehouses(ctx, query, limit, skip)
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

func (w *warehouseService) CountAuthorizedWarehouses(ctx context.Context, userID int, query string) (int64, error) {
	count, err := w.repo.CountAuthorizedWarehouses(ctx, userID, query)
	if err != nil {
		return 0, domain.ErrInternal
	}

	return count, nil
}

func (w *warehouseService) GetAuthorizedWarehouses(ctx context.Context, userID int, query string, limit int, skip int) ([]domain.Warehouse, error) {
	list, err := w.repo.GetAuthorizedWarehouses(ctx, userID, query, limit, skip)
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

func (w *warehouseService) GetUsedCapacityByID(ctx context.Context, id int) (int64, error) {
	usedCapacity, err := w.repo.GetUsedCapacityByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return 0, err
		}
		return 0, domain.ErrInternal
	}

	return usedCapacity, nil
}

func (w *warehouseService) GetInventory(ctx context.Context, id int) ([]domain.WarehouseItem, error) {
	inventory, err := w.repo.GetInventory(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	return inventory, nil
}

func (w *warehouseService) UpdateWarehouse(ctx context.Context, warehouse *domain.Warehouse) (*domain.Warehouse, error) {
	current, err := w.repo.GetWarehouseByID(ctx, warehouse.ID)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	isChangeImage := warehouse.Image != "" && warehouse.Image != current.Image
	if isChangeImage {
		err := w.file.SavePermanentFile(warehouse.Image)
		if err != nil {
			if err == domain.ErrFileIsNotExist {
				return nil, err
			}
			return nil, domain.ErrInternal
		}
	}

	updated, err := w.repo.UpdateWarehouse(ctx, warehouse)
	if err != nil {
		if isChangeImage {
			_ = w.file.DeleteFile(warehouse.Image)
		}

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
		w.file.DeleteFile(current.Image)
		w.file.DeleteTempFile(updated.Image)
	}

	return updated, nil
}

func (w *warehouseService) DeleteWarehouse(ctx context.Context, id int) error {
	err := w.repo.DeleteWarehouse(ctx, id)
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
