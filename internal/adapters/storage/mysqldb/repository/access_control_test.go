package repository

import (
	"context"
	"testing"
	"time"

	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/config"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

func NewDefaultAccessControlRepo() (ports.IAccessControlRepository, error) {
	db, err := mysqldb.NewMysqlDB(config.DB{
		DSN:             "root:@tcp(127.0.0.1:3306)/ql?charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		return nil, err
	}

	return NewAccessControlRepository(db), nil
}

func TestAccessControl_HasAccess(t *testing.T) {
	repo, err := NewDefaultAccessControlRepo()
	if err != nil {
		t.Fatal(err)
	}

	err = repo.HasAccess(context.TODO(), 2, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAccessControl_SetAccess(t *testing.T) {
	repo, err := NewDefaultAccessControlRepo()
	if err != nil {
		t.Fatal(err)
	}

	err = repo.SetAccess(context.TODO(), 2, 2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAccessControl_DelAccess(t *testing.T) {
	repo, err := NewDefaultAccessControlRepo()
	if err != nil {
		t.Fatal(err)
	}

	err = repo.DelAccess(context.TODO(), 2, 2)
	if err != nil {
		t.Fatal(err)
	}
}
