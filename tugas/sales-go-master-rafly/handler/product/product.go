package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sales-go/model"
	"sales-go/repository/product"
)

type handler struct {
	repo product.Repositorier
}

func NewHandler(repositorier product.Repositorier) *handler {
	return &handler{
		repo: repositorier,
	}
}

type Handlerer interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

func (handler *handler) GetList(w http.ResponseWriter, r *http.Request) {
	result, err := handler.repo.GetList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
		log.Println("[ERROR] get list product :", err.Error())
		return
	}
	
	jsonData, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
		log.Println("[ERROR] marshal list product :", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (handler *handler) Create(w http.ResponseWriter, r *http.Request) {
	req := model.ProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
		log.Println("[ERROR] decode request :", err.Error())
		return
	}

	_, err = handler.repo.GetProductByName(req.Name)
	if err != nil {
		if req.Price <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("message : price must be > 0"))
			return
		}

		result, err := handler.repo.Create(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
			log.Println("[ERROR] create product :", err.Error())
			return
		} else if err == nil {
			jsonData, err := json.Marshal(result)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
				log.Println("[ERROR] marshal result create product :", err.Error())
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write(jsonData)
			return
		}
	} else {
		log.Print(err.Error())
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("message : product already exist, pelase input another product name."))
		return
	}
}
