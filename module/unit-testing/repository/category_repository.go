package repository

import "unit-testing/model"

type CategoryRepository interface {
	FindById(id string) *model.Category
	FindBarang(id string) bool
	// Find(id string) bool
}
