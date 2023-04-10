package handler

import (
	"contact-go/model"
	"contact-go/repository"
	"contact-go/usecase"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type contactHandlerHttp struct {
	contactRepository repository.ContactRepository
	contactUseCase    usecase.ContactUseCase
}

func NewContactHandlerHttp(repository repository.ContactRepository, usecase usecase.ContactUseCase) ContactHandlerHttp {
	return &contactHandlerHttp{
		contactRepository: repository,
		contactUseCase:    usecase,
	}
}

// Http list
func (handle *contactHandlerHttp) HandlerGet(write http.ResponseWriter, request *http.Request) {
	write.WriteHeader(http.StatusOK)
	fmt.Println("Success", http.StatusOK)

	contacts, _ := handle.contactUseCase.List()

	result, err := json.Marshal(contacts)
	if err != nil {
		write.WriteHeader(contacts.Status)
		write.Write(nil)
		return
	}

	write.WriteHeader(contacts.Status)
	write.Write(result)
}

// Http add
func (handle *contactHandlerHttp) HandlerPost(write http.ResponseWriter, request *http.Request) {
	req := []model.ContactRequest{}
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		write.WriteHeader(http.StatusInternalServerError)
		write.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
		log.Println("[ERROR] decode request :", err.Error())
		return
	}

	var slices []model.ContactRequest
	for _, v := range req {

		if v.Name == "" || v.NoTelp == "" {
			continue
		}

		inputReq := model.ContactRequest{
			Name:   v.Name,
			NoTelp: v.NoTelp,
		}
		slices = append(slices, inputReq)
	}

	inputPurchase, err := handle.contactUseCase.Add(slices)
	if err != nil {
		write.WriteHeader(http.StatusInternalServerError)
		write.Write([]byte("kesalahan input"))
		return
	} else {
		result, err := json.Marshal(inputPurchase)
		if err != nil {
			panic(err)
		}
		write.WriteHeader(http.StatusCreated)
		write.Write(result)
	}

}

// http update
func (handle *contactHandlerHttp) HandlerUpdate(write http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("kesalahan bad request"))
		fmt.Printf("error")
	}
	data := json.NewDecoder(request.Body)
	var respon = make(map[string]interface{})
	err = data.Decode(&respon)
	if err != nil {
		panic(err)
	}

	id := int(respon["id"].(float64))
	name := respon["name"].(string)
	telp := respon["notelp"].(string)

	contact := model.ContactRequest{
		Name:   name,
		NoTelp: telp,
	}

	updateContact, err := handle.contactUseCase.Update(id, contact)
	if err != nil {
		write.WriteHeader(updateContact.Status)
		write.Write([]byte(updateContact.Message))
		return
	} else {
		result, err := json.Marshal(updateContact)
		if err != nil {
			panic(err)
		}
		write.WriteHeader(updateContact.Status)
		write.Write(result)
	}
}

// http delete
func (handle *contactHandlerHttp) HandlerDelete(write http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("kesalahan bad request"))
	}
	encoders := json.NewDecoder(request.Body)
	var respon = make(map[string]interface{})
	err = encoders.Decode(&respon)
	if err != nil {
		panic(err)
	}

	id := int(respon["id"].(float64))

	updateContact, err := handle.contactUseCase.Delete(id)
	if err != nil {
		write.WriteHeader(updateContact.Status)
		write.Write([]byte(updateContact.Message))
		return
	} else {
		result, err := json.Marshal(updateContact)
		if err != nil {
			panic(err)
		}
		write.WriteHeader(updateContact.Status)
		write.Write(result)
	}

}
