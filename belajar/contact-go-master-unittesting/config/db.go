package config

import (
	"fmt"
	"os"
)

type mysql struct {
	Username string
	Password string
	Host	 string
	Port	 string
	Database string
}

type postgresql struct {
	Username string
	Password string
	Host	 string
	Port	 string
	Database string
}

type configure struct {
	Mysqlconf		mysql
	Postgresqlconf	postgresql
}

func GetConf() configure {
	os.LookupEnv("MySQLUsername")
	fmt.Println(os.Getenv("MySQLUsername"))

	return configure {
		Mysqlconf: mysql{
			Username: "root",
			Password: "@Ugm428660",
			Host:	  "localhost",
			Port:	  "3306",
			Database: "bootcamp",
		},
		Postgresqlconf: postgresql{
			Username: os.Getenv("PostgreSQLUsername"),
			Password: os.Getenv("PostgreSQLPassword"),
			Host:	  os.Getenv("PostgreSQLHost"),
			Port:	  os.Getenv("PostgreSQLPort"),
			Database: os.Getenv("PostgreSQLDatabase"),
		},
	}
}