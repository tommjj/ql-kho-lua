package repository

import (
	"context"
	"errors"
	"time"

	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb/schema"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type importInvoiceRepository struct {
	db *mysqldb.MysqlDB
}

func NewImInvoicesRepository(db *mysqldb.MysqlDB) ports.IImportInvoicesRepository {
	return &importInvoiceRepository{
		db: db,
	}
}

func (i *importInvoiceRepository) CreateImInvoice(ctx context.Context, invoice *domain.Invoice) (*domain.Invoice, error) {
	createData := &schema.ImportInvoice{
		WarehouseID: invoice.WarehouseID,
		CustomerID:  invoice.CustomerID,
		UserID:      invoice.UserID,
		Details:     make([]schema.ImportInvoiceDetail, len(invoice.Details)),
		TotalPrice:  invoice.TotalPrice,
	}

	for i, detail := range invoice.Details {
		createData.Details[i] = schema.ImportInvoiceDetail{
			RiceID:   detail.RiceID,
			Price:    detail.Price,
			Quantity: detail.Quantity,
		}
	}

	err := i.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Preload("Details").Create(createData).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrForeignKeyViolated):
			return nil, domain.ErrDataNotFound
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, domain.ErrDataNotFound
		default:
			return nil, err
		}
	}

	return i.GetImInvoiceWithAssociationsByID(ctx, createData.ID)
}

func (i *importInvoiceRepository) GetImInvoiceByID(ctx context.Context, id int) (*domain.Invoice, error) {
	data := &schema.ImportInvoice{}

	err := i.db.WithContext(ctx).Preload("Details").Where("id = ?", id).First(data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	invoice := &domain.Invoice{
		ID:          data.ID,
		UserID:      data.UserID,
		CustomerID:  data.CustomerID,
		WarehouseID: data.WarehouseID,
		TotalPrice:  data.TotalPrice,
		Details:     make([]domain.InvoiceItem, len(data.Details)),
	}

	for i, detail := range data.Details {
		invoice.Details[i] = domain.InvoiceItem{
			Price:    detail.Price,
			Quantity: detail.Quantity,
			RiceID:   detail.RiceID,
		}
	}

	return invoice, nil
}

func (i *importInvoiceRepository) GetImInvoiceWithAssociationsByID(ctx context.Context, id int) (*domain.Invoice, error) {
	data := &schema.ImportInvoice{}

	err := i.db.WithContext(ctx).Preload("Details.Rice").
		Preload(clause.Associations).Where("id = ?", id).First(data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	invoice := &domain.Invoice{
		ID:          data.ID,
		UserID:      data.UserID,
		CustomerID:  data.CustomerID,
		WarehouseID: data.WarehouseID,
		TotalPrice:  data.TotalPrice,
		CreatedAt:   data.CreatedAt,
		Details:     make([]domain.InvoiceItem, len(data.Details)),
	}

	if data.Customer.ID != 0 {
		invoice.Customer = convertToCustomer(&data.Customer)
	}

	if data.Warehouse.ID != 0 {
		invoice.Warehouse = convertToWarehouse(&data.Warehouse)
	}

	if data.User.ID != 0 {
		invoice.CreatedBy = convertToUser(&data.User)
	}

	for i, detail := range data.Details {
		invoice.Details[i] = domain.InvoiceItem{
			Price:    detail.Price,
			Quantity: detail.Quantity,
			RiceID:   detail.RiceID,
		}

		if detail.Rice.ID != 0 {
			invoice.Details[i].Rice = convertToRice(&detail.Rice)
		}
	}

	return invoice, nil
}

func (i *importInvoiceRepository) CountImInvoices(ctx context.Context, warehouseID int, start *time.Time, end *time.Time) (int64, error) {
	var count int64

	q := i.db.WithContext(ctx).Model(&schema.ImportInvoice{})
	if start != nil {
		q.Where("created_at >= ?", start)
	}
	if end != nil {
		q.Where("created_at <= ?", end)
	}
	if warehouseID != 0 {
		q.Where("warehouse_id = ?", warehouseID)
	}

	err := q.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (i *importInvoiceRepository) GetListImInvoices(ctx context.Context, warehouseID int, start *time.Time, end *time.Time, skip, limit int) ([]domain.Invoice, error) {
	invoices := []domain.Invoice{}

	q := i.db.WithContext(ctx).Select(
		"id", "warehouse_id", "customer_id", "user_id", "created_at", "total_price",
	).Model(&schema.ImportInvoice{}).Limit(limit).Offset((skip - 1) * limit).Order("id DESC")

	if start != nil {
		q.Where("created_at >= ?", start)
	}
	if end != nil {
		q.Where("created_at <= ?", end)
	}
	if warehouseID != 0 {
		q.Where("warehouse_id = ?", warehouseID)
	}

	rows, err := q.Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		invoice := domain.Invoice{}
		rows.Scan(
			&invoice.ID,
			&invoice.WarehouseID,
			&invoice.CustomerID,
			&invoice.UserID,
			&invoice.CreatedAt,
			&invoice.TotalPrice,
		)

		invoices = append(invoices, invoice)
	}

	if len(invoices) == 0 {
		return nil, domain.ErrDataNotFound
	}

	return invoices, nil
}
