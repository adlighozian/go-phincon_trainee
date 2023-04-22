package config

import (
	"os"
	
	"github.com/joho/godotenv"
)

type sql struct {
	Username string
	Password string
	Host string
	Port string
	Database string
}

type configure struct {
	MySQL	   sql
	PostgreSQL sql
}

func NewConfig() configure {
	godotenv.Load()

	return configure{
		MySQL: sql{
			Username: os.Getenv("MYSQL_USERNAME"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Host:	  os.Getenv("MYSQL_HOST"),
			Port:	  os.Getenv("MYSQL_PORT"),
			Database: os.Getenv("MYSQL_DATABASE"),
		},
		PostgreSQL: sql{
			Username: os.Getenv("POSTGRES_USERNAME"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Host:	  os.Getenv("POSTGRES_HOST"),
			Port:	  os.Getenv("POSTGRES_PORT"),
			Database: os.Getenv("POSTGRES_DATABASE"),
		},
	}
}