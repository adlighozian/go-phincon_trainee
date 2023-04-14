package repository

import (
	"contact-go/model"
)

type ContactRepository interface {
	List() ([]model.Contact, error)
	Add(req []model.ContactRequest) ([]model.Contact, error)
	Update(id int, req model.ContactRequest) (model.Contact, error)
	Delete(id int) (int, error)
}
