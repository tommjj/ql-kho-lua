package services

import (
	"context"
	"time"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/mapmutex"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type exInvoiceService struct {
	imInvoiceRepo ports.IExportInvoiceRepository
	warehouseRepo ports.IWarehouseRepository
	l             *mapmutex.Mapmutex
}

func NewExInvoicesService(
	exInvoiceRepo ports.IExportInvoiceRepository,
	warehouseRepo ports.IWarehouseRepository,
	l *mapmutex.Mapmutex) ports.IExportInvoiceService {
	return &exInvoiceService{
		imInvoiceRepo: exInvoiceRepo,
		warehouseRepo: warehouseRepo,
		l:             l,
	}
}

func (e *exInvoiceService) CreateExInvoice(ctx context.Context, invoice *domain.Invoice) (*domain.Invoice, error) {
	e.l.Lock(invoice.WarehouseID)
	defer e.l.UnLock(invoice.WarehouseID)

	inventory, err := e.warehouseRepo.GetInventory(ctx, invoice.WarehouseID)
	if err != nil {
		if err != domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	if len(inventory) < len(invoice.Details) {
		return nil, domain.ErrInsufficientStock
	}

	for _, detail := range invoice.Details {
		isHaveItem := false
		for _, item := range inventory {
			if item.RiceID != detail.RiceID {
				continue
			}

			isInsufficientStock := item.Quantity < detail.Quantity
			if isInsufficientStock {
				return nil, domain.ErrInsufficientStock
			}

			isHaveItem = true
		}

		if !isHaveItem {
			return nil, domain.ErrInsufficientStock
		}
	}

	invoice.CalcTotalPrice()
	created, err := e.imInvoiceRepo.CreateExInvoice(ctx, invoice)
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

func (e *exInvoiceService) GetExInvoiceByID(ctx context.Context, id int) (*domain.Invoice, error) {
	invoice, err := e.imInvoiceRepo.GetExInvoiceWithAssociationsByID(ctx, id)
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

func (e *exInvoiceService) CountExInvoices(ctx context.Context, warehouseID int, start *time.Time, end *time.Time) (int64, error) {
	count, err := e.imInvoiceRepo.CountExInvoices(ctx, warehouseID, start, end)
	if err != nil {
		return 0, domain.ErrInternal
	}

	return count, nil
}

func (e *exInvoiceService) GetListExInvoices(ctx context.Context, warehouseID int, start *time.Time, end *time.Time, skip, limit int) ([]domain.Invoice, error) {
	invoice, err := e.imInvoiceRepo.GetListExInvoices(ctx, warehouseID, start, end, skip, limit)
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
