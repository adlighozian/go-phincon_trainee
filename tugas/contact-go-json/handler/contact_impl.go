package handler

import (
	"contact-go/model"
	"contact-go/repository"
	"fmt"
)

type contactHandler struct {
	ContactRepository repository.ContactRepository
}

func NewcontactHandler(contactRepo repository.ContactRepository) ContactHandler {
	return &contactHandler{
		ContactRepository: contactRepo,
	}
}

func (handler *contactHandler) List() {
	fmt.Printf("|---------------|-----------------------|-----------------------|\n")
	fmt.Printf("|\tID\t|\tNama\t\t|\tNo.Telp\t\t|\n")
	fmt.Printf("|---------------|-----------------------|-----------------------|\n")

	contacts := handler.ContactRepository.List()
	for _, v := range contacts {
		fmt.Printf("|\t%d\t|\t%s\t\t|\t%s\t\t|\n", v.ID, v.Name, v.NoTelp)
	}
	fmt.Printf("|---------------|-----------------------|-----------------------|\n")
}

func (handler *contactHandler) Add() {
	fmt.Println("Add new contact")

	fmt.Print("Name = ")
	var name string
	fmt.Scanln(&name)

	fmt.Print("NoTelp = ")
	var noTelp string
	fmt.Scanln(&noTelp)

	contactRequest := model.ContactRequest{
		Name:   name,
		NoTelp: noTelp,
	}

	contact, err := handler.ContactRepository.Add(contactRequest)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Berhasil add contact with id", contact.ID)
	}
}

func (handler *contactHandler) Update() {
	fmt.Println("Update a contact")

	fmt.Print("ID = ")
	var id int64
	fmt.Scanln(&id)

	fmt.Print("Name = ")
	var name string
	fmt.Scanln(&name)

	fmt.Print("NoTelp = ")
	var noTelp string
	fmt.Scanln(&noTelp)

	contactRequest := model.ContactRequest{
		Name:   name,
		NoTelp: noTelp,
	}

	contact, err := handler.ContactRepository.Update(id, contactRequest)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Berhasil update contact with id", contact.ID)
	}
}

func (handler *contactHandler) Delete() {
	fmt.Println("Delete a contact")

	fmt.Print("ID = ")
	var id int64
	fmt.Scanln(&id)

	err := handler.ContactRepository.Delete(id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Berhasil delete contact with id", id)
	}
}
