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

func (ss *storehouseService) CreateStorehouse(ctx context.Context, storehouse *domain.Storehouse) {

}
