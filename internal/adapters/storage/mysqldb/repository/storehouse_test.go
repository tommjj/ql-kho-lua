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

func NewDefaultStoreHouseRepo() (ports.IStorehouseRepository, error) {
	db, err := mysqldb.NewMysqlDB(config.DB{
		DSN:             "root:@tcp(127.0.0.1:3306)/ql?charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		return nil, err
	}

	return NewStorehouseRepository(db), nil
}

func TestStoreHouseRepo_CreateStorehouse(t *testing.T) {
	repo, err := NewDefaultStoreHouseRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.CreateStorehouse(context.TODO(), &domain.Storehouse{Name: "store 03", Location: "50, 53", Capacity: 1000, Image: "f20e8a37-af94-4a09-99a4-b62e7b2edbdb.png"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(data)
}

func TestStoreHouseRepo_GetStorehouseByID(t *testing.T) {
	repo, err := NewDefaultStoreHouseRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.GetStorehouseByID(context.TODO(), 3)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}

func TestStoreHouseRepo_GetListStorehouses(t *testing.T) {
	repo, err := NewDefaultStoreHouseRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.GetListStorehouses(context.TODO(), "", 5, 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}

func TestStoreHouseRepo_GetAuthorizedStorehouses(t *testing.T) {
	repo, err := NewDefaultStoreHouseRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.GetAuthorizedStorehouses(context.TODO(), 1, "", 1, 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}

func TestStoreHouseRepo_GetStorehouseUsedCapacityByID(t *testing.T) {
	repo, err := NewDefaultStoreHouseRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.GetStorehouseUsedCapacityByID(context.TODO(), 2)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", data)
}
