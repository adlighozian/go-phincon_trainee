package repository

import (
	"contact-go/model"
)

type ContactRepository interface {
	List() []model.Contact
	Add(req model.ContactRequest) (model.Contact, error)
	Update(id int, req model.ContactRequest) (model.Contact, error)
	Delete(id int) error
	GetLastID() int
	DecodeJson() []model.Contact
	EncodeJson()
}
