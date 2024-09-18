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

func (ar *accessControlRepository) HasAccess(ctx context.Context, storeHouseID int, userID int) error {
	result := struct {
		StorehouseID int
		UserID       int
	}{UserID: -1}

	err := ar.db.WithContext(ctx).
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

func (ar *accessControlRepository) SetAccess(ctx context.Context, storeHouseID int, userID int) error {
	err := ar.db.Model(&schema.User{ID: userID}).Association("AuthorizedStorehouses").Append(&schema.Storehouse{ID: storeHouseID})

	if err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return domain.ErrConflictingData
		}
		return err
	}
	return nil
}

func (ar *accessControlRepository) DelAccess(ctx context.Context, storeHouseID int, userID int) error {
	err := ar.db.Model(&schema.User{ID: userID}).Association("AuthorizedStorehouses").Delete(&schema.Storehouse{ID: storeHouseID})
	if err != nil {
		return err
	}
	return nil
}

// func (ar *accessControlRepository) GetAuthorizedStorehouses(ctx context.Context, userID int) ([]domain.Storehouse, error) {
// 	list := []schema.Storehouse{}

// 	err := ar.db.Joins("LEFT JOIN authorized on authorized.storehouse_id = storehouses.id").Where("authorized.user_id = ?", userID).Find(&list).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(list) == 0 {
// 		return nil, domain.ErrDataNotFound
// 	}

// 	return nil, nil
// }
