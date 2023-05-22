package db

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	// config := config.LoadConfig()

	dsn := "host=localhost user=postgres password=admin dbname=go-inventory port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
