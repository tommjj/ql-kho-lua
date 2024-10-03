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

func NewDefaultCustomerRepo() (ports.ICustomerRepository, error) {
	db, err := mysqldb.NewMysqlDB(config.DB{
		DSN:             "root:@tcp(127.0.0.1:3306)/ql?charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		return nil, err
	}

	return NewCustomerRepository(db), nil
}

func TestCustomerRepository_CreateCustomer(t *testing.T) {
	repo, err := NewDefaultCustomerRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.CreateCustomer(context.TODO(),
		&domain.Customer{
			Name:    "vertin",
			Email:   "vertin@m.com",
			Phone:   "0123456789",
			Address: "abc xyz",
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}

func TestCustomerRepository_GetCustomerByID(t *testing.T) {
	repo, err := NewDefaultCustomerRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.GetCustomerByID(context.TODO(), 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}

func TestCustomerRepository_GetListCustomers(t *testing.T) {
	repo, err := NewDefaultCustomerRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.GetListCustomers(context.TODO(), "cus", 5, 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}

func TestCustomerRepository_UpdateCustomer(t *testing.T) {
	repo, err := NewDefaultCustomerRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.UpdateCustomer(context.TODO(),
		&domain.Customer{
			ID:      2,
			Name:    "mostima",
			Email:   "cus@m.com",
			Phone:   "0123456789",
			Address: "abc xyz",
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}

func TestCustomerRepository_DeleteCustomer(t *testing.T) {
	repo, err := NewDefaultCustomerRepo()
	if err != nil {
		t.Fatal(err)
	}

	err = repo.DeleteCustomer(context.TODO(), 1)
	if err != nil {
		t.Fatal(err)
	}

}
