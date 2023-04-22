package client

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"sales-go/config"
	logger "sales-go/helpers/logging"
)

var (
	_ = godotenv.Load()
	conDB 				= config.NewConfig()
	connString string 	= ""
	Database 			= os.Getenv("DATABASE")
)

type dbOption struct {
	Database string
}

func NewConnection(database string) dbOption {
	return dbOption{
		Database: database,
	}
}

func (dbOpt dbOption) GetMysqlConnection() (db *sql.DB) {
	var driver string
	if dbOpt.Database == "mysql" {
		// format : "username:password@tcp(host:port)/database_name"
		driver = "mysql"
		connString = fmt.Sprintf("%s:%s@tcp(%s:%v)/%v", conDB.MySQL.Username, conDB.MySQL.Password, conDB.MySQL.Host, conDB.MySQL.Port, conDB.MySQL.Database)
	} else if dbOpt.Database == "postgresql" {
		// format : "postgres://username:password@localhost:5432/database_name"
		driver = "pgx"
		connString = fmt.Sprintf("postgres://%s:%s@%s:%v/%v", conDB.PostgreSQL.Username, conDB.PostgreSQL.Password, conDB.PostgreSQL.Host, conDB.PostgreSQL.Port, conDB.PostgreSQL.Database)
		fmt.Println(connString)
	}

	db, err := sql.Open(driver, connString)
	if err != nil {
		logger.Errorf(fmt.Errorf("unable to connect to database: %v", err))
		panic(err)
	}

	if dbOpt.Database == "mysql" {
		logger.Infof(fmt.Sprintf("Running %s on %s on port %s\n", dbOpt.Database, conDB.MySQL.Host, conDB.MySQL.Port))
	} else if dbOpt.Database == "postgresql" {
		logger.Infof(fmt.Sprintf("Running %s on %s on port %s\n", dbOpt.Database, conDB.PostgreSQL.Host, conDB.PostgreSQL.Port))
	} 

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(5)
	db.SetConnMaxIdleTime(10*time.Minute)
	db.SetConnMaxLifetime(60*time.Minute)

	return
}