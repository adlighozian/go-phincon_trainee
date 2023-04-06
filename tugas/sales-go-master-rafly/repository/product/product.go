package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sales-go/model"
)

type repository struct {}

func NewRepository() *repository {
	return &repository{}
}

type Repositorier interface {
	GetList() (listProduct []model.Product, err error)
	UpdateJSON(listProduct []model.Product) (err error)
	GetProductByName(name string) (productData model.Product, err error)
	Create(req model.ProductRequest) (newData model.Product, err error)
}

func (repo *repository) getLastID() (lastID int, err error) {
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

func (repo *repository) GetList() (listProduct []model.Product, err error) {
	reader, err := os.Open("data/product.json")
	if err != nil {
		err = errors.New(fmt.Sprintf("[ERROR] os open product json : %s", err.Error()))
		return
	}
	
	decoder := json.NewDecoder(reader)
	decoder.Decode(&listProduct)
	
	return
}

func (repo *repository) UpdateJSON(listProduct []model.Product) (err error) {
	writerJson, err := os.Create("data/product.json")
	if err != nil {
		err = errors.New(fmt.Sprintf("[ERROR] os create product json : %s", err.Error()))
		return
	}
	encodeToJson := json.NewEncoder(writerJson)
	encodeToJson.Encode(listProduct)

	writeTxt, err := os.Create("data/product.txt")
	if err != nil {
		err = errors.New(fmt.Sprintf("[ERROR] os create product txt : %s", err.Error()))
		return
	}
	encodeToTxt := json.NewEncoder(writeTxt)
	encodeToTxt.Encode(listProduct)

	return
}

func (repo *repository) GetProductByName(name string) (productData model.Product, err error) {
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
		err = errors.New("Product not found")
		return
	}
	return
}

func (repo *repository) Create(req model.ProductRequest) (newData model.Product, err error) {
	listProduct, err := repo.GetList()
	if err != nil {
		return
	}

	lastID, err := repo.getLastID()
	if err != nil {
		return
	}

	newData = model.Product{
		Id:    lastID + 1,
		Name:  req.Name,
		Price: req.Price,
	}
	listProduct = append(listProduct, newData)

	err = repo.UpdateJSON(listProduct)
	if err != nil {
		return
	}

	return
}
