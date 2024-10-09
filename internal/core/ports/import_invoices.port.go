package ports

import (
	"context"
	"time"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

type IImportInvoicesRepository interface {
	// CreateImInvoice create a new import invoice
	CreateImInvoice(ctx context.Context, invoice *domain.Invoice) (*domain.Invoice, error)
	// GetImInvoiceByID select a invoice by id
	GetImInvoiceByID(ctx context.Context, id int) (*domain.Invoice, error)
	// GetImInvoiceWithAssociationsByID select a invoice with user, storehouse, customer, rice by id
	GetImInvoiceWithAssociationsByID(ctx context.Context, id int) (*domain.Invoice, error)
	// GetListImInvoices select invoices
	GetListImInvoices(ctx context.Context, start *time.Time, end *time.Time, skip, limit int) ([]domain.Invoice, error)
}
