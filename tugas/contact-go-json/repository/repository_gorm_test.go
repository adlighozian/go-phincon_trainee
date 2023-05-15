package repository

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormClient struct {
	suite.Suite
	mock  sqlmock.Sqlmock
	repo  ContactRepository
	sqlDB *sql.DB
}

func (suite *GormClient) SetupSuite() {
	// set StartingNumber to one
	fmt.Println(">>> From SetupSuite")
}

// this function executes after all tests executed
func (suite *GormClient) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}

// this function executes before each test case
func (client *GormClient) SetupTest() {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		panic(fmt.Sprintf("Error database connection %s", err))
	}

	db, err := gorm.Open(
		mysql.New(
			mysql.Config{
				DriverName:                "mysql",
				Conn:                      dbMock,
				SkipInitializeWithVersion: true,
			},
		), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic(fmt.Sprintf("Error database connection %s", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		client.Require().NoError(err)
	}

	client.mock = mock
	client.repo = NewContactRepositoryGorm(db)
	client.sqlDB = sqlDB
}

// this function executes after each test case
func (suite *GormClient) TearDownTest() {
	fmt.Println("-- From TearDownTest")
}

func (client *GormClient) TestList_success() {
	log.Panicln("Test list success gorm")
	row := sqlmock.NewRows([]string{"id", "name", "notelp"}).AddRow(1, "Andi", "0834234235244").AddRow(2, "Umar", "0894339843943")

	client.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `clients`")).WillReturnRows(row)
	list_contact, err := client.repo.List()
	if err != nil {
		client.T().Errorf("error get list contact: %s", err)
	}

	assert.NoError(client.T(), err)
	assert.NotEmpty(client.T(), list_contact)

}

func TestRepoGormHTTP(t *testing.T) {
	suite.Run(t, new(GormClient))
}
