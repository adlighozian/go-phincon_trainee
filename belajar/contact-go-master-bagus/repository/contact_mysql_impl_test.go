package repository

import (
	"contact-go/model"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MysqlRepoSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSQL sqlmock.Sqlmock
	repo    ContactRepository
}

func (s *MysqlRepoSuite) SetupTest() {
	var err error

	db, mock, err := sqlmock.New()
	if err != nil {
		s.Require().NoError(err)
	}

	repo := NewContactMysqlRepository(db)

	s.mockDB = db
	s.mockSQL = mock
	s.repo = repo
}

func (s *MysqlRepoSuite) TearDownTest() {
	s.mockDB.Close()
}

func TestMysqlRepoSuite(t *testing.T) {
	suite.Run(t, new(MysqlRepoSuite))
}

func (s *MysqlRepoSuite) Test_contactMysqlRepository_List() {
	tests := []struct {
		name       string
		beforeTest func(sqlmock.Sqlmock)
		want       []model.Contact
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "no_telp"}).
					AddRow(1, "test", "555-555-3232")

				s.ExpectQuery("SELECT id, name, no_telp FROM contact ORDER BY id ASC").
					WillReturnRows(rows)
			},
			want: []model.Contact{
				{
					ID:     1,
					Name:   "test",
					NoTelp: "555-555-3232",
				},
			},
			wantErr: false,
		},
		{
			name: "fail",
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT id, name, no_telp FROM contact ORDER BY id ASC").
					WillReturnError(assert.AnError)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			if tt.beforeTest != nil {
				tt.beforeTest(s.mockSQL)
			}

			got, err := s.repo.List()

			if s.Equal(tt.wantErr, err != nil, "contactMysqlRepository.List() error = %v, wantErr %v", err, tt.wantErr) {
				s.Equal(tt.want, got, "contactMysqlRepository.List() = %v, want %v", got, tt.want)
			}

			if err := s.mockSQL.ExpectationsWereMet(); err != nil {
				s.Errorf(err, "there were unfulfilled expectations: %s")
			}
		})
	}
}

func (s *MysqlRepoSuite) Test_contactMysqlRepository_Add() {
	type args struct {
		contact *model.Contact
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       *model.Contact
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				contact: &model.Contact{
					Name:   "test",
					NoTelp: "555-555-3232",
				},
			},
			beforeTest: func(s sqlmock.Sqlmock) {
				sqlQuery1 := `
				INSERT INTO contact(name, no_telp) 
				VALUES (?, ?)
				`

				sqlQuery2 := `
				SELECT id, name, no_telp
				FROM contact WHERE id = ? 
				LIMIT 1
				`

				result := sqlmock.NewResult(1, 1)

				rows := sqlmock.NewRows([]string{"id", "name", "no_telp"}).
					AddRow(1, "test", "555-555-3232")

				s.ExpectBegin()
				s.ExpectPrepare(regexp.QuoteMeta(sqlQuery1)).
					WillBeClosed().
					ExpectExec().
					WithArgs("test", "555-555-3232").
					WillReturnResult(result)

				id, _ := result.LastInsertId()

				s.ExpectPrepare(regexp.QuoteMeta(sqlQuery2)).
					WillBeClosed().
					ExpectQuery().
					WithArgs(id).
					WillReturnRows(rows)

				s.ExpectCommit()
			},
			want: &model.Contact{
				ID:     1,
				Name:   "test",
				NoTelp: "555-555-3232",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			if tt.beforeTest != nil {
				tt.beforeTest(s.mockSQL)
			}

			got, err := s.repo.Add(tt.args.contact)

			if s.Equal(tt.wantErr, err != nil, "contactMysqlRepository.Add() error = %v, wantErr %v", err, tt.wantErr) {
				s.Equal(tt.want, got, "contactMysqlRepository.Add() = %v, want %v", got, tt.want)
			}

			if err := s.mockSQL.ExpectationsWereMet(); err != nil {
				s.Errorf(err, "there were unfulfilled expectations: %s")
			}
		})
	}
}

