package repository

import (
	"contact-go/model"
	"errors"
	"fmt"
	"regexp"
	"testing"

	// "contact-go/helper/get-env"
	// "github.com/joho/godotenv"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Mysqlclient struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo ContactRepository
}

// set up env
func (client *Mysqlclient) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(fmt.Sprintf("Error database connection %s", err))
	}
	client.mock = mock
	client.repo = NewContactRepository(db)
}

func (client *Mysqlclient) TestGetList_Success() {
	row := sqlmock.NewRows([]string{"id", "name", "notelp"}).AddRow(1, "Andi", "0834234235244").AddRow(2, "Umar", "0894339843943")

	client.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM clients")).WillReturnRows(row)
	list_contact, err := client.repo.List()
	if err != nil {
		client.T().Errorf("error get list contact: %s", err)
	}

	assert.NoError(client.T(), err)
	assert.NotEmpty(client.T(), list_contact)
}

func (client *Mysqlclient) TestGetList_FailedQuery() {

	client.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM clients`)).WillReturnError(errors.New("testing error"))

	list_contact, err := client.repo.List()

	require.Error(client.T(), err)
	require.Empty(client.T(), list_contact)
}

func (client *Mysqlclient) TestAddContact_Success() {
	client.mock.ExpectBegin()
	client.mock.
		ExpectPrepare(regexp.QuoteMeta(`INSERT INTO clients (name, no_telp) value (?,?)`)).
		WillBeClosed().
		ExpectExec().
		WithArgs("Andi", "0884275327327").WillReturnResult(sqlmock.NewResult(1, 1))
	client.mock.ExpectCommit()

	req := []model.ContactRequest{
		{
			Name:   "Andi",
			NoTelp: "0884275327327",
		},
	}

	list_contact, err := client.repo.Add(req)

	fmt.Println("add succese")
	require.NoError(client.T(), err)
	require.NotEmpty(client.T(), list_contact)
}

func (client *Mysqlclient) TestAddContact_FailedTransaction() {
	client.mock.ExpectBegin().WillReturnError(errors.New("testing error"))

	req := []model.ContactRequest{
		{
			Name:   "Andi",
			NoTelp: "0884275327327",
		},
	}

	list_contact, err := client.repo.Add(req)

	require.Error(client.T(), err)
	require.Empty(client.T(), list_contact)
}

func (client *Mysqlclient) TestAddContact_FailedStatement() {
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO clients (name, no_telp) value (?,?)`)).WillReturnError(errors.New("testing error"))

	req := []model.ContactRequest{
		{
			Name:   "Andi",
			NoTelp: "0884275327327",
		},
	}

	list_contact, err := client.repo.Add(req)

	require.Error(client.T(), err)
	require.Empty(client.T(), list_contact)
}

func (client *Mysqlclient) TestAddContact_FailedExec() {
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO clients (name, no_telp) value (?,?)`)).WillBeClosed().ExpectExec().WillReturnError(errors.New("testing error"))

	req := []model.ContactRequest{
		{
			Name:   "Andi",
			NoTelp: "0884275327327",
		},
	}

	list_contact, err := client.repo.Add(req)

	require.Error(client.T(), err)
	require.Empty(client.T(), list_contact)
}

func (client *Mysqlclient) TestAddContact_FailedLastId() {
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO clients (name, no_telp) value (?,?)`)).WillBeClosed().ExpectExec().WithArgs("Andi", "0884275327327").WillReturnResult(sqlmock.NewErrorResult(errors.New("testing error")))

	req := []model.ContactRequest{
		{
			Name:   "Andi",
			NoTelp: "0884275327327",
		},
	}

	list_contact, err := client.repo.Add(req)

	require.Error(client.T(), err)
	require.Empty(client.T(), list_contact)
}

func (client *Mysqlclient) TestUpdateContact_Success() {
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("UPDATE clients SET name = ?, no_telp = ? WHERE id = ?")).
		WillBeClosed().
		ExpectExec().
		WithArgs("Andi", "0884275327327", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	client.mock.ExpectCommit()

	id := 1
	req := model.ContactRequest{
		Name:   "Andi",
		NoTelp: "0884275327327",
	}
	err := client.repo.Update(id, req)

	require.NoError(client.T(), err)
}

func (client *Mysqlclient) TestUpdateContact_FailedTransaction() {
	client.mock.ExpectBegin().WillReturnError(errors.New("testing error"))

	id := 1
	req := model.ContactRequest{
		Name:   "Andi",
		NoTelp: "0884275327327",
	}
	err := client.repo.Update(id, req)

	require.Error(client.T(), err)
}

func (client *Mysqlclient) TestUpdateContact_FailedStatment() {
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta(`UPDATE clients SET name = ?, no_telp = ? WHERE id = ?`)).WillReturnError(errors.New("testing error"))

	id := 1
	req := model.ContactRequest{
		Name:   "Andi",
		NoTelp: "0884275327327",
	}
	err := client.repo.Update(id, req)

	require.Error(client.T(), err)
}

func (client *Mysqlclient) TestUpdateContact_FailedExec() {
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta(`UPDATE clients SET name = ?, no_telp = ? WHERE id = ?`)).WillBeClosed().ExpectExec().WithArgs("Andi", "0884275327327", 1).WillReturnError(errors.New("testing error"))

	id := 1
	req := model.ContactRequest{
		Name:   "Andi",
		NoTelp: "0884275327327",
	}
	err := client.repo.Update(id, req)

	require.Error(client.T(), err)
}

func (client *Mysqlclient) TestDeleteContact_Success() {
	client.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM clients WHERE id = ?")).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	id := 1
	err := client.repo.Delete(id)

	require.NoError(client.T(), err)
}

func (client *Mysqlclient) TestDeleteContact_FailContext() {
	client.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM clients WHERE id = ?")).WithArgs(1).WillReturnError(errors.New("testing error"))

	id := 1
	err := client.repo.Delete(id)

	require.Error(client.T(), err)
}

func TestRepoHTTP(t *testing.T) {
	suite.Run(t, new(Mysqlclient))
}
