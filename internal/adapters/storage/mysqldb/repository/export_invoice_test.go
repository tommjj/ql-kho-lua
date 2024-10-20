package repository

import (
	"context"
	"testing"
	"time"

	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/config"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

func NewDefaultExInvoicesRepo() (ports.IExportInvoiceRepository, error) {
	db, err := mysqldb.NewMysqlDB(config.DB{
		DSN:             "root:@tcp(127.0.0.1:3306)/ql?charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		return nil, err
	}

	return NewExInvoicesRepository(db), nil
}

func TestExInvoices_createInvoice(t *testing.T) {
	repo, err := NewDefaultExInvoicesRepo()
	if err != nil {

		t.Fatal(err)
	}

	create := &domain.Invoice{
		UserID:      1,
		CustomerID:  1,
		WarehouseID: 2,
		Details: []domain.InvoiceItem{
			{
				RiceID:   1,
				Price:    200,
				Quantity: 20,
			},
			{
				RiceID:   2,
				Price:    300,
				Quantity: 10,
			},
		},
	}
	create.CalcTotalPrice()

	data, err := repo.CreateExInvoice(context.TODO(), create)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}

func TestExInvoices_getInvoice(t *testing.T) {
	repo, err := NewDefaultExInvoicesRepo()

	if err != nil {

		t.Fatal(err)
	}

	data, err := repo.GetExInvoiceWithAssociationsByID(context.TODO(), 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data.Details)
}

func TestExInvoices_getListInvoices(t *testing.T) {
	repo, err := NewDefaultExInvoicesRepo()
	if err != nil {

		t.Fatal(err)
	}

	data, err := repo.GetListExInvoices(context.TODO(), 2, nil, nil, 1, 5)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}
