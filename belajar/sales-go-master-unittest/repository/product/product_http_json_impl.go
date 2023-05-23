package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sales-go/model"
)

type repositoryjson struct{}

func NewJsonRepository() *repositoryjson {
	return &repositoryjson{}
}

func (repo *repositoryjson) getLastID() (lastID int, err error) {
	listProduct, err := repo.GetList()
	if err != nil {
		return
	}

	if len(listProduct) == 0 {
		lastID = 0
	} else {
		lastID = len(listProduct)
	}
	return
}

func (repo *repositoryjson) GetList() (listProduct []model.Product, err error) {
	reader, err := os.Open("data/product.json")
	if err != nil {
		err = fmt.Errorf("[ERROR] os open product json : %s", err.Error())
		return
	}

	decoder := json.NewDecoder(reader)
	decoder.Decode(&listProduct)

	return
}

func (repo *repositoryjson) updateJSON(listProduct []model.Product) (err error) {
	writerJson, err := os.Create("data/product.json")
	if err != nil {
		err = fmt.Errorf("[ERROR] os create product json : %s", err.Error())
		return
	}
	encodeToJson := json.NewEncoder(writerJson)
	encodeToJson.Encode(listProduct)

	writeTxt, err := os.Create("data/product.txt")
	if err != nil {
		err = fmt.Errorf("[ERROR] os create product txt : %s", err.Error())
		return
	}
	encodeToTxt := json.NewEncoder(writeTxt)
	encodeToTxt.Encode(listProduct)

	return
}

func (repo *repositoryjson) GetProductByName(name string) (productData model.Product, err error) {
	listProduct, err := repo.GetList()
	if err != nil {
		return
	}

	for _, v := range listProduct {
		if v.Name == name {
			productData = v
		}
	}

	emptyStruct := model.Product{}
	if productData == emptyStruct {
		err = errors.New("product not found")
		return
	}
	return
}

func (repo *repositoryjson) Create(req []model.ProductRequest) (result []model.Product, err error) {
	listProduct, err := repo.GetList()
	if err != nil {
		return
	}

	lastID, err := repo.getLastID()
	if err != nil {
		return
	}

	for i, v := range req {
		newProduct := model.Product{
			Id:    lastID + 1 + i,
			Name:  v.Name,
			Price: v.Price,
		}
		listProduct = append(listProduct, newProduct)
		result = append(result, newProduct)
	}

	err = repo.updateJSON(listProduct)
	if err != nil {
		return
	}

	return
}
