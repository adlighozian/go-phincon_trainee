package service

import (
	"fmt"
	"testing"
	"time"
	"unit-testing/model"
	"unit-testing/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestMain(m *testing.M) {
	fmt.Println("start test")
	start := time.Now()
	m.Run()
	fmt.Println("end test")
	end := time.Since(start)
	fmt.Println("duration", end)
}

func TestCategoryServise_GetNotFound(t *testing.T) {

	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}

// func TestCategoryServise(t *testing.T) {

// 	categoryRepository.Mock.On("Finda", "1").Return(nil)
// 	category, err := categoryService.Get("1")

// 	assert.Nil(t, category)
// 	assert.NotNil(t, err)
// }

func TestCategoryServise_GetSuccess(t *testing.T) {
	category := model.Category{
		Id:   "2",
		Name: "baju",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}

func TestBarang(t *testing.T) {

	for _, v := range Barangs {
		categoryRepository.Mock.On("FindBarang", v.Input).Return(v.Expect)
		t.Run(v.SubTest, func(t *testing.T) {
			result := categoryService.GetBarang(v.Input)
			require.Equal(t, v.Expect, result, v.Message)
		})
	}

}

var Barangs []barang = []barang{
	{
		Input:   "1",
		SubTest: "Berhasil",
		Message: "masukkan id yang benar",
		Expect:  true,
	},
	{
		Input:   "2",
		SubTest: "salah",
		Message: "masukkan id yang salah",
		Expect:  false,
	},
}

type barang struct {
	Input   string
	SubTest string
	Message string
	Expect  bool
}
