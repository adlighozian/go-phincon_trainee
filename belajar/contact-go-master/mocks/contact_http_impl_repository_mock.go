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

func (m *RepoMock) List() (result []model.Contact, err error) {
	ret := m.Called()
	result = ret.Get(0).([]model.Contact)
	err = ret.Error(1)
	return result, err
}

func (m *RepoMock) Add(req []model.ContactRequest) (result []model.Contact, err error) {
	ret := m.Called(req)
	result = ret.Get(0).([]model.Contact)
	err = ret.Error(1)
	return result, err
}

func (m *RepoMock) Update(id int, req model.ContactRequest) (err error) {
	ret := m.Called(id, req)	
	var res1 error
	if ret.Get(0) != nil {
		res1 = ret.Get(0).(error)
	}
	return res1
}

func (m *RepoMock) Delete(id int) (err error) {
	ret := m.Called(id)
	var res1 error
	if ret.Get(0) != nil {
		res1 = ret.Get(0).(error)
	}
	return res1
}