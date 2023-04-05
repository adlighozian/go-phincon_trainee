package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func getConnectionShow() *sql.DB {
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

func test() {
	db := getConnectionShow()
	defer db.Close()

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT id, name FROM customer")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var test string
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		if id == "adli" {
			test = id
		}
		fmt.Println("id: ", id)
		fmt.Println("name: ", name)
	}

	fmt.Println("Success show database database")
	fmt.Println(test)
}
