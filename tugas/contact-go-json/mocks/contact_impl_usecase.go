package mocks

import (
	"contact-go/model"

	"github.com/stretchr/testify/mock"
)

type ContactMock struct {
	Mock mock.Mock
}

func (repository *ContactMock) List() ([]model.Contact, error) {
	return nil, nil
}

func (repository *ContactMock) Add(req []model.ContactRequest) ([]model.Contact, error) {
	return nil, nil
}

func (repository *ContactMock) Update(id int, req model.ContactRequest) (model.Contact, error) {
	var test model.Contact = model.Contact{
		Id:     1,
		Name:   "asdas",
		NoTelp: "asdjha",
	}
	return test, nil
}

func (repository *ContactMock) Delete(id int) (int, error) {

	return id, nil

}
