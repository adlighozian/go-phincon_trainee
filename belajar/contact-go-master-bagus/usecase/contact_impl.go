package usecase

import (
	"contact-go/model"
	"contact-go/repository"
)

type contactUsecase struct {
	ContactRepo repository.ContactRepository
}

func NewContactUsecase(contactRepo repository.ContactRepository) ContactUsecase {
	return &contactUsecase{
		ContactRepo: contactRepo,
	}
}

func (uc *contactUsecase) List() ([]model.Contact, error) {
	return uc.ContactRepo.List()
}

func (uc *contactUsecase) Add(req *model.ContactRequest) (*model.Contact, error) {
	contact := model.Contact{
		Name:   req.Name,
		NoTelp: req.NoTelp,
	}
	return uc.ContactRepo.Add(&contact)
}

func (uc *contactUsecase) Detail(id int64) (*model.Contact, error) {
	return uc.ContactRepo.Detail(id)
}

func (uc *contactUsecase) Update(id int64, req *model.ContactRequest) (*model.Contact, error) {
	contact := model.Contact{
		Name:   req.Name,
		NoTelp: req.NoTelp,
	}
	return uc.ContactRepo.Update(id, &contact)
}

func (uc *contactUsecase) Delete(id int64) error {
	return uc.ContactRepo.Delete(id)
}
