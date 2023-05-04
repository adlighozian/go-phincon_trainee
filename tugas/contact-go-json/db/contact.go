package db

import (
	"contact-go/config"
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetMysql(cfg *config.Config) *sql.DB {

	var connStrings = ""
	if cfg.DbMain == "mysql" {
		connStrings = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DbUsername, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	}

	db, err := sql.Open(cfg.DbMain, connStrings)
	if err != nil {
		panic(err)
	}

	// pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func NewMysqlContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
