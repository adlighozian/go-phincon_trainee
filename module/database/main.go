package main

import (
	"database/config"
	"database/repository"
)

func main() {
	config := config.LoadConfig()

	switch {
	case config.DbStart == "mysql":
		repository.Mysql()
	case config.DbStart == "psql":
		repository.Psql()
	}

}
