package handler

import (
	"contact-go/model"
	"contact-go/usecase"
	"encoding/json"
	"log"
	"net/http"
)

type contactHttpDbHandler struct {
	usecase usecase.UseCaseInterface
}

func NewContactHttpDbHandler(usecase usecase.UseCaseInterface) *contactHttpDbHandler{
	return &contactHttpDbHandler{
		usecase: usecase,
	}
}

type ContactHttpDbHandlerInterface interface {
	List(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func (handler *contactHttpDbHandler) List(w http.ResponseWriter, r *http.Request) {
	response, err := handler.usecase.List()
	if err != nil {
		log.Print(err)
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Print(err)
	}
	w.WriteHeader(response.Status)
	w.Write(jsonData)
}

func (handler *contactHttpDbHandler) Add(w http.ResponseWriter, r *http.Request) {
	// using json
	req := []model.ContactRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// using post form
	/*
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
	}

	name := r.PostForm.Get("name")
	noTelp := r.PostForm.Get("no_telp")

	if name == "" || noTelp == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := model.ContactRequest{
		Name:   name,
		NoTelp: noTelp,
	}*/

	response, err := handler.usecase.Add(req)
	if err != nil {
		log.Print(err)
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Print(err)
	}
	w.WriteHeader(response.Status)
	w.Write(jsonData)
}

func (handler *contactHttpDbHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	// using json
	req := model.ContactRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// using post form
	/*
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
	}

	name := r.PostForm.Get("name")
	noTelp := r.PostForm.Get("no_telp")

	if name == "" || noTelp == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req = model.ContactRequest{
		Name:   name,
		NoTelp: noTelp,
	}*/

	response, err := handler.usecase.Update(idStr, req)
	if err != nil {
		log.Print(err)
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Print(err)
	}

	w.WriteHeader(response.Status)
	w.Write(jsonData)
}

func (handler *contactHttpDbHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	log.Print(idStr)
	response, err := handler.usecase.Delete(idStr)
	if err != nil {
		log.Print(err)
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Print(err)
	}

	w.WriteHeader(response.Status)
	w.Write(jsonData)
}