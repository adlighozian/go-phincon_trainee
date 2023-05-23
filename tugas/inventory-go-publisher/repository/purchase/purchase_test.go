package purchase

import (
	"fmt"
	"inventory/helper/middleware"
	"inventory/mocks"
	"inventory/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repoMock struct {
	suite.Suite
	mock          sqlmock.Sqlmock
	DB            *gorm.DB
	repo          PurchaseRepository
	mockPubsliher *mocks.PublisherMock
	mockRandom    *mocks.RandomMock
}

// this function executes before the test suite begins execution
func (s *repoMock) SetupSuite() {

	db, mock, err := sqlmock.New()
	middleware.FailError(err, "Error database connection")
	s.mock = mock

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	middleware.FailError(err, "Error database connection")
	s.DB = gormDB

	s.mockRandom = mocks.NewRandom()
	s.mockPubsliher = mocks.NewPublisher()

	s.repo = NewPurchaseRepository(s.DB, s.mockPubsliher, s.mockRandom)
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
	s.mockRandom.On("Randomizer").Return("548262741")
	s.mockPubsliher.On("PubPurchase", mock.Anything).Return(nil)

	query := `select * from purchase p join purchase_detail pd on p.id = pd.purchase_id where order_number = $1`

	row := sqlmock.NewRows([]string{"id"}).AddRow(1)

	// s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(row)
	// s.mock.ExpectCommit()

	res, err := s.repo.InputPurchase(
		[]model.ReqPurchase{
			{
				Item:  "hp",
				Price: 15000,
				From:  "bagas",
				Total: 10,
			},
		},
	)

	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res)
}

func (s *repoMock) TestInputPurchase_Error1() {
	s.mockRandom.On("Randomizer").Return("548262741")
	s.mockPubsliher.On("PubPurchase", mock.Anything).Return(nil)

	query := `select * from purchase p join purchase_detail pd on p.id = pd.purchase_id where order_number = $1`

	row := sqlmock.NewRows([]string{"id"}).AddRow(1)

	// s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(row)
	// s.mock.ExpectCommit()

	res, err := s.repo.InputPurchase(
		[]model.ReqPurchase{
			{
				Item:  "hp",
				Price: 15000,
				From:  "bagas",
				Total: 0,
			},
		},
	)

	require.NoError(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *repoMock) TestDetailPurchase_Success() {

	query := `select id from purchase where order_number = $1`
	row1 := sqlmock.NewRows([]string{"id"}).AddRow(1)

	selectPurchaseDetail := `select * from purchase p join purchase_detail pd on p.id = pd.purchase_id where order_number = $1`
	row2 := sqlmock.NewRows([]string{"id"}).AddRow(1)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(row1)

	s.mock.ExpectQuery(regexp.QuoteMeta(selectPurchaseDetail)).WillReturnRows(row2)

	res, err := s.repo.DetailPurchase("1221")
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res)
}

func (s *repoMock) TestDetailPurchase_Error() {

	query := `select id from purchase where order_number = $1`
	row1 := sqlmock.NewRows([]string{"id"}).AddRow(0)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(row1)

	res, err := s.repo.DetailPurchase("1221")
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func TestRepoHTTP(t *testing.T) {
	suite.Run(t, new(repoMock))
}
