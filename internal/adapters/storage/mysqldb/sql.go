package mysqldb

import (
	"github.com/tommjj/ql-kho-lua/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	*gorm.DB
}

func NewMysqlDB(conf config.DB) (*MysqlDB, error) {
	dialector := mysql.New(mysql.Config{
		DSN: conf.DSN,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	mysql, err := db.DB()
	if err != nil {
		return nil, err
	}

	mysql.SetMaxIdleConns(conf.MaxIdleConns)
	mysql.SetMaxOpenConns(conf.MaxOpenConns)
	mysql.SetConnMaxLifetime(conf.ConnMaxLifetime)

	err = mysql.Ping()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate()
	if err != nil {
		return nil, err
	}

	return &MysqlDB{
		db,
	}, nil
}
