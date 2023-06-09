package handler

import (
	"contact-go/model"
	"contact-go/usecase"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type contactHandlerHttp struct {
	// contactRepository repository.ContactRepository
	contactUseCase usecase.ContactUseCase
}

func NewContactHandlerHttp(usecase usecase.ContactUseCase) *contactHandlerHttp {
	return &contactHandlerHttp{
		contactUseCase: usecase,
	}
}

// Http list
func (handle *contactHandlerHttp) HandlerGet(write http.ResponseWriter, request *http.Request) {
	log.Println("list handler")

	contacts, err := handle.contactUseCase.List()
	if err != nil {
		log.Println(err)
	}

	result, err := json.Marshal(contacts)
	if err != nil {
		log.Println(err)
	}

	write.WriteHeader(contacts.Status)
	write.Write(result)

}

// Http add
func (handle *contactHandlerHttp) HandlerPost(write http.ResponseWriter, request *http.Request) {
	log.Println("add handler")
	req := []model.ContactRequest{}
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		write.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
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

	contacts, err := handle.contactUseCase.Add(slices)
	if err != nil {
		log.Println(err)
	}

	result, err := json.Marshal(contacts)
	if err != nil {
		log.Println(err)
	}

	write.WriteHeader(contacts.Status)
	write.Write(result)

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
	} else {
		result, err := json.Marshal(updateContact)
		if err != nil {
			panic(err)
		}
		write.WriteHeader(updateContact.Status)
		write.Write(result)
	}

}
