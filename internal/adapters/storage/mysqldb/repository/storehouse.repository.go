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
		Image:    storehouses.Image,
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

func (sr *storehouseRepository) CountStorehouses(ctx context.Context, query string) (int64, error) {
	var count int64
	var err error

	q := sr.db.WithContext(ctx).Table("storehouses").Where("deleted_at is NULL")

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

func (sr *storehouseRepository) GetListStorehouses(ctx context.Context, query string, limit, skip int) ([]domain.Storehouse, error) {
	stores := []domain.Storehouse{}
	var err error
	var rows *sql.Rows

	sql := sr.db.WithContext(ctx).Table("storehouses").
		Select("id", "name", "location", "capacity", "image").
		Limit(limit).Offset((skip - 1) * limit).Order("id desc").Where("deleted_at is NULL")

	trimQuery := strings.TrimSpace(query)
	if trimQuery != "" {
		sql.Where("name LIKE ?", fmt.Sprintf("%%%v%%", trimQuery))
	}

	rows, err = sql.Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		store := domain.Storehouse{}
		rows.Scan(
			&store.ID,
			&store.Name,
			&store.Location,
			&store.Capacity,
			&store.Image,
		)

		stores = append(stores, store)
	}

	if len(stores) == 0 {
		return nil, domain.ErrDataNotFound
	}

	return stores, nil
}

func (ar *storehouseRepository) CountAuthorizedStorehouses(ctx context.Context, userID int, query string) (int64, error) {
	var count int64
	var err error

	q := ar.db.WithContext(ctx).Table("storehouses").Joins("LEFT JOIN authorized on authorized.storehouse_id = storehouses.id").
		Where("authorized.user_id = ?", userID).Where("deleted_at is NULL")

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

func (ar *storehouseRepository) GetAuthorizedStorehouses(ctx context.Context, userID int, query string, limit, skip int) ([]domain.Storehouse, error) {
	list := []schema.Storehouse{}
	var err error

	trimQuery := strings.TrimSpace(query)

	q := ar.db.WithContext(ctx).Joins("LEFT JOIN authorized on authorized.storehouse_id = storehouses.id").
		Limit(limit).Offset((skip-1)*limit).Where("authorized.user_id = ? AND deleted_at is NULL", userID).Order("id desc")

	if trimQuery != "" {
		q.Where("name LIKE ?", fmt.Sprintf("%%%v%%", trimQuery))
	}

	err = q.Find(&list).Error
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

func (sr *storehouseRepository) GetUsedCapacityByID(ctx context.Context, id int) (int64, error) {
	err := sr.db.WithContext(ctx).First(&schema.Storehouse{ID: id}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, domain.ErrDataNotFound
		}
		return 0, err
	}

	total := struct {
		Total int64
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
	result := sr.db.WithContext(ctx).
		Model(&schema.Storehouse{}).Where("id = ?", storehouses.ID).
		Updates(&schema.Storehouse{
			Name:     storehouses.Name,
			Location: storehouses.Location,
			Capacity: storehouses.Capacity,
			Image:    storehouses.Image,
		})

	err := result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, domain.ErrNoUpdatedData
	}

	return sr.GetStorehouseByID(ctx, storehouses.ID)
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
