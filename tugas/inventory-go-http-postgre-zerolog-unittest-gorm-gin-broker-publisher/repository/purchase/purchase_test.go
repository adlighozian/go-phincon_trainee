package purchase

import (
	"fmt"
	"inventory/helper/middleware"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repoMock struct {
	suite.Suite
	mock sqlmock.Sqlmock
	DB   *gorm.DB
	repo PurchaseRepository
}

// this function executes before the test suite begins execution
func (s *repoMock) SetupSuite() {

	db, mock, err := sqlmock.New()
	middleware.FailError(err, "Error database connection")
	s.mock = mock

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	middleware.FailError(err, "Error database connection")
	s.DB = gormDB

	s.repo = NewPurchaseRepository(s.DB)
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
func (s *repoMock) TestInputPurchase_Success() {

	// res, err := s.repo.InputPurchase()
}

func TestRepoHTTP(t *testing.T) {
	suite.Run(t, new(repoMock))
}
