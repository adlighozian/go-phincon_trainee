package usecase

import (
	"contact-go/model"
	"contact-go/repository"
	"errors"
	"log"
	"net/http"
)

type contactUseCase struct {
	Repository repository.ContactRepository
}

func NewContactUseCase(repository repository.ContactRepository) ContactUseCase {
	return &contactUseCase{
		Repository: repository,
	}
}

const sresult string = "Internal Database Error"

func (usecase *contactUseCase) List() (model.ContactResponse, error) {
	log.Println("list usecase")
	collection_contact, err := usecase.Repository.List()
	if err != nil {
		return model.ContactResponse{
			Status:  http.StatusBadGateway,
			Message: sresult,
			Data:    nil,
		}, err
	} else {
		return model.ContactResponse{
			Status:  http.StatusOK,
			Message: "oke",
			Data:    collection_contact,
		}, nil
	}

}

func (usecase *contactUseCase) Add(req []model.ContactRequest) (model.ContactResponse, error) {
	log.Println("add usecase")
	if req == nil {
		return model.ContactResponse{
			Status:  http.StatusBadRequest,
			Message: "Status Bad Request",
			Data:    nil,
		}, errors.New("data tidak ada")
	}
	collection_contact, err := usecase.Repository.Add(req)
	if err != nil {
		return model.ContactResponse{
			Status:  http.StatusBadGateway,
			Message: sresult,
			Data:    nil,
		}, err
	}
	return model.ContactResponse{
		Status:  http.StatusCreated,
		Message: "oke",
		Data:    collection_contact,
	}, nil
}

func (usecase *contactUseCase) Update(id int, req model.ContactRequest) (model.ContactResponse, error) {

	err := usecase.Repository.Update(id, req)
	if err != nil {
		return model.ContactResponse{
			Status:  http.StatusBadGateway,
			Message: sresult,
			Data:    nil,
		}, err
	}
	return model.ContactResponse{
		Status:  http.StatusOK,
		Message: "oke",
		Data:    nil,
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

	err := usecase.Repository.Delete(id)
	if err != nil {
		return model.ContactResponse{
			Status:  http.StatusBadGateway,
			Message: sresult,
			Data:    nil,
		}, err
	}
	return model.ContactResponse{
		Status:  http.StatusOK,
		Message: "oke",
		Data:    nil,
	}, nil
}
