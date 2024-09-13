package repository

import "github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"

type accessControlRepository struct {
	db *mysqldb.MysqlDB
}

func NewAccessControlRepository(db *mysqldb.MysqlDB) *accessControlRepository {
	return &accessControlRepository{
		db: db,
	}
}
