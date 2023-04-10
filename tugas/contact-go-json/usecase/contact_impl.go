package usecase

import (
	"contact-go/model"
	"contact-go/repository"
	"errors"
	"net/http"
)

type contactUseCase struct{}

func NewContactUseCase() ContactUseCase {
	return new(contactUseCase)
}

func (usecase *contactUseCase) List() (model.ContactResponse, error) {
	collection_contact := repository.NewContactRepository().List()
	if collection_contact == nil {
		return model.ContactResponse{
			Status:  http.StatusBadGateway,
			Message: "Internal Database Error",
			Data:    nil,
		}, errors.New("internal Database Error")
	}
	return model.ContactResponse{
		Status:  http.StatusOK,
		Message: "oke",
		Data:    collection_contact,
	}, nil
}

func (usecase *contactUseCase) Add(req []model.ContactRequest) (model.ContactResponse, error) {
	collection_contact, _ := repository.NewContactRepository().Add(req)
	if collection_contact == nil {
		return model.ContactResponse{
			Status:  http.StatusBadGateway,
			Message: "Internal Database Error",
			Data:    nil,
		}, errors.New("internal Database Error")
	}
	return model.ContactResponse{
		Status:  http.StatusOK,
		Message: "oke",
		Data:    collection_contact,
	}, nil
}

func (usecase *contactUseCase) Update(id int, req model.ContactRequest) (model.ContactResponse, error) {

	collection_contact, err := repository.NewContactRepository().Update(id, req)
	if err != nil {
		return model.ContactResponse{
			Status:  http.StatusBadGateway,
			Message: "Internal Database Error",
			Data:    nil,
		}, err
	}
	return model.ContactResponse{
		Status:  http.StatusOK,
		Message: "Berhasil diupdate",
		Data:    collection_contact,
	}, nil
}

func (usecase *contactUseCase) Delete(id int) (model.ContactResponse, error) {

	if id <= 0 {
		return model.ContactResponse{
			Status:  http.StatusNotFound,
			Message: "id tidak ditemukan",
			Data:    nil,
		}, nil
	}

	collection_contact, err := repository.NewContactRepository().Delete(id)
	if err != nil {
		return model.ContactResponse{
			Status:  http.StatusBadGateway,
			Message: "Internal Database Error",
			Data:    nil,
		}, err
	}
	return model.ContactResponse{
		Status:  http.StatusOK,
		Message: "Berhasil Dihapus",
		Data:    collection_contact,
	}, nil
}
