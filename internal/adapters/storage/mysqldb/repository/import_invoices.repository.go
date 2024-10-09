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

type importInvoicesRepository struct {
	db *mysqldb.MysqlDB
}

func NewImInvoicesRepository(db *mysqldb.MysqlDB) ports.IImportInvoicesRepository {
	return &importInvoicesRepository{
		db: db,
	}
}

func (eir *importInvoicesRepository) CreateImInvoice(ctx context.Context, invoice *domain.Invoice) (*domain.Invoice, error) {
	createData := &schema.ImportInvoice{
		StorehouseID: invoice.StorehouseID,
		CustomerID:   invoice.CustomerID,
		UserID:       invoice.UserID,
		Details:      make([]schema.ImportInvoiceDetail, len(invoice.Details)),
	}

	for i, detail := range invoice.Details {
		createData.Details[i] = schema.ImportInvoiceDetail{
			RiceID:   detail.RiceID,
			Price:    detail.Price,
			Quantity: detail.Quantity,
		}
		createData.TotalPrice += detail.Price * float64(detail.Quantity)
	}

	err := eir.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Preload("Details").Create(createData).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrForeignKeyViolated):
			return nil, domain.ErrConflictingData
		default:
			return nil, err
		}
	}

	created := &domain.Invoice{
		ID:           createData.ID,
		UserID:       createData.UserID,
		CustomerID:   createData.CustomerID,
		StorehouseID: createData.CustomerID,
		TotalPrice:   createData.TotalPrice,
		Details:      make([]domain.InvoiceItem, len(createData.Details)),
	}

	for i, detail := range createData.Details {
		created.Details[i] = domain.InvoiceItem{
			Price:    detail.Price,
			Quantity: detail.Quantity,
			RiceID:   detail.RiceID,
		}
	}

	return created, nil
}

func (eir *importInvoicesRepository) GetImInvoiceByID(ctx context.Context, id int) (*domain.Invoice, error) {
	data := &schema.ImportInvoice{}

	err := eir.db.WithContext(ctx).Preload("Details").Where("id = ?", id).First(data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	invoice := &domain.Invoice{
		ID:           data.ID,
		UserID:       data.UserID,
		CustomerID:   data.CustomerID,
		StorehouseID: data.StorehouseID,
		TotalPrice:   data.TotalPrice,
		Details:      make([]domain.InvoiceItem, len(data.Details)),
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

func (eir *importInvoicesRepository) GetImInvoiceWithAssociationsByID(ctx context.Context, id int) (*domain.Invoice, error) {
	data := &schema.ImportInvoice{}

	err := eir.db.WithContext(ctx).Preload("Details.Rice").
		Preload(clause.Associations).Where("id = ?", id).First(data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	invoice := &domain.Invoice{
		ID:           data.ID,
		UserID:       data.UserID,
		CustomerID:   data.CustomerID,
		StorehouseID: data.StorehouseID,
		TotalPrice:   data.TotalPrice,
		Details:      make([]domain.InvoiceItem, len(data.Details)),
	}

	if data.Customer.ID != 0 {
		invoice.Customer = convertToCustomer(&data.Customer)
	}

	if data.Storehouse.ID != 0 {
		invoice.Storehouse = convertToStorehouse(&data.Storehouse)
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

func (eir *importInvoicesRepository) CountImInvoices(ctx context.Context) (int64, error) {
	var count int64

	err := eir.db.WithContext(ctx).Model(&schema.ImportInvoice{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (eir *importInvoicesRepository) GetListImInvoices(ctx context.Context, start *time.Time, end *time.Time, skip, limit int) ([]domain.Invoice, error) {
	invoices := []domain.Invoice{}

	rows, err := eir.db.WithContext(ctx).Select(
		"id", "storehouse_id", "customer_id", "user_id", "created_at", "total_price",
	).Model(&schema.ImportInvoice{}).Limit(limit).Offset((skip - 1) * limit).Order("id DESC").Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		invoice := domain.Invoice{}
		rows.Scan(
			&invoice.ID,
			&invoice.StorehouseID,
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
