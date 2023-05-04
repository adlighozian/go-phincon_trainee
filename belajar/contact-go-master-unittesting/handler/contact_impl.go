package handler

import (
	"contact-go/model"
	"contact-go/repository"
	"fmt"
	"log"
)

type contactHandler struct {
	repo repository.ContactRepositorier
}

func NewContactHandler(contactrepo repository.ContactRepositorier) *contactHandler {
	return &contactHandler{
		repo: contactrepo,
	}
}

func (handler *contactHandler) List() {
	fmt.Printf("\nID\t\t| Nama\t\t | No Telp \n")
	contacts, err := handler.repo.List()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range contacts {
		if len(v.Name) > 5 {
			fmt.Printf("%d \t\t| %s\t | %s\n", v.Id, v.Name, v.NoTelp)
		} else {
			fmt.Printf("%d \t\t| %s\t\t | %s\n", v.Id, v.Name, v.NoTelp)
		}
	}
}

func (handler *contactHandler) Add() {
	fmt.Println("Add new contact")
	
	fmt.Print("Name = ")
	var name string
	fmt.Scanln(&name)
	
	fmt.Print("No telp = ")
	var no_telp string
	fmt.Scanln(&no_telp)

	contactRequest := model.ContactRequest{
		Name:   name,
		NoTelp: no_telp,
	}
	contact, err := handler.repo.Add([]model.ContactRequest{contactRequest})
	if err != nil {
		log.Print("\nThere is an error : ", err)
	}
	
	for _, v := range contact {
		fmt.Println("New data added success with id : ", v.Id)
	}
}

func (handler *contactHandler) Update() {
	fmt.Println("Update existing contact")
	fmt.Print("Input id : ")

	var id int
	fmt.Scanln(&id)

	fmt.Print("Name = ")
	var name string
	fmt.Scanln(&name)

	fmt.Print("No telp = ")
	var no_telp string
	fmt.Scanln(&no_telp)

	req := model.ContactRequest{
		Name:   name,
		NoTelp: no_telp,
	}

	err := handler.repo.Update(id, req)
	if err != nil {
		log.Print("\nThere is an error : ", err)
	}
	
	fmt.Printf("Contact with id %d updated\n", id)
}

func (handler *contactHandler) Delete() {
	fmt.Println("Delete existing contact")
	fmt.Print("Input id : ")
	var id int
	fmt.Scanln(&id)

	err := handler.repo.Delete(id)
	if err != nil {
		log.Print("\nThere is an error : ", err)
	}

	fmt.Printf("Contact with id %d deleted\n", id)
}