package mysqldb

import (
	"testing"
	"time"

	"github.com/tommjj/ql-kho-lua/internal/config"
)

func TestConnection(t *testing.T) {
	_, err := NewMysqlDB(config.DB{
		DSN:             "root:@tcp(127.0.0.1:3306)/ql?charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	})

	if err != nil {
		t.Fatal(err)
	}

}

func TestRaw(t *testing.T) {
	db, err := NewMysqlDB(config.DB{
		DSN:             "root:@tcp(127.0.0.1:3306)/ql?charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	})

	if err != nil {
		t.Fatal(err)
	}

	db.Begin()

}
