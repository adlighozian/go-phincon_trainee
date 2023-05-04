package db

import (
	"contact-go/config"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type dbOpt struct {
	Database string
}

func GetDB(db string) dbOpt {
	return dbOpt{
		Database: db,
	}
}

func (dbs dbOpt) GetConnectionMysql() (*sql.DB, error) {
	config := config.LoadConfig()

	var connStrings = ""
	if dbs.Database == "mysql" {
		connStrings = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	}

	db, err := sql.Open(config.DbMain, connStrings)
	if err != nil {
		panic(err)
	}

	// pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}
