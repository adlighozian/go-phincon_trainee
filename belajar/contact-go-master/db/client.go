package client

import (
	"fmt"
	"log"
	"time"

	"contact-go/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql	   = "mysql"
	postgresql = "postgres"
)

type dbOption struct {
	Database string
}

func GetDB(dbSelected string) dbOption {
	return dbOption{
		Database: dbSelected,
	}
}

var (
	conDB 			  = config.GetConf()
	connString string = ""
)

func (dbOpt dbOption) GetMysqlConnection() (db *sql.DB) {
	if dbOpt.Database == mysql {
		// "username:password@tcp(host:port)/database_name"
		connString = fmt.Sprintf("%s:%s@tcp(%s:%v)/%v", conDB.Mysqlconf.Username, conDB.Mysqlconf.Password, conDB.Mysqlconf.Host, conDB.Mysqlconf.Port, conDB.Mysqlconf.Database)
	}

	db, err := sql.Open(mysql, connString)
	if err != nil {
		panic(err)
	}

	log.Printf("Running mysql on %s on port %s\n", conDB.Mysqlconf.Host, conDB.Mysqlconf.Port)
	
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(5)
	db.SetConnMaxIdleTime(10*time.Minute)
	db.SetConnMaxLifetime(60*time.Minute)

	return
}
