package voucher

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sales-go/model"
	"sales-go/repository/voucher"
)

type jsonhttphandler struct {
	repo voucher.Repositorier
}

func NewJsonHTTPHandler(repositorier voucher.Repositorier) *jsonhttphandler {
	return &jsonhttphandler{
		repo: repositorier,
	}
}

func (handler *jsonhttphandler) GetList(w http.ResponseWriter, r *http.Request) {
	result, err := handler.repo.GetList()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
		log.Println("[ERROR] get list voucher :", err.Error())
		return
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
		log.Println("[ERROR] marshal list voucher :", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (handler *jsonhttphandler) Create(w http.ResponseWriter, r *http.Request) {	
	req := []model.VoucherRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
		log.Println("[ERROR] decode request :", err.Error())
		return
	}

	for _, req := range req {
		_, err = handler.repo.GetVoucherByCode(req.Code)
		if err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("message : voucher already exist, pelase input another product name."))
			return
		}
		if req.Persen <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("message : persen must be > 0"))
			return
		}
	}

	result, err := handler.repo.Create(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
		log.Println("[ERROR] create voucher :", err.Error())
		return
	} else {
		jsonData, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
			log.Println("[ERROR] marshal result create voucher :", err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonData)
		return
	}
}