package controller

import (
	"encoding/json"
	"fmt"
	"inventory/model"
	"inventory/service"
	"net/http"
)

type handlerHttp struct {
	productService  service.ProductService
	purchaseService service.PurchaseService
	// PurchaseRepository repository.PurchaseOrderRepository
	// SalesRepository    repository.SalesOrderRepository
}

func NewHandlerHttp(product service.ProductService, purchase service.PurchaseService) InventoryHandlerHttp {
	return &handlerHttp{
		productService:  product,
		purchaseService: purchase,
	}
}

func (handler *handlerHttp) ProductShow(w http.ResponseWriter, r *http.Request) {
	product, _ := handler.productService.ShowProduct()
	result, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	w.WriteHeader(product.Status)
	w.Write(result)
}

func (handler *handlerHttp) PurchaseInput(w http.ResponseWriter, r *http.Request) {
	req := []model.ReqPurchase{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("kesalahan bad request"))
		fmt.Printf("error")
	}

	var purchase []model.ReqPurchase
	for _, v := range req {

		input := model.ReqPurchase{
			Item:  v.Item,
			Price: v.Price,
			From:  v.From,
			Total: v.Total,
		}
		purchase = append(purchase, input)

	}

	data, _ := handler.purchaseService.InputPurchase(purchase)
	result, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	w.WriteHeader(data.Status)
	w.Write(result)

}

func (handler *handlerHttp) PurchaseDetail(w http.ResponseWriter, r *http.Request) {
	fmt.Println("controller")
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("kesalahan bad request"))
		fmt.Printf("error")
	}
	sli := json.NewDecoder(r.Body)
	var respon = make(map[string]interface{})
	err = sli.Decode(&respon)
	if err != nil {
		panic(err)
	}

	order := respon["order"].(string)

	data, _ := handler.purchaseService.DetailPurchase(order)
	result, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	w.WriteHeader(data.Status)
	w.Write(result)

}
