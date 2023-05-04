package mocks

import (
	"contact-go/model"
	"github.com/stretchr/testify/mock"
)

type UsecaseMock struct {
	mock.Mock
}

func NewUseCaseMock() *UsecaseMock {
	return &UsecaseMock{}
}

func (m *UsecaseMock) List() (model.ContactResponse, error) {
	ret := m.Called()
	result := ret.Get(0).(model.ContactResponse)
	err := ret.Error(1)
	return result, err
}

func (m *UsecaseMock) Add(req []model.ContactRequest) (model.ContactResponse, error) {
	ret := m.Called(req)
	result := ret.Get(0).(model.ContactResponse)
	err := ret.Error(1)
	return result, err
}

func (m *UsecaseMock) Update(idStr string, req model.ContactRequest) (model.ContactResponse, error) {
	ret := m.Called(idStr, req)
	result := ret.Get(0).(model.ContactResponse)
	err := ret.Error(1)
	return result, err
}

func (m *UsecaseMock) Delete(idStr string) (model.ContactResponse, error) {
	ret := m.Called(idStr)
	result := ret.Get(0).(model.ContactResponse)
	err := ret.Error(1)
	return result, err
}