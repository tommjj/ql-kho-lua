package services

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type storehouseService struct {
	repo ports.IStorehouseRepository
	file ports.IFileStorage
}

func NewStorehouseService(repo ports.IStorehouseRepository, fileStorage ports.IFileStorage) ports.IStorehouseService {
	return &storehouseService{
		repo: repo,
		file: fileStorage,
	}
}

func (ss *storehouseService) CreateStorehouse(ctx context.Context, storehouse *domain.Storehouse) (*domain.Storehouse, error) {
	err := ss.file.SavePermanentFile(storehouse.Image)
	if err != nil {
		if err == domain.ErrFileIsNotExist {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	created, err := ss.repo.CreateStorehouse(ctx, storehouse)
	if err != nil {
		_ = ss.file.DeleteFile(storehouse.Image)

		switch err {
		case domain.ErrConflictingData:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}
	ss.file.DeleteTempFile(storehouse.Image)

	return created, nil
}

func (ss *storehouseService) GetStorehouseByID(ctx context.Context, id int) (*domain.Storehouse, error) {
	store, err := ss.repo.GetStorehouseByID(ctx, id)
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

func (ss *storehouseService) CountStorehouses(ctx context.Context, query string) (int64, error) {
	count, err := ss.repo.CountStorehouses(ctx, query)
	if err != nil {
		return 0, domain.ErrInternal
	}

	return count, nil
}

func (ss *storehouseService) GetListStorehouses(ctx context.Context, query string, limit, skip int) ([]domain.Storehouse, error) {
	list, err := ss.repo.GetListStorehouses(ctx, query, limit, skip)
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

func (ss *storehouseService) CountAuthorizedStorehouses(ctx context.Context, userID int, query string) (int64, error) {
	count, err := ss.repo.CountAuthorizedStorehouses(ctx, userID, query)
	if err != nil {
		return 0, domain.ErrInternal
	}

	return count, nil
}

func (ss *storehouseService) GetAuthorizedStorehouses(ctx context.Context, userID int, query string, limit int, skip int) ([]domain.Storehouse, error) {
	list, err := ss.repo.GetAuthorizedStorehouses(ctx, userID, query, limit, skip)
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

func (ss *storehouseService) GetStorehouseUsedCapacityByID(ctx context.Context, id int) (float64, error) {
	usedCapacity, err := ss.repo.GetStorehouseUsedCapacityByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return 0, err
		}
		return 0, domain.ErrInternal
	}

	return usedCapacity, nil
}

func (ss *storehouseService) UpdateStorehouse(ctx context.Context, storehouse *domain.Storehouse) (*domain.Storehouse, error) {
	current, err := ss.repo.GetStorehouseByID(ctx, storehouse.ID)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	isChangeImage := storehouse.Image != "" && storehouse.Image != current.Image
	if isChangeImage {
		err := ss.file.SavePermanentFile(storehouse.Image)
		if err != nil {
			if err == domain.ErrFileIsNotExist {
				return nil, err
			}
			return nil, domain.ErrInternal
		}
	}

	updated, err := ss.repo.UpdateStorehouse(ctx, storehouse)
	if err != nil {
		ss.file.DeleteFile(storehouse.Image)

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

func (ss *storehouseService) DeleteStorehouse(ctx context.Context, id int) error {
	err := ss.repo.DeleteStorehouse(ctx, id)
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
