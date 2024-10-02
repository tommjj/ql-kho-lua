package repository

import (
	"context"
	"errors"

	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb/schema"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"gorm.io/gorm"
)

type importInvoicesRepository struct {
	db *mysqldb.MysqlDB
}

func NewImInvoicesRepository(db *mysqldb.MysqlDB) *importInvoicesRepository {
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
