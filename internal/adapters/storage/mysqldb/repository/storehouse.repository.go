package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb/schema"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type storehouseRepository struct {
	db *mysqldb.MysqlDB
}

func NewStorehouseRepository(db *mysqldb.MysqlDB) ports.IStorehouseRepository {
	return &storehouseRepository{
		db: db,
	}
}

func (sr *storehouseRepository) CreateStorehouse(ctx context.Context, storehouses *domain.Storehouse) (*domain.Storehouse, error) {
	createData := &schema.Storehouse{
		Name:     storehouses.Name,
		Location: storehouses.Location,
		Capacity: storehouses.Capacity,
	}

	err := sr.db.WithContext(ctx).Create(createData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return convertToDomainStorehouse(createData), nil
}

func (sr *storehouseRepository) GetStorehouseByID(ctx context.Context, id int) (*domain.Storehouse, error) {
	store := &schema.Storehouse{}

	err := sr.db.WithContext(ctx).Where("id = ?", id).First(store).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return convertToDomainStorehouse(store), nil
}

func (sr *storehouseRepository) GetListStorehouses(ctx context.Context, query string, limit, skip int) ([]domain.Storehouse, error) {
	stores := []domain.Storehouse{}
	var err error

	sql := sr.db.Table("storehouses").WithContext(ctx).
		Select("id", "name", "location", "capacity").
		Limit(limit).Offset((skip - 1) * limit)

	trimQuery := strings.TrimSpace(query)
	if trimQuery == "" {
		err = sql.Scan(&stores).Error
	} else {
		err = sql.Where("name LIKE ?", fmt.Sprintf("%%%v%%", trimQuery)).Scan(&stores).Error
	}

	if err != nil {
		return nil, err
	}
	if len(stores) == 0 {
		return nil, domain.ErrDataNotFound
	}

	return stores, nil
}

func (sr *storehouseRepository) UpdateStorehouse(ctx context.Context, storehouses *domain.Storehouse) (*domain.Storehouse, error) {
	updatedData := &schema.Storehouse{}

	result := sr.db.WithContext(ctx).Clauses(clause.Returning{}).
		Model(updatedData).Where("id = ?", storehouses.ID).
		Updates(&schema.Storehouse{
			Name:     storehouses.Name,
			Location: storehouses.Location,
			Capacity: storehouses.Capacity,
		})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, domain.ErrNoUpdatedData
	}

	return convertToDomainStorehouse(updatedData), nil
}

func (sr *storehouseRepository) DeleteStorehouse(ctx context.Context, id int) error {
	result := sr.db.WithContext(ctx).Where("id = ?", id).Delete(&schema.Storehouse{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrDataNotFound
	}

	return nil
}
