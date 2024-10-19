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

type warehouseRepository struct {
	db *mysqldb.MysqlDB
}

func NewWarehouseRepository(db *mysqldb.MysqlDB) ports.IWarehouseRepository {
	return &warehouseRepository{
		db: db,
	}
}

func (sr *warehouseRepository) CreateWarehouse(ctx context.Context, warehouses *domain.Warehouse) (*domain.Warehouse, error) {
	createData := &schema.Warehouse{
		Name:     warehouses.Name,
		Location: warehouses.Location,
		Capacity: warehouses.Capacity,
		Image:    warehouses.Image,
	}

	err := sr.db.WithContext(ctx).Create(createData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return convertToWarehouse(createData), nil
}

func (sr *warehouseRepository) GetWarehouseByID(ctx context.Context, id int) (*domain.Warehouse, error) {
	store := &schema.Warehouse{}

	err := sr.db.WithContext(ctx).Where("id = ?", id).First(store).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return convertToWarehouse(store), nil
}

func (sr *warehouseRepository) CountWarehouses(ctx context.Context, query string) (int64, error) {
	var count int64
	var err error

	q := sr.db.WithContext(ctx).Table("warehouses").Where("deleted_at is NULL")

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

func (sr *warehouseRepository) GetListWarehouses(ctx context.Context, query string, limit, skip int) ([]domain.Warehouse, error) {
	stores := []domain.Warehouse{}
	var err error
	var rows *sql.Rows

	sql := sr.db.WithContext(ctx).Table("warehouses").
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
		store := domain.Warehouse{}
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

func (ar *warehouseRepository) CountAuthorizedWarehouses(ctx context.Context, userID int, query string) (int64, error) {
	var count int64
	var err error

	q := ar.db.WithContext(ctx).Table("warehouses").Joins("LEFT JOIN authorized on authorized.warehouse_id = warehouses.id").
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

func (ar *warehouseRepository) GetAuthorizedWarehouses(ctx context.Context, userID int, query string, limit, skip int) ([]domain.Warehouse, error) {
	list := []schema.Warehouse{}
	var err error

	trimQuery := strings.TrimSpace(query)

	q := ar.db.WithContext(ctx).Joins("LEFT JOIN authorized on authorized.warehouse_id = warehouses.id").
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

	warehouse := make([]domain.Warehouse, 0, len(list))

	for _, v := range list {
		warehouse = append(warehouse, *convertToWarehouse(&v))
	}

	return warehouse, nil
}

func (sr *warehouseRepository) GetUsedCapacityByID(ctx context.Context, id int) (int64, error) {
	err := sr.db.WithContext(ctx).First(&schema.Warehouse{ID: id}).Error
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
    			WHERE export_invoices.warehouse_id = @id 
    			UNION ALL
    			SELECT SUM( import_invoice_details.quantity) as total, "i" as type 
    			FROM import_invoices INNER JOIN import_invoice_details on import_invoices.id = import_invoice_details.invoice_id 
    			WHERE import_invoices.warehouse_id = @id) as t`, sql.Named("id", id)).Scan(&total).Error
	if err != nil {
		return 0, err
	}

	return total.Total, nil
}

func (sr *warehouseRepository) UpdateWarehouse(ctx context.Context, warehouse *domain.Warehouse) (*domain.Warehouse, error) {
	result := sr.db.WithContext(ctx).
		Model(&schema.Warehouse{}).Where("id = ?", warehouse.ID).
		Updates(&schema.Warehouse{
			Name:     warehouse.Name,
			Location: warehouse.Location,
			Capacity: warehouse.Capacity,
			Image:    warehouse.Image,
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

	return sr.GetWarehouseByID(ctx, warehouse.ID)
}

func (sr *warehouseRepository) DeleteWarehouse(ctx context.Context, id int) error {
	result := sr.db.WithContext(ctx).Where("id = ?", id).Delete(&schema.Warehouse{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrDataNotFound
	}

	return nil
}
