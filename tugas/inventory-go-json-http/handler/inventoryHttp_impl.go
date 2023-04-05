package handler

import (
	"encoding/json"
	"fmt"
	"inventory/model"
	"inventory/repository"
	"net/http"
)

type HandlerHttp struct {
	ProductRepository  repository.ProductRepository
	PurchaseRepository repository.PurchaseOrderRepository
	SalesRepository    repository.SalesOrderRepository
}

func NewHandlerHttp(product repository.ProductRepository, purchase repository.PurchaseOrderRepository, sales repository.SalesOrderRepository) InventoryHandlerHttp {
	return &HandlerHttp{
		ProductRepository:  product,
		PurchaseRepository: purchase,
		SalesRepository:    sales,
	}
}

func (handler *HandlerHttp) ProductGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("Success", http.StatusOK)

	contacts := handler.ProductRepository.DecodeProduct()
	result, err := json.Marshal(contacts)

	if err != nil {
		panic(err)
	}
	w.WriteHeader(200)
	w.Write(result)
}

func (handler *HandlerHttp) PurchaseGet(w http.ResponseWriter, r *http.Request) {
	// cek form
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("kesalahan bad request"))
		panic(err)
	}
	// membuat tampungan untuk body request
	data := json.NewDecoder(r.Body)
	var respon = make(map[string]interface{})
	err = data.Decode(&respon)
	if err != nil {
		panic(err)
	}

	// mengambil body request
	order := respon["order"].(string)

	showPurchaseDetail, err := handler.PurchaseRepository.ShowPurchaseOrderDetail(order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Order tidak ditemukan"))
	} else {
		result, err := json.Marshal(showPurchaseDetail)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}

}

func (handler *HandlerHttp) PurchasePost(w http.ResponseWriter, r *http.Request) {
	// cek form dan membuat tampungan untuk body request
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("kesalahan bad request"))
		panic(err)
	}
	data := json.NewDecoder(r.Body)
	var respon = make(map[string]interface{})
	err = data.Decode(&respon)
	if err != nil {
		panic(err)
	}

	item := respon["item"].(string)
	price := int(respon["price"].(float64))
	from := respon["from"].(string)
	total := int(respon["total"].(float64))

	inputReq := model.ReqPurchaseOrder{
		Item:  item,
		Price: price,
		From:  from,
		Total: total,
	}

	inputPurchase, err := handler.PurchaseRepository.InputPurchaseOrder(inputReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("kesalahan input"))
	} else {
		result, err := json.Marshal(inputPurchase)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}

}

func (handler *HandlerHttp) SalesGet(w http.ResponseWriter, r *http.Request) {
	// cek form
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("kesalahan bad request"))
		panic(err)
	}
	// membuat tampungan untuk body request
	data := json.NewDecoder(r.Body)
	var respon = make(map[string]interface{})
	err = data.Decode(&respon)
	if err != nil {
		panic(err)
	}

	// mengambil body request
	order := respon["order"].(string)

	showPurchaseDetail, err := handler.SalesRepository.ShowSalesOrderDetail(order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Order tidak ditemukan"))
	} else {
		result, err := json.Marshal(showPurchaseDetail)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}

}

func (handler *HandlerHttp) SalesPost(w http.ResponseWriter, r *http.Request) {
	// cek form dan membuat tampungan untuk body request
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("kesalahan bad request"))
		panic(err)
	}
	data := json.NewDecoder(r.Body)
	var respon = make(map[string]interface{})
	err = data.Decode(&respon)
	if err != nil {
		panic(err)
	}

	item := respon["item"].(string)
	price := int(respon["price"].(float64))
	from := respon["from"].(string)
	total := int(respon["total"].(float64))

	inputReq := model.ReqSalesOrder{
		Item:  item,
		Price: price,
		From:  from,
		Total: total,
	}

	inputPurchase, err := handler.SalesRepository.InputSalesOrder(inputReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Stok tidak cukup"))
	} else {
		result, err := json.Marshal(inputPurchase)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}

}
