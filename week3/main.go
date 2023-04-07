package main

import (
	"week3/config"
	"week3/repository"
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
