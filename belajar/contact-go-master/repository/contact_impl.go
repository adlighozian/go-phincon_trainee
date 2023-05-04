package repository

import (
	"errors"

	model "contact-go/model"
)

type repository struct {}

func NewContactRepository () *repository {
	return &repository{}
}

func (repo *repository) getLastID() (lastID int, err error) {
	list, _ := repo.List()

	if len(list) == 0 {
		lastID = 0
	} else {
		for _, v := range list {
			if lastID < int(v.Id) {
				lastID = int(v.Id)
			}
		}
	}
	return
}

func (repo *repository) GetIndexById(id int) (index int, value model.Contact, err error) {
	for i, v := range model.ContactSlice {
		if v.Id == id {
			index = int(i)
			value = v
			return index, value, nil
		}
	}
	return -1, model.Contact{}, errors.New("id not found")
}

func (repo *repository) List() ([]model.Contact, error) {
	return model.ContactSlice, nil
}

func (repo *repository) Add(req []model.ContactRequest) (result []model.Contact, err error) {
	lastID, err := repo.getLastID()
	if err != nil {
		return
	}

	for i, v := range req {
		newContact := model.Contact{
			Id:     lastID+1+i,
			Name:  	v.Name,
			NoTelp: v.NoTelp,
		}
		model.ContactSlice = append(model.ContactSlice, newContact)
		result = append(result, newContact)

	}
	return
}

func (repo *repository) Update(id int, req model.ContactRequest) (err error) {
	index, value, err := repo.GetIndexById(id)
	if err != nil {
		return
	}

	if req.Name == "" {
		req.Name = value.Name
	}

	if req.NoTelp == "" {
		req.NoTelp = value.NoTelp
	}

	model.ContactSlice[index] = model.Contact{
		Id:     value.Id,
		Name:   req.Name,
		NoTelp: req.NoTelp,
	}

	return
}

func (repo *repository) Delete(id int) (err error) {
	index, _, err := repo.GetIndexById(id)
	if err != nil {
		return
	}

	deletedItemIndex := index
	model.ContactSlice = append(model.ContactSlice[:deletedItemIndex], model.ContactSlice[deletedItemIndex+1:]...)
	return
}