package repository

import (
	"context"
	"errors"

	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb/schema"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type storehouseRepository struct {
	db *mysqldb.MysqlDB
}

func NewStorehouseRepository(db *mysqldb.MysqlDB) *storehouseRepository {
	return &storehouseRepository{
		db: db,
	}
}

func (sr *storehouseRepository) Create(ctx context.Context, storehouses *domain.Storehouse) (*domain.Storehouse, error) {
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

func (sr *storehouseRepository) Update(ctx context.Context, storehouses *domain.Storehouse) (*domain.Storehouse, error) {
	updateData := &schema.Storehouse{
		ID:       storehouses.ID,
		Name:     storehouses.Name,
		Location: storehouses.Location,
		Capacity: storehouses.Capacity,
	}

	updatedData := &schema.Storehouse{}

	result := sr.db.WithContext(ctx).Clauses(clause.Returning{}).Model(updatedData).Where("id = ?", storehouses.ID).Updates(updateData)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, domain.ErrNoUpdatedData
	}

	return convertToDomainStorehouse(updateData), nil
}
