package services

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type accessControlService struct {
	repo ports.IAccessControlRepository
}

func NewAccessControlService(repo ports.IAccessControlRepository) ports.IAccessControlService {
	return &accessControlService{
		repo: repo,
	}
}

func (acs *accessControlService) HasAccess(ctx context.Context, storeHouseID int, userID int) error {
	err := acs.repo.HasAccess(ctx, storeHouseID, userID)
	if err != nil {
		switch err {
		case domain.ErrForbidden:
			return err
		default:
			return domain.ErrInternal
		}
	}

	return nil
}

func (acs *accessControlService) SetAccess(ctx context.Context, storeHouseID int, userID int) error {
	err := acs.repo.SetAccess(ctx, storeHouseID, userID)
	if err != nil {
		switch err {
		case domain.ErrConflictingData:
			return err
		default:
			return domain.ErrInternal
		}
	}

	return nil
}

func (acs *accessControlService) DelAccess(ctx context.Context, storeHouseID int, userID int) error {
	err := acs.repo.DelAccess(ctx, storeHouseID, userID)
	if err != nil {
		return domain.ErrInternal
	}

	return nil
}
