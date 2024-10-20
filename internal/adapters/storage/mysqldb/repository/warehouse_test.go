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

func NewDefaultWarehouseRepo() (ports.IWarehouseRepository, error) {
	db, err := mysqldb.NewMysqlDB(config.DB{
		DSN:             "root:@tcp(127.0.0.1:3306)/ql?charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		return nil, err
	}

	return NewWarehouseRepository(db), nil
}

func TestWarehouseRepo_CreateWarehouse(t *testing.T) {
	repo, err := NewDefaultWarehouseRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.CreateWarehouse(context.TODO(), &domain.Warehouse{Name: "Store 04", Location: "40.431858734948605,-99.95028183893876", Capacity: 1000, Image: "f20e8a37-af94-4a09-99a4-b62e7b2edbdb.png"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(data)
}

func TestWarehouseRepo_GetWarehouseByID(t *testing.T) {
	repo, err := NewDefaultWarehouseRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.GetWarehouseByID(context.TODO(), 3)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}

func TestWareHouseRepo_CountWarehouses(t *testing.T) {
	repo, err := NewDefaultWarehouseRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.CountWarehouses(context.TODO(), "")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}
func TestWarehouseRepo_GetListWarehouses(t *testing.T) {
	repo, err := NewDefaultWarehouseRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.GetListWarehouses(context.TODO(), "3", 5, 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}

func TestWarehouseRepo_CountAuthorizedWarehouses(t *testing.T) {
	repo, err := NewDefaultWarehouseRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.CountAuthorizedWarehouses(context.TODO(), 1, "0")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}

func TestWarehouseRepo_GetAuthorizedWarehouses(t *testing.T) {
	repo, err := NewDefaultWarehouseRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.GetAuthorizedWarehouses(context.TODO(), 1, "03", 1, 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}

func TestWarehouseRepo_GetWarehouseUsedCapacityByID(t *testing.T) {
	repo, err := NewDefaultWarehouseRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.GetUsedCapacityByID(context.TODO(), 2)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}

func TestWarehouseRepo_GetInventory(t *testing.T) {
	repo, err := NewDefaultWarehouseRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.GetInventory(context.TODO(), 2)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}
