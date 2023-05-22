package product

import (
	"fmt"
	"inventory/helper/middleware"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repoMock struct {
	suite.Suite
	mock    sqlmock.Sqlmock
	DB      *gorm.DB
	product ProductRepository
}

// this function executes before the test suite begins execution
func (s *repoMock) SetupSuite() {

	db, mock, err := sqlmock.New()
	middleware.FailError(err, "Error database connection")
	s.mock = mock

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	middleware.FailError(err, "Error database connection")
	s.DB = gormDB

	s.product = NewProductRepository(s.DB)
}

// this function executes after all tests executed
func (s *repoMock) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}

// this function executes before each test case
func (s *repoMock) SetupTest() {
	// reset StartingNumber to one
	fmt.Println("-- From SetupTest")
}

// this function executes after each test case
func (s *repoMock) TearDownTest() {
	fmt.Println("-- From TearDownTest")
}

// repository Product start
func (s *repoMock) TestShowProduct_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "price", "stock"}).
		AddRow(1, "Kaos Phincon", 30000, 10).
		AddRow(2, "Lanyard_Phincon", 30000, 15)

	query := "SELECT * FROM products"
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	data, err := s.product.ShowProduct()

	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), data)
}

func (s *repoMock) TestShowProduct_Error1() {
	rows := sqlmock.NewRows([]string{"id", "name", "price", "stock"})

	query := "SELECT * FROM products"
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	data, err := s.product.ShowProduct()

	require.Error(s.T(), err)
	require.Empty(s.T(), data)
}

func TestRepoHTTP(t *testing.T) {
	suite.Run(t, new(repoMock))
}
