package product

import (
	"fmt"
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

func (handler *handler) GetList(w http.ResponseWriter, r *http.Request) {
	result, err := handler.repo.GetList()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("\nId\t\t|Name\t\t\t|Price\t\t")
	for _, v := range result {
		if len(v.Name) > 13 {
			fmt.Printf("\n%d\t\t|%s\t|%0.2f", v.Id, v.Name, v.Price)
		} else if len(v.Name) > 5 && len(v.Name) < 13 {
			fmt.Printf("\n%d\t\t|%s\t\t|%0.2f", v.Id, v.Name, v.Price)
		} else {
			fmt.Printf("\n%d\t\t|%s\t\t\t|%0.2f", v.Id, v.Name, v.Price)
		}
	}
}

func (handler *handler) Create(w http.ResponseWriter, r *http.Request) {
	var name string
	var price float64
	fmt.Println("\nInput name data : ")
	fmt.Scanln(&name)
	fmt.Println("\nInput price data : ")
	fmt.Scanln(&price)

	data := model.ProductRequest{
		Name:  name,
		Price: price,
	}
	result, err := handler.repo.Create([]model.ProductRequest{data})
	if err != nil {
		fmt.Println(err)
		handler.Create(w, r)
	}

	fmt.Println("Product has been added : ", result)
}