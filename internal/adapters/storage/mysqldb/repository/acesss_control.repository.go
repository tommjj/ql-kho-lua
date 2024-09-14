package repository

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

type accessControlRepository struct {
	db *mysqldb.MysqlDB
}

func NewAccessControlRepository(db *mysqldb.MysqlDB) *accessControlRepository {
	return &accessControlRepository{
		db: db,
	}
}

func (r *accessControlRepository) hasAccess(ctx context.Context, storeHouseID int, token *domain.TokenPayload) error {
	if token.Role == domain.Root {
		return nil
	}

	return nil
}
