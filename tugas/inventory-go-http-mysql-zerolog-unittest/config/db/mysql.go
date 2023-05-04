package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"inventory/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnectionMysql(cfg *config.Config) (*sql.DB, error) {
	if cfg.DbMain == "" {
		return nil, errors.New("error")
	}

	connStrings := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DbUsername, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)

	db, err := sql.Open(cfg.DbMain, connStrings)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}

func NewMysqlContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Minute)
}
