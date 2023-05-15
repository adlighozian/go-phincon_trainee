package repository

import (
	"contact-go/db"
	"contact-go/model"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type contactRepository struct {
	Conn *sql.DB
}

func NewContactRepository(connection *sql.DB) ContactRepository {
	return &contactRepository{
		Conn: connection,
	}
}

func (repo *contactRepository) List() ([]model.Client, error) {
	// defer repo.Conn.Close()
	log.Println("list repository")
	ctx, cancel := db.NewMysqlContext()
	defer cancel()

	var data []model.Client
	query := `SELECT * FROM clients`
	rows, err := repo.Conn.QueryContext(ctx, query)
	if err != nil {
		log.Println("list repository error query")
		log.Println(err)
		return data, errors.New("error")
	}

	var temp model.Client
	for rows.Next() {
		rows.Scan(&temp.Id, &temp.Name, &temp.NoTelp)
		data = append(data, temp)
	}

	return data, nil

}

func (repo *contactRepository) Add(req []model.ContactRequest) ([]model.Client, error) {
	// defer repo.Conn.Close()
	ctx, cancel := db.NewMysqlContext()
	defer cancel()

	var contacts []model.Client

	query := `INSERT INTO clients (name, no_telp) value (?,?)`
	txr, err := repo.Conn.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return contacts, errors.New("error")
	}

	stmt, err := txr.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return contacts, errors.New("error")
	}

	defer stmt.Close()

	for _, v := range req {

		result, err := stmt.ExecContext(ctx, v.Name, v.NoTelp)
		if err != nil {
			log.Println(err)
			txr.Rollback()
			return contacts, errors.New("error")
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			log.Println(err)
			return contacts, errors.New("error")
		}

		fmt.Println("id:", lastInsertId)

		contacts = append(contacts, model.Client{
			Id:     int(lastInsertId),
			Name:   v.Name,
			NoTelp: v.NoTelp,
		})
	}

	txr.Commit()

	return contacts, nil

}

func (repo *contactRepository) Update(id int, req model.ContactRequest) error {
	// defer repo.Conn.Close()
	ctx, cancel := db.NewMysqlContext()
	defer cancel()

	query := `UPDATE clients SET name = ?, no_telp = ? WHERE id = ?`
	trx, err := repo.Conn.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return errors.New("error")
	}

	stmt, err := trx.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return errors.New("error")
	}

	_, err = stmt.ExecContext(ctx, req.Name, req.NoTelp, id)
	if err != nil {
		trx.Rollback()
		log.Println(err)
		return errors.New("error")
	}

	trx.Commit()

	return nil

}

func (repo *contactRepository) Delete(id int) error {
	// defer repo.Conn.Close()
	ctx, cancel := db.NewMysqlContext()
	defer cancel()

	query := `DELETE FROM clients WHERE id = ?`
	_, err := repo.Conn.ExecContext(ctx, query, id)
	if err != nil {
		log.Println(err)
		return errors.New("gagal delete")
	}

	return nil
}
