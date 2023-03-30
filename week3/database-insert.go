package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func getConnection() *sql.DB {
	connString := "root:@tcp(localhost:3306)/golang-trainee"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}

func main() {
	db := getConnection()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO customer(id, name) VALUES ('adli','Adli')")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Success input database")
}
