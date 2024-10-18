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

func NewDefaultImInvoicesRepo() (ports.IImportInvoicesRepository, error) {
	db, err := mysqldb.NewMysqlDB(config.DB{
		DSN:             "root:@tcp(127.0.0.1:3306)/ql?charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		return nil, err
	}

	return NewImInvoicesRepository(db), nil
}

func TestImInvoices_createInvoice(t *testing.T) {
	repo, err := NewDefaultImInvoicesRepo()
	if err != nil {

		t.Fatal(err)
	}

	data, err := repo.CreateImInvoice(context.TODO(), &domain.Invoice{
		UserID:       1,
		CustomerID:   1,
		StorehouseID: 2,
		Details: []domain.InvoiceItem{
			{
				RiceID:   1,
				Price:    200,
				Quantity: 120,
			},
			{
				RiceID:   2,
				Price:    500,
				Quantity: 5052,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}

func TestImInvoices_getInvoice(t *testing.T) {
	repo, err := NewDefaultImInvoicesRepo()

	if err != nil {

		t.Fatal(err)
	}

	data, err := repo.GetImInvoiceByID(context.TODO(), 3)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data.Details)
}

func TestImInvoices_getListInvoices(t *testing.T) {
	repo, err := NewDefaultImInvoicesRepo()
	if err != nil {

		t.Fatal(err)
	}

	start := time.Date(2024, 10, 3, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC)

	data, err := repo.GetListImInvoices(context.TODO(), &start, &end, 1, 5)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}
