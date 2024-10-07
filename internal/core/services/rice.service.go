package services

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type riceService struct {
	repo ports.IRiceRepository
}

func NewRiceService(repo ports.IRiceRepository) ports.IRiceService {
	return &riceService{
		repo: repo,
	}
}

func (r *riceService) CreateRice(ctx context.Context, rice *domain.Rice) (*domain.Rice, error) {
	created, err := r.repo.CreateRice(ctx, rice)
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

func (r *riceService) GetRiceByID(ctx context.Context, id int) (*domain.Rice, error) {
	rice, err := r.repo.GetRiceByID(ctx, id)
	if err != nil {
		switch err {
		case domain.ErrDataNotFound:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}

	return rice, nil
}

func (r *riceService) GetListRice(ctx context.Context, query string, limit, skip int) ([]domain.Rice, error) {
	rice, err := r.repo.GetListRice(ctx, query, limit, skip)
	if err != nil {
		switch err {
		case domain.ErrDataNotFound:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}

	return rice, nil
}

func (r *riceService) UpdateRice(ctx context.Context, rice *domain.Rice) (*domain.Rice, error) {
	_, err := r.repo.GetRiceByID(ctx, rice.ID)
	if err != nil {
		switch err {
		case domain.ErrDataNotFound:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}

	updated, err := r.repo.UpdateRice(ctx, rice)
	if err != nil {
		switch err {
		case domain.ErrNoUpdatedData:
			return nil, err
		case domain.ErrConflictingData:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}

	return updated, nil
}

func (r *riceService) DeleteRice(ctx context.Context, id int) error {
	err := r.repo.DeleteRice(ctx, id)
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
