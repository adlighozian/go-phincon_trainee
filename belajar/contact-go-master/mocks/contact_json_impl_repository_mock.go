package mocks

import (
	"contact-go/model"
	"github.com/stretchr/testify/mock"
)

type JsonRepoMock struct {
	mock.Mock
}

func NewJsonRepoMock() *JsonRepoMock {
	return &JsonRepoMock{}
}

func (m *JsonRepoMock) UpdateJSON(list []model.Contact) (err error) {
	return nil
}

func (m *JsonRepoMock) List() (result []model.Contact, err error) {
	ret := m.Called()
	result = ret.Get(0).([]model.Contact)
	return result, nil
}
