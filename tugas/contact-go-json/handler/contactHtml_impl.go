package handler

import (
	"contact-go/model"
	"contact-go/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

type contactHandlerHttp struct {
	contactRepository repository.ContactRepository
}

func NewContactHandlerHttp(contact repository.ContactRepository) ContactHandlerHttp {
	return &contactHandlerHttp{
		contactRepository: contact,
	}
}

// Http list
func (handle *contactHandlerHttp) HandlerGet(write http.ResponseWriter, request *http.Request) {
	write.WriteHeader(http.StatusOK)
	fmt.Println("Success", http.StatusOK)

	contacts := handle.contactRepository.List()
	result, err := json.Marshal(contacts)

	if err != nil {
		panic(err)
	}
	write.WriteHeader(200)
	write.Write(result)
}

// Http add
func (handle *contactHandlerHttp) HandlerPost(write http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("kesalahan bad request"))
	}
	encoder_ := json.NewDecoder(request.Body)
	var respon = make(map[string]interface{})
	err = encoder_.Decode(&respon)
	if err != nil {
		panic(err)
	}

	name := respon["name"].(string)
	telp := respon["telp"].(string)

	contact := model.ContactRequest{
		Name:   name,
		NoTelp: telp,
	}

	data, _ := handle.contactRepository.Add(contact)

	result, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	write.WriteHeader(http.StatusCreated)
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

	handle.contactRepository.Update(id, contact)

	write.WriteHeader(http.StatusCreated)
	fmt.Println("Success", http.StatusCreated)
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

	handle.contactRepository.Delete(id)

	write.WriteHeader(http.StatusCreated)
	fmt.Println("Success", http.StatusCreated)
	fmt.Fprintf(write, "Id berhasil dihapus")

}
