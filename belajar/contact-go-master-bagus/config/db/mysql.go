package db

import (
	"contact-go/config"
	"contact-go/helper"
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlDatabase(cfg *config.Config) (*sql.DB, error) {
	if cfg.Database.URL == "" {
		return nil, helper.NewAppError(helper.ErrDbUrlNotExist)
	}

	db, err := sql.Open(cfg.Database.Driver, cfg.Database.URL)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}

func NewMysqlContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
