package repository

import (
	"unit-testing/model"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *model.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(model.Category)
		return &category
	}
}

func (repository *CategoryRepositoryMock) FindBarang(id string) bool {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == false {
		return false
	} else {
		return true
	}
}

// func (repository *CategoryRepositoryMock) Find(id string) bool {
// 	arguments := repository.Mock.Called(id)
// 	if arguments.Get(0) == false {
// 		return false
// 	} else {
// 		return true
// 	}
// }
