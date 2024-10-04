package services

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type storehouseService struct {
	repo ports.IStorehouseRepository
}

func NewStorehouseService(repo ports.IStorehouseRepository) *storehouseService {
	return &storehouseService{
		repo: repo,
	}
}

func (ss *storehouseService) CreateStorehouse(ctx context.Context, storehouse *domain.Storehouse) (*domain.Storehouse, error) {
	created, err := ss.repo.CreateStorehouse(ctx, storehouse)
	if err != nil {
		switch err {
		case domain.ErrConflictingData:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}

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

func (ss *storehouseService) GetAuthorizedStorehouses(ctx context.Context, userID int) ([]domain.Storehouse, error) {
	list, err := ss.repo.GetAuthorizedStorehouses(ctx, userID, "", 1, 1)
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

//func (ss *storehouseService)

func (ss *storehouseService) UpdateStorehouse(ctx context.Context, storehouse *domain.Storehouse) (*domain.Storehouse, error) {
	created, err := ss.repo.UpdateStorehouse(ctx, storehouse)
	if err != nil {
		switch err {
		case domain.ErrNoUpdatedData:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}

	return created, nil
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
