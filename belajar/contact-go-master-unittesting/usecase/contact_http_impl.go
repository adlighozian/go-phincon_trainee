package usecase

import (
	"contact-go/model"
	"contact-go/repository"
	"errors"
	"net/http"
	"strconv"
)

type usecase struct {
	repo repository.ContactRepositorier
}

func NewUseCase(repository repository.ContactRepositorier) *usecase {
	return &usecase{
		repo: repository,
	}
}

func (uc *usecase) IsValidID(idStr string) (int, model.ContactResponse, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, model.ContactResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}, err
	}
	if id <= 0 {
		return 0, model.ContactResponse{
			Status:  http.StatusBadRequest,
			Message: "Id should not be more than 0",
			Data:    nil,
		}, errors.New("id should not be more than 0")
	}
	return id, model.ContactResponse{}, nil
}

func (uc *usecase) IsValidNameAndNoTelp(name string, noTelp string) (model.ContactResponse, error) {
	if name == "" {
		return model.ContactResponse{
			Status:  http.StatusBadRequest,
			Message: "Name should not be empty",
			Data:    nil,
		}, errors.New("name should not be empty")
	}

	if noTelp == "" || noTelp == "0" {
		return model.ContactResponse{
			Status:  http.StatusBadRequest,
			Message: "No telp should not be empty",
			Data:    nil,
		}, errors.New("no telp should not be empty")
	}
	return model.ContactResponse{}, nil
}

func (uc *usecase) List() (model.ContactResponse, error) {
	res, err := uc.repo.List()
	if err != nil {
		return model.ContactResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		}, err
	}
	return model.ContactResponse{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    res,
	}, nil
}

func (uc *usecase) Add(req []model.ContactRequest) (model.ContactResponse, error) {
	for _, v := range req {
		res, err := uc.IsValidNameAndNoTelp(v.Name, v.NoTelp)
		if err != nil {
			return res, err
		}
	}

	res, err := uc.repo.Add(req)
	if err != nil {
		return model.ContactResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		}, err
	}
	return model.ContactResponse{
		Status:  http.StatusCreated,
		Message: "Created",
		Data:    res,
	}, nil
}

func (uc *usecase) Update(idStr string, req model.ContactRequest) (model.ContactResponse, error) {
	id, res, err := uc.IsValidID(idStr)
	if err != nil {
		return res, err
	}

	res, err = uc.IsValidNameAndNoTelp(req.Name, req.NoTelp)
	if err != nil {
		return res, err
	}

	err = uc.repo.Update(id, req)
	if err != nil {
		return model.ContactResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		}, err
	}
	return model.ContactResponse{
		Status:  http.StatusOK,
		Message: "Updated",
		Data:    nil,
	}, nil
}

func (uc *usecase) Delete(idStr string) (model.ContactResponse, error) {
	id, res, err := uc.IsValidID(idStr)
	if err != nil {
		return res, err
	}

	err = uc.repo.Delete(id)
	if err != nil {
		return model.ContactResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		}, err
	}
	return model.ContactResponse{
		Status:  http.StatusOK,
		Message: "Deleted",
		Data:    nil,
	}, nil
}
