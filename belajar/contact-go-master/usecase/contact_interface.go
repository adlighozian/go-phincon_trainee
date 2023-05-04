package usecase

import "contact-go/model"

type UseCaseInterface interface {
	List() (model.ContactResponse, error)
	Add(req []model.ContactRequest) (model.ContactResponse, error)
	Update(idStr string, req model.ContactRequest) (model.ContactResponse, error)
	Delete(idStr string) (model.ContactResponse, error)
}