package main

import (
	"week3/config"
)

func main() {
	config := config.LoadConfig()

	switch {
	case config.DbStart == "mysql":
		mysql()
	case config.DbStart == "psql":
		psql()
	}

}
