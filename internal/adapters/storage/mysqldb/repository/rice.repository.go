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

type riceRepository struct {
	db *mysqldb.MysqlDB
}

func NewRiceRepository(db *mysqldb.MysqlDB) ports.IRiceRepository {
	return &riceRepository{
		db: db,
	}
}

func (rr *riceRepository) CreateRice(ctx context.Context, rice *domain.Rice) (*domain.Rice, error) {
	createData := &schema.Rice{
		Name: rice.Name,
	}

	err := rr.db.WithContext(ctx).Create(createData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return convertToDomainRice(createData), nil
}

func (rr *riceRepository) GetRiceByID(ctx context.Context, id int) (*domain.Rice, error) {
	rice := &schema.Rice{}

	err := rr.db.WithContext(ctx).Where("id = ?", id).First(rice).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return convertToDomainRice(rice), nil
}

func (rr *riceRepository) GetListRice(ctx context.Context, query string, skip, limit int) ([]domain.Rice, error) {
	rice := []domain.Rice{}
	var err error

	if strings.TrimSpace(query) == "" {
		err = rr.db.Table("rice").WithContext(ctx).Limit(limit).Offset((skip - 1) * limit).Scan(&rice).Error
	} else {
		err = rr.db.Table("rice").WithContext(ctx).Where("name LIKE ?", fmt.Sprintf("%%%v%%", query)).Limit(limit).Offset((skip - 1) * limit).Scan(&rice).Error
	}

	if err != nil {
		return nil, err
	}
	if len(rice) == 0 {
		return nil, domain.ErrDataNotFound
	}

	return rice, nil
}

func (rr *riceRepository) UpdateRice(ctx context.Context, rice *domain.Rice) (*domain.Rice, error) {
	updateData := &schema.Rice{
		ID:   rice.ID,
		Name: rice.Name,
	}

	updated := &schema.Rice{}

	result := rr.db.WithContext(ctx).Clauses(clause.Returning{}).Model(updated).Where("id = ?", rice.ID).Updates(updateData)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrConflictingData
		}
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, domain.ErrNoUpdatedData
	}

	return convertToDomainRice(updated), nil
}

func (rr *riceRepository) DeleteRice(ctx context.Context, id int) error {
	result := rr.db.WithContext(ctx).Where("id = ?", id).Delete(&schema.Rice{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrDataNotFound
	}

	return nil
}
