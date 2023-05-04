package db

import (
	"database/sql"
	"fmt"
	"inventory/config"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	config := config.LoadConfig()
	connStrings := fmt.Sprintf("postgres://%s:%s@%s:%v/%v", config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	db, err := sql.Open(config.DbMain, connStrings)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
