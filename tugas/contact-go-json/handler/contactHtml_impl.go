package handler

import (
	"contact-go/model"
	"contact-go/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

type contactHandlerHttp struct{}

func NewContactHandlerHttp() ContactHandlerHttp {
	return new(contactHandlerHttp)
}

// Http GET
func (repo *contactHandlerHttp) HandlerGet(write http.ResponseWriter, request *http.Request) {
	write.WriteHeader(http.StatusOK)
	fmt.Println("Success", http.StatusOK)

	contacts := repository.NewContactRepository().List()
	result, err := json.Marshal(contacts)

	if err != nil {
		panic(err)
	}
	write.WriteHeader(200)
	write.Write(result)
}

// Http POST
func (repo *contactHandlerHttp) HandlerPost(write http.ResponseWriter, request *http.Request) {
	id := repository.NewContactRepository().GetLastID()

	err := request.ParseForm()
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("kesalahan bad request"))
	}

	name := request.PostForm.Get("name")
	telp := request.PostForm.Get("telp")

	contact := model.Contact{
		ID:     id + 1,
		Name:   name,
		NoTelp: telp,
	}

	model.Contacts = append(repository.NewContactRepository().DecodeJson(), contact)
	repository.NewContactRepository().EncodeJson()

	write.WriteHeader(http.StatusCreated)
	fmt.Println("Success", http.StatusCreated)
	fmt.Printf("name: %s dan telp: %s Berhasil\n", name, telp)
	fmt.Fprintf(write, "name: %s dan telp: %s Berhasil\n", name, telp)
}
