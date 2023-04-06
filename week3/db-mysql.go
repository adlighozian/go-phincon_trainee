package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"week3/db"
	"week3/model"

	_ "github.com/go-sql-driver/mysql"
)

func mysql() {
	db := db.GetConnectionMysql()

	MakeSlice(2)

	// insert exec
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	nama := "tidak"
	no_telp := Randomizer()

	query := `INSERT INTO client(nama,no_telp) VALUES (?,?)`

	result, err := db.ExecContext(ctx, query, nama, no_telp)
	if err != nil {
		panic(err)
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println(lastInsertId)

	// prepare statment
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	for _, v := range model.Clients {
		result, err := stmt.ExecContext(ctx, v.Nama, v.No_telp)
		if err != nil {
			panic(err)
		}
		lastUserId, _ := result.LastInsertId()
		fmt.Println(lastUserId)
	}

	// transaction
	txr, _ := db.BeginTx(ctx, nil)
	stmts, _ := txr.PrepareContext(ctx, query)
	var (
		errs       error
		lastUserId int64
	)
	for _, v := range model.Clients {
		result, err := stmts.ExecContext(ctx, v.Nama, v.No_telp)
		if err != nil {
			panic(err)
		}
		lastUserId, errs = result.LastInsertId()
		fmt.Println(lastUserId)
	}
	if errs == nil {
		txr.Commit()
	} else {
		txr.Rollback()
	}

	// Select query
	query = `SELECT id, nama, no_telp FROM client`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, nama, noTelp string
		err := rows.Scan(&id, &nama, &noTelp)
		if err != nil {
			panic(err)
		}
		// fmt.Println("id: ", id, "nama: ", nama, "no telp: ", noTelp)
	}

	defer db.Close()
}

func Randomizer() string {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	letters := []rune("1234567890")

	b := make([]rune, 7)

	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	rand := string(b)

	return rand
}

func RandomizerName() string {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	letters := []rune("qwertyuioplkjhgfdsazxcvbnm")

	b := make([]rune, 7)

	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	rand := string(b)

	return rand
}

func MakeSlice(param int) model.Client {
	var slice model.Client

	for i := 1; i <= param; i++ {
		time.Sleep(1 * time.Second)
		slice = model.Client{
			Nama:    RandomizerName(),
			No_telp: Randomizer(),
		}
		model.Clients = append(model.Clients, slice)
		fmt.Println(slice)
	}

	return slice
}
