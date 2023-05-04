package repository

import (
	"contact-go/db"
	"contact-go/model"
	"database/sql"
	"errors"
	"fmt"
)

type contactRepository struct {
	Conn *sql.DB
}

func NewContactRepository(connection *sql.DB) ContactRepository {
	return &contactRepository{
		Conn: connection,
	}
}

func (repo *contactRepository) List() ([]model.Contact, error) {

	var data []model.Contact

	ctx, cancel := db.NewMysqlContext()
	defer cancel()

	query := `SELECT * FROM client`
	rows, err := repo.Conn.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	var temp model.Contact
	for rows.Next() {
		rows.Scan(&temp.Id, &temp.Name, &temp.NoTelp)
		data = append(data, temp)
	}

	return data, nil

}

func (repo *contactRepository) Add(req []model.ContactRequest) ([]model.Contact, error) {

	var contacts []model.Contact
	ctx, cancel := db.NewMysqlContext()
	defer cancel()

	query := `INSERT INTO client(nama,no_telp) VALUES (?,?)`

	txr, errs := repo.Conn.BeginTx(ctx, nil)
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

	if contacts == nil {
		return nil, errors.New("gagal menambahkan data")
	} else {
		return contacts, nil
	}

}

func (repo *contactRepository) Update(id int, req model.ContactRequest) (model.Contact, error) {

	defer repo.Conn.Close()
	ctx, cancel := db.NewMysqlContext()
	defer cancel()
	contact := model.Contact{
		Id:     id,
		Name:   req.Name,
		NoTelp: req.NoTelp,
	}

	query := `UPDATE client SET nama = ?, no_telp = ? WHERE id = ?`

	_, err := repo.Conn.ExecContext(ctx, query, req.Name, req.NoTelp, id)
	if err != nil {
		fmt.Println("error update", id, req.Name, req.NoTelp)
		return contact, errors.New("gagal update")
	}

	fmt.Println("Berhasil di update dengan id:", id)
	return contact, nil

}

func (repo *contactRepository) Delete(id int) (int, error) {

	defer repo.Conn.Close()
	ctx, cancel := db.NewMysqlContext()
	defer cancel()

	query := `DELETE FROM client WHERE id = ?`
	_, err := repo.Conn.ExecContext(ctx, query, id)
	if err != nil {
		fmt.Println("error delete 3", id)
		return id, errors.New("gagal delete")
	}

	return id, nil
}
