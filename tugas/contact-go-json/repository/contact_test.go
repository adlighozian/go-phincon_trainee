package repository

import (
	"contact-go/model"
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	// "contact-go/helper/get-env"
	// "github.com/joho/godotenv"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Mysqlclient struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo ContactRepository
}

// set up env
func (client *Mysqlclient) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(fmt.Sprintf("Error database connection %s", err))
	}

	client.db = db
	client.mock = mock
	client.repo = NewContactHTTPRepository(db)
}

func (client *Mysqlclient) TestGetListContact() {
	row := sqlmock.NewRows([]string{"id", "name", "no_telp"}).AddRow(1, "Andi", "0834234235244").AddRow(2, "Umar", "0894339843943")
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare("SELECT id, name, no_telp FROM contact").WillBeClosed().ExpectQuery().WillReturnRows(row)
	client.mock.ExpectCommit()

	list_contact, err := client.repo.List()
	if err != nil {
		client.T().Errorf("error get list contact: %s", err)
	}

	require.NoError(client.T(), err)
	require.NotEmpty(client.T(), list_contact)
}

func (client *Mysqlclient) TestAddContact() {
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO contact (name, no_telp) value (?,?)")).ExpectExec().WithArgs("Andi", "0884275327327").WillReturnResult(sqlmock.NewResult(1, 1))
	client.mock.ExpectCommit()

	req := model.ContactRequest{
		Name:   "Andi",
		NoTelp: "0884275327327",
	}
	list_contact, err := client.repo.Add([]model.ContactRequest{req})
	if err != nil {
		client.T().Errorf("error get list contact: %s", err)
	}

	require.NoError(client.T(), err)
	require.NotEmpty(client.T(), list_contact)
}

func (client *Mysqlclient) TestUpdateContact() {
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("UPDATE contact SET name = ?, no_telp = ? WHERE id = ?")).ExpectExec().WithArgs("Andi", "0884275327327", 1).WillReturnResult(sqlmock.NewResult(1, 1))
	client.mock.ExpectCommit()

	id := 1
	req := model.ContactRequest{
		Name:   "Andi",
		NoTelp: "0884275327327",
	}
	err := client.repo.Update(id, req)
	if err != nil {
		client.T().Errorf("error update contact: %s", err)
	}

	require.NoError(client.T(), err)
}

func (client *Mysqlclient) TestDeleteContact() {
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM contact WHERE id = ?")).ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	client.mock.ExpectCommit()

	id := 1
	err := client.repo.Delete(id)
	if err != nil {
		client.T().Errorf("error delete contact: %s", err)
	}

	require.NoError(client.T(), err)
}

func TestRepoHTTP(t *testing.T) {
	suite.Run(t, new(Mysqlclient))
}
