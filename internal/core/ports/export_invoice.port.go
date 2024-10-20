package ports

import (
	"context"
	"time"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

type IExportInvoiceRepository interface {
	// CreateExInvoice create a new export invoice
	CreateExInvoice(ctx context.Context, invoice *domain.Invoice) (*domain.Invoice, error)
	// GetExInvoiceByID select a invoice by id
	GetExInvoiceByID(ctx context.Context, id int) (*domain.Invoice, error)
	// GetExInvoiceWithAssociationsByID select a invoice with user, warehouse, customer, rice by id
	GetExInvoiceWithAssociationsByID(ctx context.Context, id int) (*domain.Invoice, error)
	// CountExInvoices
	CountExInvoices(ctx context.Context, warehouseID int, start *time.Time, end *time.Time) (int64, error)
	// GetListExInvoices select invoices
	GetListExInvoices(ctx context.Context, warehouseID int, start *time.Time, end *time.Time, skip, limit int) ([]domain.Invoice, error)
}

type IExportInvoiceService interface {
	// CreateExInvoice create a new export invoice
	CreateExInvoice(ctx context.Context, invoice *domain.Invoice) (*domain.Invoice, error)
	// GetExInvoiceByID select a invoice by id
	GetExInvoiceByID(ctx context.Context, id int) (*domain.Invoice, error)
	// CountExInvoices
	CountExInvoices(ctx context.Context, warehouseID int, start *time.Time, end *time.Time) (int64, error)
	// GetListExInvoices select invoices
	GetListExInvoices(ctx context.Context, warehouseID int, start *time.Time, end *time.Time, skip, limit int) ([]domain.Invoice, error)
}
