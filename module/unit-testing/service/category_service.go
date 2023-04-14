package service

import (
	"errors"
	"unit-testing/model"
	"unit-testing/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service *CategoryService) Get(id string) (*model.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}

func (service *CategoryService) GetBarang(id string) bool {
	category := service.Repository.FindBarang(id)
	return category
}
