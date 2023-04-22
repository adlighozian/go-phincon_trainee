package product

import (
	"fmt"
	"errors"
	"sales-go/model"
	"sales-go/repository/product"
)

type usecase struct {
	repo product.Repositorier
}

func NewDBHTTPUsecase(repository product.Repositorier) *usecase {
	return &usecase{
		repo: repository,
	}
}

func (uc *usecase) GetList() (response []model.Product, err error) {
	return uc.repo.GetList()
}

func (uc *usecase) GetProductByName(name string) (response model.Product, err error) {
	response, err = uc.repo.GetProductByName(name)
	if err != nil {
		return
	}

	emptyStruct := model.Product{}
	if response == emptyStruct {
		err = errors.New("product not found")
		return
	}
	return
}

func (uc *usecase) Create(req []model.ProductRequest) (response []model.Product, err error) {
	for _, product := range req {
		if product.Name == "" {
			err = fmt.Errorf("product %s : name should not be empty", product.Name)
			return
		} else if product.Price <= 0 {
			err = fmt.Errorf("product %s : price should be > 0", product.Name)
		} else {
			_, err = uc.GetProductByName(product.Name)
			if err != nil {
				continue
			} else {
				fmt.Println(err)
				err = fmt.Errorf("product %s already exist", product.Name)
				return
			}
		}
	}

	response, err = uc.repo.Create(req)
	if err != nil {
		return
	}
	return
}