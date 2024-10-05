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

type keyRepository struct {
	db *mysqldb.MysqlDB
}

func NewKeyRepository(db *mysqldb.MysqlDB) ports.IKeyRepository {
	return &keyRepository{
		db: db,
	}
}

func (k *keyRepository) GetKey(ctx context.Context, id int) (string, error) {
	user := &schema.User{}

	err := k.db.Table("users").Select("key").Where("id = ?", id).Find(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", domain.ErrDataNotFound
		}
		return "", err
	}

	key := user.Key

	return key.String, nil
}

func (k *keyRepository) SetKey(ctx context.Context, id int, key string) error {
	result := k.db.Table("users").Where("id = ?", id).Update("key", key)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrNoUpdatedData
	}

	return nil
}

func (k *keyRepository) DelKey(ctx context.Context, id int) error {
	result := k.db.Table("users").Where("id = ?", id).Update("key", nil)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrNoUpdatedData
	}

	return nil
}