func (s *MysqlRepoSuite) Test_contactMysqlRepository_Detail() {
	type args struct {
		id int64
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       *model.Contact
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				id: 1,
			},
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "no_telp"}).
					AddRow(1, "test", "555-555-3232")

				s.ExpectQuery(regexp.QuoteMeta("SELECT id, name, no_telp FROM contact WHERE id = ? LIMIT 1")).
					WillReturnRows(rows)
			},
			want: &model.Contact{
				ID:     1,
				Name:   "test",
				NoTelp: "555-555-3232",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			if tt.beforeTest != nil {
				tt.beforeTest(s.mockSQL)
			}

			got, err := s.repo.Detail(tt.args.id)

			if s.Equal(tt.wantErr, err != nil, "contactMysqlRepository.Detail() error = %v, wantErr %v", err, tt.wantErr) {
				s.Equal(tt.want, got, "contactMysqlRepository.Detail() = %v, want %v", got, tt.want)
			}

			if err := s.mockSQL.ExpectationsWereMet(); err != nil {
				s.Errorf(err, "there were unfulfilled expectations: %s")
			}
		})
	}
}

func (s *MysqlRepoSuite) Test_contactMysqlRepository_Update() {
	type args struct {
		id      int64
		contact *model.Contact
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       *model.Contact
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				id: 1,
				contact: &model.Contact{
					Name:   "jangkrik",
					NoTelp: "555-555-4000",
				},
			},
			beforeTest: func(s sqlmock.Sqlmock) {
				sqlQuery1 := `
				UPDATE contact SET name = ?, no_telp = ? 
				WHERE id = ?
				`

				sqlQuery2 := `
				SELECT id, name, no_telp
				FROM contact WHERE id = ? 
				LIMIT 1
				`

				result := sqlmock.NewResult(1, 1)

				rows := sqlmock.NewRows([]string{"id", "name", "no_telp"}).
					AddRow(1, "jangkrik", "555-555-4000")

				s.ExpectBegin()
				s.ExpectPrepare(regexp.QuoteMeta(sqlQuery1)).
					WillBeClosed().
					ExpectExec().
					WithArgs("jangkrik", "555-555-4000", 1).
					WillReturnResult(result)

				id, _ := result.LastInsertId()

				s.ExpectPrepare(regexp.QuoteMeta(sqlQuery2)).
					WillBeClosed().
					ExpectQuery().
					WithArgs(id).
					WillReturnRows(rows)

				s.ExpectCommit()
			},
			want: &model.Contact{
				ID:     1,
				Name:   "jangkrik",
				NoTelp: "555-555-4000",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			if tt.beforeTest != nil {
				tt.beforeTest(s.mockSQL)
			}

			got, err := s.repo.Update(tt.args.id, tt.args.contact)

			if s.Equal(tt.wantErr, err != nil, "contactMysqlRepository.Update() error = %v, wantErr %v", err, tt.wantErr) {
				s.Equal(tt.want, got, "contactMysqlRepository.Update() = %v, want %v", got, tt.want)
			}

			if err := s.mockSQL.ExpectationsWereMet(); err != nil {
				s.Errorf(err, "there were unfulfilled expectations: %s")
			}
		})
	}
}

func (s *MysqlRepoSuite) Test_contactMysqlRepository_Delete() {
	type args struct {
		id int64
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				id: 1,
			},
			beforeTest: func(s sqlmock.Sqlmock) {
				result := sqlmock.NewResult(1, 1)

				s.ExpectExec(regexp.QuoteMeta("DELETE FROM contact WHERE id = ?")).
					WillReturnResult(result)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			if tt.beforeTest != nil {
				tt.beforeTest(s.mockSQL)
			}

			err := s.repo.Delete(tt.args.id)

			s.Equal(tt.wantErr, err != nil, "contactUsecase.Delete() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}
