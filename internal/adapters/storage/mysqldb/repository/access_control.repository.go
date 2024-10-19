package repository

import (
	"context"
	"errors"

	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb/schema"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
	"gorm.io/gorm"
)

type accessControlRepository struct {
	db *mysqldb.MysqlDB
}

func NewAccessControlRepository(db *mysqldb.MysqlDB) ports.IAccessControlRepository {
	return &accessControlRepository{
		db: db,
	}
}

func (ar *accessControlRepository) HasAccess(ctx context.Context, warehouseID int, userID int) error {
	result := struct {
		warehouseID int
		UserID      int
	}{UserID: -1}

	err := ar.db.WithContext(ctx).
		Raw("SELECT * FROM authorized WHERE warehouse_id = ? AND user_id = ?", warehouseID, userID).
		Scan(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrDataNotFound
		}
		return err
	}

	if result.UserID != userID {
		return domain.ErrForbidden
	}
	return nil

}

func (ar *accessControlRepository) SetAccess(ctx context.Context, warehouseID int, userID int) error {
	err := ar.db.WithContext(ctx).Model(&schema.User{ID: userID}).Association("AuthorizedWarehouses").Append(&schema.Warehouse{ID: warehouseID})

	if err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return domain.ErrConflictingData
		}
		return err
	}
	return nil
}

func (ar *accessControlRepository) DelAccess(ctx context.Context, warehouseID int, userID int) error {
	err := ar.db.WithContext(ctx).Model(&schema.User{ID: userID}).Association("AuthorizedWarehouses").Delete(&schema.Warehouse{ID: warehouseID})
	if err != nil {
		return err
	}
	return nil
}
