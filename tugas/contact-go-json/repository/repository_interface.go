package repository

import (
	"contact-go/model"
)

type ContactRepository interface {
	List() ([]model.Client, error)
	Add(req []model.ContactRequest) ([]model.Client, error)
	Update(id int, req model.ContactRequest) error
	Delete(id int) error
}
