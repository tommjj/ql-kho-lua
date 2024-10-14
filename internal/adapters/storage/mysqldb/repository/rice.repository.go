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

	return convertToRice(createData), nil
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

	return convertToRice(rice), nil
}

func (rr *riceRepository) CountRice(ctx context.Context, query string) (int64, error) {
	var count int64
	var err error

	q := rr.db.WithContext(ctx).Table("rice").Where("deleted_at is NULL")

	trimQuery := strings.TrimSpace(query)
	if trimQuery != "" {
		q.Where("name LIKE ?", fmt.Sprintf("%%%v%%", trimQuery))
	}

	err = q.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (rr *riceRepository) GetListRice(ctx context.Context, query string, limit, skip int) ([]domain.Rice, error) {
	rice := []domain.Rice{}
	var err error

	q := rr.db.WithContext(ctx).Table("rice").
		Limit(limit).Offset((skip - 1) * limit).
		Order("id desc").Where("deleted_at is NULL")

	trimQuery := strings.TrimSpace(query)
	if trimQuery != "" {
		q.Where("name LIKE ?", fmt.Sprintf("%%%v%%", trimQuery))
	}

	err = q.Scan(&rice).Error
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

	result := rr.db.WithContext(ctx).Model(updated).Where("id = ?", rice.ID).Updates(updateData)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrConflictingData
		}
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, domain.ErrNoUpdatedData
	}

	return convertToRice(updated), nil
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
