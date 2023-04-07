package repository

import (
	"contact-go/config"
	"contact-go/db"
	"contact-go/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type contactRepository struct{}

func NewContactRepository() ContactRepository {
	return new(contactRepository)
}

func (repo *contactRepository) DecodeJson() []model.Contact {
	reader, err := os.Open("./assets/contacts.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(reader)
	decoder.Decode(&model.Contacts)
	return model.Contacts
}

func (repo *contactRepository) EncodeJson() {
	writer, err := os.Create("./assets/contacts.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(writer)
	encoder.Encode(repo.DecodeJson())
}

func (repo *contactRepository) GetLastID() int {
	contacts := repo.List()

	var tempID int
	for _, v := range contacts {
		if tempID < v.Id {
			tempID = v.Id
		}
	}
	return tempID
}

func (repo *contactRepository) GetIndexByID(id int) (int, error) {
	contacts := repo.List()

	for i, v := range contacts {
		if id == v.Id {
			return i, nil
		}
	}

	return -1, errors.New("ID tidak ditemukan")
}

func (repo *contactRepository) List() (data []model.Contact) {

	switch {
	case config.LoadConfig().Storage == "db":
		db := db.GetConnectionMysql()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		query := `SELECT * FROM client`
		rows, err := db.QueryContext(ctx, query)
		if err != nil {
			panic(err)
		}

		var temp model.Contact
		for rows.Next() {
			rows.Scan(&temp.Id, &temp.Name, &temp.NoTelp)
			data = append(data, temp)
		}
		return data
	case config.LoadConfig().Storage == "json":
		return repo.DecodeJson()
	}

	return
}

func (repo *contactRepository) Add(req []model.ContactRequest) ([]model.Contact, error) {

	db := db.GetConnectionMysql()
	var contacts []model.Contact
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	query := `INSERT INTO client(nama,no_telp) VALUES (?,?)`

	txr, errs := db.BeginTx(ctx, nil)
	stmt, _ := txr.PrepareContext(ctx, query)
	defer stmt.Close()

	for _, v := range req {
		result, err := stmt.ExecContext(ctx, v.Name, v.NoTelp)
		if err != nil {
			panic(err)
		}
		lastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", lastInsertId)

		contact := model.Contact{
			Id:     int(lastInsertId),
			Name:   v.Name,
			NoTelp: v.NoTelp,
		}
		contacts = append(contacts, contact)
	}

	if errs != nil {
		txr.Rollback()
	} else {
		txr.Commit()
	}

	return contacts, nil
}

func (repo *contactRepository) Update(id int, req model.ContactRequest) (model.Contact, error) {

	db := db.GetConnectionMysql()
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	contact := model.Contact{
		Id:     id,
		Name:   req.Name,
		NoTelp: req.NoTelp,
	}

	query := `UPDATE client SET nama = ?, no_telp = ? WHERE id = ?`

	_, err := db.ExecContext(ctx, query, req.Name, req.NoTelp, id)
	if err != nil {
		fmt.Println("error update", id, req.Name, req.NoTelp)
		return contact, errors.New("gagal update")
	}

	fmt.Println("Berhasil di update dengan id:", id)
	return contact, nil

}

func (repo *contactRepository) Delete(id int) (int, error) {
	db := db.GetConnectionMysql()
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	query := `DELETE FROM client WHERE id = ?`
	_, err := db.ExecContext(ctx, query, id)

	if err != nil {
		fmt.Println("error delete", id)
		return id, errors.New("gagal update")
	}
	return id, nil
}
