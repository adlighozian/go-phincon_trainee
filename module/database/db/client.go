package db

import (
	"database/sql"
	"fmt"
	"time"
	"week3/config"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnectionMysql() *sql.DB {
	config := config.LoadConfig()
	connStrings := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	db, err := sql.Open(config.DbMain, connStrings)
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

func GetConnectionPsql() *sql.DB {
	config := config.LoadConfig()
	connStrings := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
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
