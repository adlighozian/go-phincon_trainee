package usecase

import (
	"contact-go/model"
	"contact-go/repository"
	"fmt"
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

func (usecase *contactUseCase) List() (model.ContactResponse, error) {
	collection_contact, err := usecase.Repository.List()
	if err != nil {
		fmt.Println("error")
		return model.ContactResponse{
			Status:  http.StatusBadGateway,
			Message: "Internal Database Error",
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
	collection_contact, err := usecase.Repository.Add(req)
	if err != nil {
		return model.ContactResponse{
			Status:  http.StatusBadGateway,
			Message: "Internal Database Error",
			Data:    nil,
		}, err
	}
	return model.ContactResponse{
		Status:  http.StatusOK,
		Message: "oke",
		Data:    collection_contact,
	}, nil
}

func (usecase *contactUseCase) Update(id int, req model.ContactRequest) (model.ContactResponse, error) {

	collection_contact, err := usecase.Repository.Update(id, req)
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

	collection_contact, err := usecase.Repository.Delete(id)
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
