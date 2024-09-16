package repository

import "github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"

type customerRepository struct {
	db *mysqldb.MysqlDB
}

func NewCustomerRepository(db *mysqldb.MysqlDB) *customerRepository {
	return &customerRepository{
		db: db,
	}
}
