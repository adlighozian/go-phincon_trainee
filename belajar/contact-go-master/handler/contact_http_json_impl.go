package handler

import (
	"contact-go/model"
	"contact-go/repository"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type contactHttpJsonHandler struct {
	repo repository.ContactRepositorier
}

func NewContactHttpJsonHandler(contactrepository repository.ContactRepositorier) *contactHttpJsonHandler{
	return &contactHttpJsonHandler{
		repo: contactrepository,
	}
}

type ContactHttpJsonHandlerInterface interface {
	List(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func (handler *contactHttpJsonHandler) List(w http.ResponseWriter, r *http.Request) {
	contact, err := handler.repo.List()
	if err != nil {
		log.Print(err)
	}

	jsonData, err := json.Marshal(contact)
	if err != nil {
		log.Print(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (handler *contactHttpJsonHandler) Add(w http.ResponseWriter, r *http.Request) {
	// using json
	req := []model.ContactRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, v := range req {
		if v.Name == "" || v.NoTelp == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
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

	result, err := handler.repo.Add(req)
	if err != nil {
		log.Print(err)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (handler *contactHttpJsonHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// using json
	req := model.ContactRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.NoTelp == "" {
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

	err = handler.repo.Update(id, req)
	if err != nil {
		log.Print(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *contactHttpJsonHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = handler.repo.Delete(id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}