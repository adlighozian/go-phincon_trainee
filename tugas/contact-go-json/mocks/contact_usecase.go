package mocks

import (
	"contact-go/model"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

func NewContactRepoMock() *RepoMock {
	return &RepoMock{}
}

func (m *RepoMock) List() ([]model.Contact, error) {
	var result []model.Contact
	ret := m.Called()
	result = ret.Get(0).([]model.Contact)
	return result, nil
}

func (m *RepoMock) Add(req []model.ContactRequest) ([]model.Contact, error) {
	var result []model.Contact
	ret := m.Called(req)
	result = ret.Get(0).([]model.Contact)
	return result, nil
}

func (m *RepoMock) Update(id int, req model.ContactRequest) (model.Contact, error) {
	var result model.Contact
	ret := m.Called(id, req)
	var res1 error
	if ret.Get(0) != nil {
		res1 = ret.Get(0).(error)
	}
	return result, res1
}

func (m *RepoMock) Delete(id int) (int, error) {
	ret := m.Called(id)
	var res1 error
	if ret.Get(0) != nil {
		res1 = ret.Get(0).(error)
	}
	return 0, res1
}
