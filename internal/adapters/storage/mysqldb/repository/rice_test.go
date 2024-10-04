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

func NewDefaultRiceRepo() (ports.IRiceRepository, error) {
	db, err := mysqldb.NewMysqlDB(config.DB{
		DSN:             "root:@tcp(127.0.0.1:3306)/ql?charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		return nil, err
	}

	return NewRiceRepository(db), nil
}

func TestRice_Create(t *testing.T) {
	repo, err := NewDefaultRiceRepo()
	if err != nil {
		t.Fatal(err)
	}

	rice, err := repo.CreateRice(context.TODO(), &domain.Rice{
		Name: "rice 03",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", rice)
}

func TestRice_GetRiceByID(t *testing.T) {
	repo, err := NewDefaultRiceRepo()
	if err != nil {
		t.Fatal(err)
	}

	rice, err := repo.GetRiceByID(context.TODO(), 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", rice)
}

func TestRice_GetListRice(t *testing.T) {
	repo, err := NewDefaultRiceRepo()
	if err != nil {
		t.Fatal(err)
	}

	rice, err := repo.GetListRice(context.TODO(), "1", 1, 5)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", rice)
}

func TestRice_UpdateRice(t *testing.T) {
	repo, err := NewDefaultRiceRepo()
	if err != nil {
		t.Fatal(err)
	}

	rice, err := repo.UpdateRice(context.TODO(), &domain.Rice{
		ID:   2,
		Name: "rice 03",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", rice)
}

func TestRice_DeleteRice(t *testing.T) {
	repo, err := NewDefaultRiceRepo()
	if err != nil {
		t.Fatal(err)
	}

	err = repo.DeleteRice(context.TODO(), 3)
	if err != nil {
		t.Fatal(err)
	}
}
