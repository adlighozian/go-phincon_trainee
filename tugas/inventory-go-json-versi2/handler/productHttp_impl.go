package handler

import (
	"encoding/json"
	"fmt"
	"inventory/model"
	"net/http"
)

type productHandlerHttp struct{}

func NewProductHandlerHttp() ProductHandlerHttp {
	return new(productHandlerHttp)
}

func (repo *productHandlerHttp) ProductGet(write http.ResponseWriter, request *http.Request) {
	write.WriteHeader(http.StatusOK)
	fmt.Println("Success", http.StatusOK)

	contacts := model.PurchaseOrders
	result, err := json.Marshal(contacts)

	if err != nil {
		panic(err)
	}
	write.WriteHeader(200)
	write.Write(result)
}
