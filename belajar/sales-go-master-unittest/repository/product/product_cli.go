package product

import (
	"errors"
	"sales-go/model"
)

type repositorycli struct {}

func NewCLIRepository() *repositorycli {
	return &repositorycli{}
}

func (repo *repositorycli) GetList() (listProduct []model.Product, err error) {
	return model.ProductSlice, nil
}

func (repo *repositorycli) GetProductByName(name string) (productData model.Product, err error) {
	for _, v := range model.ProductSlice {
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

func (repo *repositorycli) Create(req []model.ProductRequest) (result []model.Product, err error) {
	for _, v := range req {
		newData := model.Product{
			Id:    len(model.ProductSlice) + 1,
			Name:  v.Name,
			Price: v.Price,
		}
		model.ProductSlice = append(model.ProductSlice, newData)
		result = append(result, newData)
	}
	return
}