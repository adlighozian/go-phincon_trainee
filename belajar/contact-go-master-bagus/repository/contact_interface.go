//go:generate mockery --output=../mocks --name ContactRepository
package repository

import "contact-go/model"

type ContactRepository interface {
	List() ([]model.Contact, error)
	Add(contact *model.Contact) (*model.Contact, error)
	Detail(id int64) (*model.Contact, error)
	Update(id int64, contact *model.Contact) (*model.Contact, error)
	Delete(id int64) error
}
