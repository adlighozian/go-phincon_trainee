package repository

import (
	"contact-go/config/db"
	"contact-go/model"
	"database/sql"
)

type contactMysqlRepository struct {
	db *sql.DB
}

func NewContactMysqlRepository(db *sql.DB) ContactRepository {
	return &contactMysqlRepository{
		db: db,
	}
}

// db.ExecContext(...) function is used for executing SQL statements
// that do not return any rows, such as INSERT, UPDATE, and DELETE statements.

// On the other hand, the db.QueryRowContext(...) function is used for
// executing SQL queries that return a single row of result set.

func (repo *contactMysqlRepository) List() ([]model.Contact, error) {
	var contacts []model.Contact
	var contact model.Contact
	var err error

	ctx, cancel := db.NewMysqlContext()
	defer cancel()

	sqlQuery := "SELECT id, name, no_telp FROM contact ORDER BY id ASC"
	rows, err := repo.db.QueryContext(ctx, sqlQuery)
	if err != nil {
		return contacts, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&contact.ID, &contact.Name, &contact.NoTelp)
		if err != nil {
			return contacts, err
		}

		contacts = append(contacts, contact)
	}

	err = rows.Err()
	if err != nil {
		return contacts, err
	}

	return contacts, nil
}

func (repo *contactMysqlRepository) Add(contact *model.Contact) (*model.Contact, error) {
	newContact := new(model.Contact)
	var err error

	ctx, cancel := db.NewMysqlContext()
	defer cancel()

	sqlQuery1 := `
	INSERT INTO contact(name, no_telp) 
	VALUES (?, ?)
	`

	sqlQuery2 := `
	SELECT id, name, no_telp
	FROM contact WHERE id = ? 
	LIMIT 1
	`

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	stmt1, err := tx.PrepareContext(ctx, sqlQuery1)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	row1, err := stmt1.ExecContext(ctx, contact.Name, contact.NoTelp)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	id, err := row1.LastInsertId()
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	stmt2, err := tx.PrepareContext(ctx, sqlQuery2)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	row2 := stmt2.QueryRowContext(ctx, id)
	err = row2.Scan(&newContact.ID, &newContact.Name, &newContact.NoTelp)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	return newContact, nil
}

func (repo *contactMysqlRepository) Detail(id int64) (*model.Contact, error) {
	contact := new(model.Contact)
	var err error

	ctx, cancel := db.NewMysqlContext()
	defer cancel()

	sqlQuery := "SELECT id, name, no_telp FROM contact WHERE id = ? LIMIT 1"
	row := repo.db.QueryRowContext(ctx, sqlQuery, id)
	err = row.Scan(&contact.ID, &contact.Name, &contact.NoTelp)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (repo *contactMysqlRepository) Update(id int64, contact *model.Contact) (*model.Contact, error) {
	updatedContact := new(model.Contact)
	var err error

	ctx, cancel := db.NewMysqlContext()
	defer cancel()

	sqlQuery1 := `
	UPDATE contact SET name = ?, no_telp = ? 
	WHERE id = ?
	`

	sqlQuery2 := `
	SELECT id, name, no_telp
	FROM contact WHERE id = ? 
	LIMIT 1
	`

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	stmt1, err := tx.PrepareContext(ctx, sqlQuery1)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	row1, err := stmt1.ExecContext(ctx, contact.Name, contact.NoTelp, id)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	contactId, err := row1.LastInsertId()
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	stmt2, err := tx.PrepareContext(ctx, sqlQuery2)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	row2 := stmt2.QueryRowContext(ctx, contactId)
	err = row2.Scan(&updatedContact.ID, &updatedContact.Name, &updatedContact.NoTelp)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	return updatedContact, nil
}

func (repo *contactMysqlRepository) Delete(id int64) error {
	ctx, cancel := db.NewMysqlContext()
	defer cancel()

	sqlQuery := "DELETE FROM contact WHERE id = ?"
	_, err := repo.db.ExecContext(ctx, sqlQuery, id)
	if err != nil {
		return err
	}

	return nil
}
