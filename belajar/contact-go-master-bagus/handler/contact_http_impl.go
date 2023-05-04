package handler

import (
	"contact-go/helper"
	"contact-go/model"
	"contact-go/usecase"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type contactHTTPHandler struct {
	ContactUC usecase.ContactUsecase
}

func NewContactHTTPHandler(contactUC usecase.ContactUsecase) ContactHTTPHandler {
	return &contactHTTPHandler{
		ContactUC: contactUC,
	}
}

func (handler *contactHTTPHandler) List(w http.ResponseWriter, r *http.Request) {
	contacts, err := handler.ContactUC.List()
	if err != nil {
		_ = helper.NewJsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = helper.NewJsonResponse(w, http.StatusOK, "OK", contacts)
	if err != nil {
		_ = helper.NewJsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

func (handler *contactHTTPHandler) Add(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// name := r.PostForm.Get("name")
	// if name == "" {
	// 	http.Error(w, "name yang dimasukkan tidak valid", http.StatusBadRequest)
	// 	return
	// }

	// noTelp := r.PostForm.Get("no_telp")
	// if noTelp == "" {
	// 	http.Error(w, "no_telp yang dimasukkan tidak valid", http.StatusBadRequest)
	// 	return
	// }

	var contactRequest model.ContactRequest
	err := json.NewDecoder(r.Body).Decode(&contactRequest)
	if err != nil {
		_ = helper.NewJsonResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if contactRequest.Name == "" {
		_ = helper.NewJsonResponse(w, http.StatusBadRequest, helper.ErrContactNameNotValid, nil)
		return
	}

	if contactRequest.NoTelp == "" {
		_ = helper.NewJsonResponse(w, http.StatusBadRequest, helper.ErrContactNoTelpNotValid, nil)
		return
	}

	contact, err := handler.ContactUC.Add(&contactRequest)
	if err != nil {
		code, message := helper.HandleAppError(err)
		_ = helper.NewJsonResponse(w, code, message, nil)
		return
	}

	// msg := fmt.Sprintf(`{ "message":"Berhasil add contact with id %d" }`, contact.ID)
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(msg))

	err = helper.NewJsonResponse(w, http.StatusCreated, "Created", contact)
	if err != nil {
		_ = helper.NewJsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

func (handler *contactHTTPHandler) Detail(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/contacts/")
	if idStr == "" {
		_ = helper.NewJsonResponse(w, http.StatusBadRequest, helper.ErrContactIdNotValid, nil)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		_ = helper.NewJsonResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	contact, err := handler.ContactUC.Detail(int64(id))
	if err != nil {
		code, message := helper.HandleAppError(err)
		_ = helper.NewJsonResponse(w, code, message, nil)
		return
	}

	err = helper.NewJsonResponse(w, http.StatusOK, "OK", contact)
	if err != nil {
		_ = helper.NewJsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

func (handler *contactHTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/contacts/")
	if idStr == "" {
		_ = helper.NewJsonResponse(w, http.StatusBadRequest, helper.ErrContactIdNotValid, nil)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		_ = helper.NewJsonResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var contactRequest model.ContactRequest
	err = json.NewDecoder(r.Body).Decode(&contactRequest)
	if err != nil {
		_ = helper.NewJsonResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	contact, err := handler.ContactUC.Update(int64(id), &contactRequest)
	if err != nil {
		code, message := helper.HandleAppError(err)
		_ = helper.NewJsonResponse(w, code, message, nil)
		return
	}

	err = helper.NewJsonResponse(w, http.StatusOK, "OK", contact)
	if err != nil {
		_ = helper.NewJsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

func (handler *contactHTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/contacts/")
	if idStr == "" {
		_ = helper.NewJsonResponse(w, http.StatusBadRequest, helper.ErrContactIdNotValid, nil)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		_ = helper.NewJsonResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = handler.ContactUC.Delete(int64(id))
	if err != nil {
		code, message := helper.HandleAppError(err)
		_ = helper.NewJsonResponse(w, code, message, nil)
		return
	}

	err = helper.NewJsonResponse(w, http.StatusOK, "OK", nil)
	if err != nil {
		_ = helper.NewJsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}
