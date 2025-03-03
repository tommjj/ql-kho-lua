package services

import (
	"context"
	"time"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/mapmutex"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type imInvoiceService struct {
	imInvoiceRepo ports.IImportInvoicesRepository
	warehouseRepo ports.IWarehouseRepository
	l             *mapmutex.Mapmutex
}

func NewImInvoicesService(
	imInvoiceRepo ports.IImportInvoicesRepository,
	warehouseRepo ports.IWarehouseRepository,
	l *mapmutex.Mapmutex) ports.IImportInvoicesService {
	return &imInvoiceService{
		imInvoiceRepo: imInvoiceRepo,
		warehouseRepo: warehouseRepo,
		l:             l,
	}
}

func (i *imInvoiceService) CreateImInvoice(ctx context.Context, invoice *domain.Invoice) (*domain.Invoice, error) {
	i.l.Lock(invoice.WarehouseID)
	defer i.l.UnLock(invoice.WarehouseID)

	store, err := i.warehouseRepo.GetWarehouseByID(ctx, invoice.WarehouseID)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, err
	}

	used, err := i.warehouseRepo.GetUsedCapacityByID(ctx, invoice.WarehouseID)
	if err != nil {
		if err != domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	var capacity int
	for _, v := range invoice.Details {
		capacity += v.Quantity
	}
	if (int(used) + capacity) > store.Capacity {
		return nil, domain.ErrWarehouseFull
	}

	invoice.CalcTotalPrice()
	created, err := i.imInvoiceRepo.CreateImInvoice(ctx, invoice)
	if err != nil {
		switch err {
		case domain.ErrDataNotFound:
			return nil, domain.ErrDataNotFound
		default:
			return nil, domain.ErrInternal
		}
	}

	return created, nil
}

func (i *imInvoiceService) GetImInvoiceByID(ctx context.Context, id int) (*domain.Invoice, error) {
	invoice, err := i.imInvoiceRepo.GetImInvoiceWithAssociationsByID(ctx, id)
	if err != nil {
		switch err {
		case domain.ErrDataNotFound:
			return nil, domain.ErrDataNotFound
		default:
			return nil, domain.ErrInternal
		}
	}

	return invoice, nil
}

func (i *imInvoiceService) CountImInvoices(ctx context.Context, warehouseID int, start *time.Time, end *time.Time) (int64, error) {
	count, err := i.imInvoiceRepo.CountImInvoices(ctx, warehouseID, start, end)
	if err != nil {
		return 0, domain.ErrInternal
	}

	return count, nil
}

func (i *imInvoiceService) GetListImInvoices(ctx context.Context, warehouseID int, start *time.Time, end *time.Time, skip, limit int) ([]domain.Invoice, error) {
	invoice, err := i.imInvoiceRepo.GetListImInvoices(ctx, warehouseID, start, end, skip, limit)
	if err != nil {
		switch err {
		case domain.ErrDataNotFound:
			return nil, domain.ErrDataNotFound
		default:
			return nil, domain.ErrInternal
		}
	}

	return invoice, nil
}
