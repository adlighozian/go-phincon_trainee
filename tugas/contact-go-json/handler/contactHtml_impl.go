package handler

import (
	"contact-go/model"
	"contact-go/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

	encoder_ := json.NewDecoder(request.Body)

	var respon = make(map[string]interface{})

	err = encoder_.Decode(&respon)

	if err != nil {
		panic(err)
	}

	// fmt.Printf(respon["name"].(string))

	name := respon["name"].(string)

	telp := respon["telp"].(string)

	fmt.Printf("name: %s ", name)

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

func (repo *contactHandlerHttp) HandlerUpdate(write http.ResponseWriter, request *http.Request) {

	err := request.ParseForm()

	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("kesalahan bad request"))
		fmt.Printf("error")
	}

	encoder_ := json.NewDecoder(request.Body)

	var respon = make(map[string]interface{})

	err = encoder_.Decode(&respon)

	if err != nil {
		panic(err)
	}

	fmt.Printf(respon["id"].(string))
	fmt.Printf(respon["name"].(string))
	fmt.Printf(respon["notelp"].(string))

	id := respon["id"].(string)
	name := respon["name"].(string)
	telp := respon["notelp"].(string)

	a, _ := strconv.Atoi(id)

	var index int
	for i, v := range model.Contacts {
		if name == v.Name {
			index = i
		}
	}

	if index == 0 {
		fmt.Fprintf(write, "Nama tidak ditemukan")
	}

	contacts := model.Contacts
	contact := &contacts[a]
	contact.Name = name
	contact.NoTelp = telp

	repository.NewContactRepository().EncodeJson()

	write.WriteHeader(http.StatusCreated)
	fmt.Println("Success", http.StatusCreated)
	fmt.Printf("name: %s dan telp: %s Berhasil\n", name, telp)
	fmt.Fprintf(write, "name: %s dan telp: %s Berhasil\n", name, telp)
}

func (repo *contactHandlerHttp) HandlerDelete(write http.ResponseWriter, request *http.Request) {
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

	id := respon["id"].(string)

	a, _ := strconv.Atoi(id)

	val := a - 1
	var index int
	for i := range model.Contacts {
		if a == i {
			index = i
		}
	}

	if index == 0 {
		fmt.Fprintf(write, "Nama tidak ditemukan")
	}

	model.Contacts = append(model.Contacts[:val], model.Contacts[val+1:]...)
	repository.NewContactRepository().EncodeJson()

	write.WriteHeader(http.StatusCreated)
	fmt.Println("Success", http.StatusCreated)
	fmt.Fprintf(write, "Id berhasil dihapus")

}
