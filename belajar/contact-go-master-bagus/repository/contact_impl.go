package repository

import (
	"contact-go/helper"
	"contact-go/model"
)

type contactRepository struct{}

func NewContactRepository() ContactRepository {
	return new(contactRepository)
}

func (repo *contactRepository) List() ([]model.Contact, error) {
	return model.Contacts, nil
}

func (repo *contactRepository) getLastID() (int64, error) {
	contacts, err := repo.List()

	var tempID int64
	for _, v := range contacts {
		if tempID < v.ID {
			tempID = v.ID
		}
	}
	return tempID, err
}

func (repo *contactRepository) getIndexByID(id int64) (int, error) {
	contacts, err := repo.List()
	if err != nil {
		return -1, err
	}

	for i, v := range contacts {
		if id == v.ID {
			return i, nil
		}
	}

	return -1, helper.NewAppError(helper.ErrContactNotFound)
}

func (repo *contactRepository) Add(contact *model.Contact) (*model.Contact, error) {
	id, err := repo.getLastID()
	if err != nil {
		return nil, err
	}

	newContact := contact
	newContact.ID = id + 1

	model.Contacts = append(model.Contacts, *newContact)

	return newContact, nil
}

func (repo *contactRepository) Detail(id int64) (*model.Contact, error) {
	contacts, err := repo.List()
	if err != nil {
		return nil, err
	}

	index, err := repo.getIndexByID(id)
	if err != nil {
		return nil, err
	}

	contact := contacts[index]

	return &contact, nil
}

func (repo *contactRepository) Update(id int64, contact *model.Contact) (*model.Contact, error) {
	contacts, err := repo.List()
	if err != nil {
		return nil, err
	}

	index, err := repo.getIndexByID(id)
	if err != nil {
		return nil, err
	}

	updatedContact := &contacts[index]
	updatedContact.Name = contact.Name
	updatedContact.NoTelp = contact.NoTelp

	return updatedContact, nil
}

func (repo *contactRepository) Delete(id int64) error {
	index, err := repo.getIndexByID(id)
	if err != nil {
		return err
	}

	model.Contacts = append(model.Contacts[:index], model.Contacts[index+1:]...)

	return nil
}
