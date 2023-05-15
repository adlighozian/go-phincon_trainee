package mocks

import (
	"contact-go/model"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

func NewRepoMock() *RepoMock {
	return &RepoMock{}
}

func (m *RepoMock) List() ([]model.Client, error) {
	ret := m.Called()
	result := ret.Get(0).([]model.Client)
	err := ret.Error(1)
	return result, err
}

func (m *RepoMock) Add(req []model.ContactRequest) ([]model.Client, error) {
	ret := m.Called(req)
	result := ret.Get(0).([]model.Client)
	err := ret.Error(1)
	return result, err
}

func (m *RepoMock) Update(id int, req model.ContactRequest) error {
	ret := m.Called(id, req)
	err := ret.Error(0)
	return err
}

func (m *RepoMock) Delete(id int) error {
	ret := m.Called(id)
	err := ret.Error(0)
	return err
}
