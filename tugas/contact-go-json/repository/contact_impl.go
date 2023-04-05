package repository

import (
	"contact-go/model"
	"encoding/json"
	"errors"
	"os"
)

type contactRepository struct{}

func NewContactRepository() ContactRepository {
	return new(contactRepository)
}

func (repo *contactRepository) DecodeJson() []model.Contact {
	reader, err := os.Open("./assets/contacts.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(reader)
	decoder.Decode(&model.Contacts)
	return model.Contacts
}

func (repo *contactRepository) EncodeJson() {
	writer, err := os.Create("./assets/contacts.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(writer)
	encoder.Encode(repo.DecodeJson())
}

func (repo *contactRepository) List() []model.Contact {
	return repo.DecodeJson()
}

func (repo *contactRepository) GetLastID() int {
	contacts := repo.List()

	var tempID int
	for _, v := range contacts {
		if tempID < v.Id {
			tempID = v.Id
		}
	}
	return tempID
}

func (repo *contactRepository) GetIndexByID(id int) (int, error) {
	contacts := repo.List()

	for i, v := range contacts {
		if id == v.Id {
			return i, nil
		}
	}

	return -1, errors.New("ID tidak ditemukan")
}

func (repo *contactRepository) Add(req model.ContactRequest) (model.Contact, error) {
	id := repo.GetLastID()

	contact := model.Contact{
		Id:     id + 1,
		Name:   req.Name,
		NoTelp: req.NoTelp,
	}

	model.Contacts = append(repo.DecodeJson(), contact)
	repo.EncodeJson()

	return contact, nil
}

func (repo *contactRepository) Update(id int, req model.ContactRequest) (model.Contact, error) {
	contacts := repo.List()
	index, err := repo.GetIndexByID(id)

	if err != nil {
		return model.Contact{}, err
	}

	contact := &contacts[index]
	contact.Name = req.Name
	contact.NoTelp = req.NoTelp

	repo.EncodeJson()

	return *contact, nil
}

func (repo *contactRepository) Delete(id int) error {
	index, err := repo.GetIndexByID(id)

	if err != nil {
		return err
	}

	model.Contacts = append(model.Contacts[:index], model.Contacts[index+1:]...)
	repo.EncodeJson()

	return nil
}
