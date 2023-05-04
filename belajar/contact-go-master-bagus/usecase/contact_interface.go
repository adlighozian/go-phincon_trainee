//go:generate mockery --output=../mocks --name ContactUsecase

package usecase

import "contact-go/model"

type ContactUsecase interface {
	List() ([]model.Contact, error)
	Add(req *model.ContactRequest) (*model.Contact, error)
	Detail(id int64) (*model.Contact, error)
	Update(id int64, req *model.ContactRequest) (*model.Contact, error)
	Delete(id int64) error
}
