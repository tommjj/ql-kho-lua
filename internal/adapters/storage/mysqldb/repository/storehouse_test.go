package repository

import (
	"context"
	"testing"
	"time"

	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/config"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

func NewDefaultStoreHouseRepo() (*storehouseRepository, error) {
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

func TestStoreHouseRepo_Create(t *testing.T) {
	repo, err := NewDefaultStoreHouseRepo()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.Create(context.TODO(), &domain.Storehouse{Name: "store01", Location: "0, 0", Capacity: 1000})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(data)
}
