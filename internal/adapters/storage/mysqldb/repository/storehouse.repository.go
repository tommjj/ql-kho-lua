package repository

import (
	"context"
	"database/sql"
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

	return convertToStorehouse(createData), nil
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

	return convertToStorehouse(store), nil
}

func (sr *storehouseRepository) GetListStorehouses(ctx context.Context, query string, limit, skip int) ([]domain.Storehouse, error) {
	stores := []domain.Storehouse{}
	var err error

	sql := sr.db.WithContext(ctx).Table("storehouses").
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

func (ar *storehouseRepository) GetAuthorizedStorehouses(ctx context.Context, userID int, query string, limit, skip int) ([]domain.Storehouse, error) {
	list := []schema.Storehouse{}
	var err error

	trimQuery := strings.TrimSpace(query)

	if trimQuery != "" {
		err = ar.db.WithContext(ctx).Joins("LEFT JOIN authorized on authorized.storehouse_id = storehouses.id").
			Where("authorized.user_id = ? AND name LIKE ?", userID, fmt.Sprintf("%%%v%%", trimQuery)).
			Limit(limit).Offset((skip - 1) * limit).Find(&list).Error
	} else {
		err = ar.db.WithContext(ctx).Joins("LEFT JOIN authorized on authorized.storehouse_id = storehouses.id").
			Where("authorized.user_id = ?", userID).Limit(limit).Offset((skip - 1) * limit).Find(&list).Error
	}

	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, domain.ErrDataNotFound
	}

	storehouse := make([]domain.Storehouse, 0, len(list))

	for _, v := range list {
		storehouse = append(storehouse, *convertToStorehouse(&v))
	}

	return storehouse, nil
}

func (sr *storehouseRepository) GetStorehouseUsedCapacityByID(ctx context.Context, id int) (float64, error) {
	err := sr.db.WithContext(ctx).First(&schema.Storehouse{ID: id}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, domain.ErrDataNotFound
		}
		return 0, err
	}

	total := struct {
		Total float64
	}{}

	err = sr.db.WithContext(ctx).
		Raw(`SELECT (SUM(IF(t.type = "i", t.total, 0)) - SUM(IF(t.type = "e", t.total, 0))) as "total"
			FROM 
				(SELECT SUM( export_invoice_details.quantity) as total, "e" as type
    			FROM export_invoices INNER JOIN export_invoice_details on export_invoices.id = export_invoice_details.invoice_id 
    			WHERE export_invoices.storehouse_id = @id 
    			UNION ALL
    			SELECT SUM( import_invoice_details.quantity) as total, "i" as type 
    			FROM import_invoices INNER JOIN import_invoice_details on import_invoices.id = import_invoice_details.invoice_id 
    			WHERE import_invoices.storehouse_id = @id) as t`, sql.Named("id", id)).Scan(&total).Error
	if err != nil {
		return 0, err
	}

	return total.Total, nil
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
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, domain.ErrNoUpdatedData
	}

	return convertToStorehouse(updatedData), nil
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
