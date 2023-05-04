package repository

import (
	"contact-go/model"
	"context"
	"database/sql"
	"log"
	"time"
)

type contacthttp struct{
	db *sql.DB
}

func NewContactHTTPRepository(client *sql.DB) *contacthttp {
	return &contacthttp{
		db: client,
	}
}

func (repo *contacthttp) List() (result []model.Contact, err error) {
	defer repo.db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	trx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		log.Print("1")
		return
	}

	// prepare statement
	query := `SELECT id, name, no_telp FROM contact`
	stmt, err := trx.PrepareContext(ctx, query)
	if err != nil {
		log.Print("2")
		return
	}

	res, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Print("3")
		trx.Rollback()
		return
	}

	for res.Next() {
		var temp model.Contact
		res.Scan(&temp.Id, &temp.Name, &temp.NoTelp)
		result = append(result, temp)
	}

	trx.Commit()
	return
}

func (repo *contacthttp) Add(req []model.ContactRequest) (result []model.Contact, err error) {
	defer repo.db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	query := `INSERT INTO contact (name, no_telp) value (?,?)`
	trx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	for _, v := range req {
		res, err := stmt.ExecContext(ctx, v.Name, v.NoTelp)
		if err != nil {
			trx.Rollback()
			return []model.Contact{}, err
		}

		lastID, err := res.LastInsertId()
		if err != nil {
			return []model.Contact{}, err
		}

		result = append(result, model.Contact{
			Id:   	int(lastID),
			Name: 	v.Name,
			NoTelp: v.NoTelp,
		})
	}

	trx.Commit()

	return
}

func (repo *contacthttp) Update(id int, req model.ContactRequest) (err error) {
	defer repo.db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	trx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	query := `UPDATE contact SET name = ?, no_telp = ? WHERE id = ?`
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, req.Name, req.NoTelp, id)
	if err != nil {
		trx.Rollback()
		return
	}

	trx.Commit()
	
	return
}

func (repo *contacthttp) Delete(id int) (err error) {
	defer repo.db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	trx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	query := `DELETE FROM contact WHERE id = ?`
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		trx.Rollback()
	}

	_, err = res.RowsAffected()
	if err != nil {
		return
	}

	trx.Commit()

	return
}
