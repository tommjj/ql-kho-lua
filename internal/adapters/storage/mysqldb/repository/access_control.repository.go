package repository

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

type accessControlRepository struct {
	db *mysqldb.MysqlDB
}

func NewAccessControlRepository(db *mysqldb.MysqlDB) *accessControlRepository {
	return &accessControlRepository{
		db: db,
	}
}

func (r *accessControlRepository) HasAccess(ctx context.Context, storeHouseID int, userID int) error {
	result := struct {
		StorehouseID int
		UserID       int
	}{
		UserID: -1,
	}

	err := r.db.WithContext(ctx).
		Raw("SELECT * FROM authorized WHERE storehouse_id = ? AND user_id = ?", storeHouseID, userID).
		Scan(&result).Error
	if err != nil {
		return err
	}

	if result.UserID != userID {
		return domain.ErrForbidden
	}
	return nil
}
