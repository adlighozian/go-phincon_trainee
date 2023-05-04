package handler

import (
	"contact-go/helper"
	"contact-go/model"
	"contact-go/usecase"
	"fmt"
)

type contactHandler struct {
	ContactUC usecase.ContactUsecase
}

func NewContactHandler(contactUC usecase.ContactUsecase) ContactHandler {
	return &contactHandler{
		ContactUC: contactUC,
	}
}

func (handler *contactHandler) List() {
	err := helper.ClearTerminal()
	if err != nil {
		fmt.Println(err)
	}

	contacts, err := handler.ContactUC.List()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("|---------------|-----------------------|-----------------------|\n")
		fmt.Printf("| ID\t\t| Nama\t\t\t| No.Telp\t\t|\n")
		fmt.Printf("|---------------|-----------------------|-----------------------|\n")

		for _, v := range contacts {
			fmt.Printf("| %d\t\t| %s\t\t| %s\t\t|\n", v.ID, v.Name, v.NoTelp)
		}
		fmt.Printf("|---------------|-----------------------|-----------------------|\n")
	}
}

func (handler *contactHandler) Add() {
	err := helper.ClearTerminal()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Add new contact")

	fmt.Print("Name = ")
	var name string
	fmt.Scanln(&name)
	if name == "" {
		fmt.Println("Name yang dimasukkan tidak valid")
		return
	}

	fmt.Print("NoTelp = ")
	var noTelp string
	fmt.Scanln(&noTelp)
	if noTelp == "" {
		fmt.Println("NoTelp yang dimasukkan tidak valid")
		return
	}

	contactRequest := model.ContactRequest{
		Name:   name,
		NoTelp: noTelp,
	}

	contact, err := handler.ContactUC.Add(&contactRequest)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Berhasil add contact with id", contact.ID)
	}
}

func (handler *contactHandler) Detail() {
	err := helper.ClearTerminal()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Contact Detail")

	fmt.Print("Contact ID = ")
	var contactID int64
	fmt.Scanln(&contactID)
	if contactID == 0 {
		fmt.Println("Contact ID yang dimasukkan tidak valid")
		return
	}

	contact, err := handler.ContactUC.Detail(contactID)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("ID : \t\t%d\nNama : \t\t%s\nNo.Telp : \t%s\n", contact.ID, contact.Name, contact.NoTelp)

	}
}

func (handler *contactHandler) Update() {
	err := helper.ClearTerminal()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Update a contact")

	fmt.Print("ID = ")
	var id int64
	fmt.Scanln(&id)
	if id == 0 {
		fmt.Println("ID yang dimasukkan tidak valid")
		return
	}

	fmt.Print("Name = ")
	var name string
	fmt.Scanln(&name)
	if name == "" {
		fmt.Println("Name yang dimasukkan tidak valid")
		return
	}

	fmt.Print("NoTelp = ")
	var noTelp string
	fmt.Scanln(&noTelp)
	if noTelp == "" {
		fmt.Println("NoTelp yang dimasukkan tidak valid")
		return
	}

	contactRequest := model.ContactRequest{
		Name:   name,
		NoTelp: noTelp,
	}

	contact, err := handler.ContactUC.Update(id, &contactRequest)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Berhasil update contact with id", contact.ID)
	}
}

func (handler *contactHandler) Delete() {
	err := helper.ClearTerminal()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Delete a contact")

	fmt.Print("ID = ")
	var id int64
	fmt.Scanln(&id)
	if id == 0 {
		fmt.Println("ID yang dimasukkan tidak valid")
		return
	}

	err = handler.ContactUC.Delete(id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Berhasil delete contact with id", id)
	}
}
