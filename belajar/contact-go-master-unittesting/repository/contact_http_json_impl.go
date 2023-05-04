package repository

import (
	"encoding/json"
	"errors"
	"os"

	model "contact-go/model"
)

type contactjson struct {}

func NewContactJsonRepository() *contactjson {
	return &contactjson{}
}

func (repo *contactjson) getLastID() (lastID int, err error) {
	list, err := repo.List()
	if err != nil {
		return
	}

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

func (repo *contactjson) GetIndexById(id int) (index int, value model.Contact, err error) {
	list, err := repo.List()
	if err != nil {
		return
	}

	for i, v := range list {
		if v.Id == id {
			index = int(i)
			value = v
			return index, value, nil
		}
	}
	return -1, model.Contact{}, errors.New("id not found")
}

func (repo *contactjson) updateJSON(list []model.Contact) (err error) {
	// struct -> JSON
	write, err := os.Create("data/contact.txt")
	if err != nil {
		return
	}
	encoder := json.NewEncoder(write)
	encoder.Encode(list)
	return
}

func (repo *contactjson) List() (result []model.Contact, err error) {
	// JSON -> struct
	reader, err := os.Open("data/contact.txt")
	if err != nil {
		return
	}
	decoder := json.NewDecoder(reader)
	decoder.Decode(&result)

	return
}

func (repo *contactjson) Add(req []model.ContactRequest) (result []model.Contact, err error) {
	// JSON to struct
	list, err := repo.List()
	if err != nil {
		return
	}

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
		list = append(list, newContact)
		result = append(result, newContact)
	}

	err = repo.updateJSON(list)
	if err != nil {
		return
	}
	return
}

func (repo *contactjson) Update(id int, req model.ContactRequest) (err error) {
	index, value, err := repo.GetIndexById(id)
	if err != nil {
		return
	}

	list, err := repo.List()
	if err != nil {
		return
	}

	if req.Name == "" {
		req.Name = value.Name
	}

	if req.NoTelp == "" {
		req.NoTelp = value.NoTelp
	}

	list[index] = model.Contact{
		Id:     id,
		Name:   req.Name,
		NoTelp: req.NoTelp,
	}

	err = repo.updateJSON(list)
	if err != nil {
		return
	}
	return
}

func (repo *contactjson) Delete(id int) (err error) {
	list, err := repo.List()
	if err != nil {
		return
	}

	index, _, err := repo.GetIndexById(id)
	if err != nil {
		return
	}

	deletedItemIndex := index
	list = append(list[:deletedItemIndex], list[deletedItemIndex+1:]...)

	err = repo.updateJSON(list)
	if err != nil {
		return
	}
	return
}