package mysqldb

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	*gorm.DB
}

func NewMysqlDB() (*MysqlDB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/ql?charset=utf8mb4&parseTime=True&loc=Local"

	dbConn := mysql.New(mysql.Config{
		DSN: dsn,
	})

	db, err := gorm.Open(dbConn, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	mysql, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	mysql.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	mysql.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	mysql.SetConnMaxLifetime(time.Hour)

	err = mysql.Ping()
	if err != nil {
		return nil, err
	}

	return &MysqlDB{
		db,
	}, nil
}
