package repository

import (
	"contact-go/model"
)

type ContactRepositorier interface {
	List() (result []model.Contact, err error)
	Add(req []model.ContactRequest) (result []model.Contact, err error)
	Update(id int, req model.ContactRequest) (err error)
	Delete(id int) (err error)
}