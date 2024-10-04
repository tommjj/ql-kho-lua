package mysqldb

import (
	"context"
	"sync"
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
		MaxIdleConns:    100,
		MaxOpenConns:    140,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		t.Fatal(err)
	}

	q := db.Table("import_invoices").Select("*").Limit(1).WithContext(context.TODO())

	wg := sync.WaitGroup{}

	wg.Add(140)

	for range 140 {
		go func() {
			for range 1000 {
				a := schema.ImportInvoice{}
				var count int64
				q.Count(&count)
				q.Scan(&a)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestDown(t *testing.T) {
	db, err := NewMysqlDB(config.DB{
		DSN:             "root:@tcp(127.0.0.1:3306)/ql?charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	})

	if err != nil {
		t.Fatal(err)
	}

	m := db.Migrator()
	m.DropTable(
		&schema.ExportInvoiceDetail{},
		&schema.ExportInvoice{},
		&schema.ImportInvoiceDetail{},
		&schema.ImportInvoice{},
		&schema.Customer{},
		"authorized",
		&schema.User{},
		&schema.Rice{},
		&schema.Storehouse{},
	)
}

func TestUp(t *testing.T) {
	db, err := NewMysqlDB(config.DB{
		DSN:             "root:@tcp(127.0.0.1:3306)/ql?charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	})

	if err != nil {
		t.Fatal(err)
	}

	m := db.Migrator()
	m.AutoMigrate(
		&schema.User{},
		&schema.Storehouse{},
		&schema.Customer{},
		&schema.Rice{},
		&schema.ExportInvoice{},
		&schema.ExportInvoiceDetail{},
		&schema.ImportInvoice{},
		&schema.ImportInvoiceDetail{},
	)
}
