package usecase

import "contact-go/model"

type ContactUseCase interface {
	List() (model.ContactResponse, error)
	Add(req []model.ContactRequest) (model.ContactResponse, error)
	Update(id int, req model.ContactRequest) (model.ContactResponse, error)
	Delete(id int) (model.ContactResponse, error)
}
