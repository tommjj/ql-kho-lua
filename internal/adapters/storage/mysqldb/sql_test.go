package mysqldb

import (
	"context"
	"testing"
	"time"

	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb/schema"
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
	var count int64
	a := schema.ImportInvoice{}
	q := db.Table("import_invoices").WithContext(context.TODO())
	q.Select("*").Limit(1).Where("id = 2").Scan(&a)
	q.Select("*").Limit(1).Count(&count)
	now := time.Now()
	t.Log(count)
	t.Log(a)
	t.Log(time.Since(now))
}
