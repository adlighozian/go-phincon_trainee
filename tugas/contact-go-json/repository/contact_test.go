package repository

import (
	"contact-go/helper"
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MySqlSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
	c    ContactRepository
}

func (s *MySqlSuite) SetupSuite() {
	db_, err := helper.GetEnv("dbms", "mysql")
	if err != nil {
		panic(err)
	}

	s.db, s.mock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}

	s.c = NewContactRepository(s.db, db_)
}

func (s *MySqlSuite) TestContactRepo() {
	rows := sqlmock.NewRows([]string{"id", "nama", "no_telp"}).AddRow(2, "Romo", "98329009238").AddRow(2, "Romo", "98329009238")
	s.mock.ExpectPrepare("SELECT id, name, no_telp FROM contact").WillBeClosed().ExpectQuery().WillReturnRows(rows)

	listed_contact, _ := s.c.List()
	require.NotEmpty(s.T(), listed_contact)
}
