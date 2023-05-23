package client

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"sales-go/config"
	"sales-go/helpers/logging"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	_ = godotenv.Load()

	conDB 				= config.NewConfig()
	connString string 	= ""
	Database 			= os.Getenv("DATABASE")
	req					= new(http.Request)
)

type dbOption struct {
	Database string
}

func NewConnection(database string) dbOption {
	return dbOption{
		Database: database,
	}
}

func (dbOpt dbOption) GetDBConnection() (db *sql.DB, err error) {
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

	db, err = sql.Open(driver, connString)
	if err != nil {
		logging.Errorf(fmt.Errorf("unable to connect to database: %v", err), req)
		return
	}

	if dbOpt.Database == "mysql" {
		logging.Infof(fmt.Sprintf("Running %s on %s on port %s\n", dbOpt.Database, conDB.MySQL.Host, conDB.MySQL.Port), req)
	} else if dbOpt.Database == "postgresql" {
		logging.Infof(fmt.Sprintf("Running %s on %s on port %s\n", dbOpt.Database, conDB.PostgreSQL.Host, conDB.PostgreSQL.Port), req)
	} 

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(5)
	db.SetConnMaxIdleTime(10*time.Minute)
	db.SetConnMaxLifetime(60*time.Minute)

	return
}

func (dbOpt dbOption) GetDBGormConnection() (*gorm.DB, error) {
	var connString string
	if dbOpt.Database == "mysql" {
		// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		connString = fmt.Sprintf("%s:%s@tcp(%s:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", 
			conDB.MySQL.Username, 
			conDB.MySQL.Password, 
			conDB.MySQL.Host, 
			conDB.MySQL.Port, 
			conDB.MySQL.Database,
		)
	} else if dbOpt.Database == "postgresql" {
		// "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
		connString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			conDB.PostgreSQL.Host,
			conDB.PostgreSQL.Username,
			conDB.PostgreSQL.Password,
			conDB.PostgreSQL.Database,
			conDB.PostgreSQL.Port,
		)
	}

	db, err := gorm.Open(
		postgres.Open(connString), &gorm.Config{
			Logger: 				logger.Default.LogMode(logger.Info),
			SkipDefaultTransaction: true,
			PrepareStmt:			false,
		},
	)
	if err != nil {
		logging.Errorf(fmt.Errorf("unable to connect to database: %v", err), req)
		return nil, err
	}

	logging.Infof(fmt.Sprintf("Running mysql on %s on port %s\n", conDB.MySQL.Host, conDB.MySQL.Port), req)

	return db, nil
}