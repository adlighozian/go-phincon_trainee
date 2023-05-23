package product

import (
	"sales-go/model"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct{
	mock.Mock
}

func NewProductRepoMock() *RepoMock {
	return &RepoMock{}
}

func (m *RepoMock) GetList() (listProduct []model.Product, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called()
	listProduct = ret.Get(0).([]model.Product)
	err = ret.Error(1)
	return listProduct, err
}

func (m *RepoMock) GetProductByName(name string) (productData model.Product, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(name)
	productData = ret.Get(0).(model.Product)
	err = ret.Error(1)
	return productData, err
}

func (m *RepoMock) Create(req []model.ProductRequest) (result []model.Product, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(req)
	result = ret.Get(0).([]model.Product)
	err = ret.Error(1)
	return result, err
}
