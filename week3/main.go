package main

import (
	"context"
	"fmt"
	"time"
	"week3/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type listContact struct {
	name   string
	noTelp string
}

func main() {

	// var listContacts []listContact = []listContact{
	// 	{
	// 		name:   "adli",
	// 		noTelp: "123221",
	// 	},
	// 	{
	// 		name:   "agung",
	// 		noTelp: "54521",
	// 	},
	// 	{
	// 		name:   "aris",
	// 		noTelp: "52351235",
	// 	},
	// }

	db := db.GetConnection()

	// ctx := context.Background()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `INSERT INTO client(name,no_telp) VALUES ('Test1',NULL)`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	query = `SELECT id, name, no_telp FROM client`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name, noTelp string
		err := rows.Scan(&id, &name, &noTelp)
		if err != nil {
			panic(err)
		}
		fmt.Println("id: ", id)
		fmt.Println("name: ", name)
		fmt.Println("no telephone: ", noTelp)
	}

	// query = `INSERT INTO client(name,no_telp) VALUES ('?','?')`
	// rows, err := db.QueryContext(ctx, query)
	// for i, v := range listContacts {

	// }

	defer db.Close()
	fmt.Println("Success input database")
}
