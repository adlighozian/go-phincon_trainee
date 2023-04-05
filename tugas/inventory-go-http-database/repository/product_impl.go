package repository

import (
	"encoding/json"
	"inventory/model"
	"os"
)

type productRepository struct {
}

func NewProductRepository() ProductRepository {
	return new(productRepository)
}

func (repo *productRepository) DecodeProduct() []model.Product {
	reader, err := os.Open("./assets/products.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(reader)
	decoder.Decode(&model.Products)
	return model.Products
}

func (repo *productRepository) EncodeProduct() {
	writer, err := os.Create("./assets/products.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(writer)
	encoder.Encode(model.Products)
}

func (repo *productRepository) GetIdProduct() int {
	model := repo.DecodeProduct()
	tempId := 1
	for _, v := range model {
		tempId = int(v.Id) + 1
	}
	return tempId
}

func (repo *productRepository) SearchItem(param string) bool {
	model := repo.DecodeProduct()
	for _, v := range model {
		if param == v.Name {
			return true
		}
	}
	return false
}

func (repo *productRepository) ShowProduct() []model.Product {
	return repo.DecodeProduct()
}
