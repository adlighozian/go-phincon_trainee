package repository

import (
	"contact-go/model"
)

type ContactRepository interface {
	List() []model.Contact
	Add(req model.ContactRequest) (model.Contact, error)
	Update(id int64, req model.ContactRequest) (model.Contact, error)
	Delete(id int64) error
	GetLastID() int64
	DecodeJson() []model.Contact
	EncodeJson()
}
