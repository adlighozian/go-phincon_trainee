package product

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/require"

	"sales-go/mocks/product"
	"sales-go/model"
)

func TestUseCaseProduct(t *testing.T) {
	t.Run("test get list product success", func(t *testing.T) {
		mockSuccess := product.NewProductRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		mockSuccess.On("GetList").Return([]model.Product{
			{
				Id: 7,
				Name: "Kaos_Phincon_2",
				Price: 30000,
			},
			{
				Id: 8,
				Name: "Lanyard_Phincon_2",
				Price: 80000,
			},
			{
				Id: 9,
				Name: "Tumbler_Phincon",
				Price: 30000,
			},
		}, nil)

		res, err := usecase.GetList()
		if err != nil {
			t.Error(err)
		}

		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	t.Run("test get list product failed", func(t *testing.T) {
		mockSuccess := product.NewProductRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		mockSuccess.On("GetList").Return([]model.Product{}, fmt.Errorf("some error"))

		res, err := usecase.GetList()
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test get product by name success", func(t *testing.T) {
		mockSuccess := product.NewProductRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		productName := "Kaos Phincon"

		// mock input output
		mockSuccess.On("GetProductByName", productName).Return(model.Product{
			Id: 1,
			Name: "Kaos Phincon",
			Price: 50000,
		}, nil)

		res, err := usecase.GetProductByName(productName)
		if err != nil {
			t.Error(err)
		}

		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	t.Run("test get product by name failed : get product by name error", func(t *testing.T) {
		mockSuccess := product.NewProductRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		productName := "Kaos Phincon"

		// mock input output
		mockSuccess.On("GetProductByName", productName).Return(model.Product{}, fmt.Errorf("some error"))

		res, err := usecase.GetProductByName(productName)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test get product by name failed : return empty struct", func(t *testing.T) {
		mockSuccess := product.NewProductRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		productName := "Kaos Phincon"

		// mock input output
		mockSuccess.On("GetProductByName", productName).Return(model.Product{}, nil)

		res, err := usecase.GetProductByName(productName)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test create product success", func(t *testing.T) {
		mockSuccess := product.NewProductRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		req := []model.ProductRequest{
			{
				Name: "Kaos_Phincon",
				Price: 30000,
			},
			{
				Name: "Lanyard_Phincon",
				Price: 80000,
			},
			{
				Name: "Tumbler_Phincon",
				Price: 30000,
			},
		}

		// mock input output
		mockSuccess.On("GetProductByName", "Kaos_Phincon").Return(model.Product{}, fmt.Errorf("product not found"))
		mockSuccess.On("GetProductByName", "Lanyard_Phincon").Return(model.Product{}, fmt.Errorf("product not found"))
		mockSuccess.On("GetProductByName", "Tumbler_Phincon").Return(model.Product{}, fmt.Errorf("product not found"))
		mockSuccess.On("Create", req).Return([]model.Product{
			{
				Id: 1,
				Name: "Kaos_Phincon",
				Price: 30000,
			},
			{
				Id: 2,
				Name: "Lanyard_Phincon",
				Price: 80000,
			},
			{
				Id: 3,
				Name: "Tumbler_Phincon",
				Price: 30000,
			},
		}, nil)

		res, err := usecase.Create(req)
		if err != nil {
			t.Error(err)
		}

		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	t.Run("test create product fail : some product's name is empty", func(t *testing.T) {
		mockSuccess := product.NewProductRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		req := []model.ProductRequest{
			{
				Name: "",
				Price: 30000,
			},
			{
				Name: "Lanyard_Phincon",
				Price: 80000,
			},
			{
				Name: "Tumbler_Phincon",
				Price: 30000,
			},
		}

		res, err := usecase.Create(req)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test create product fail : some product price is not positive number", func(t *testing.T) {
		mockSuccess := product.NewProductRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		req := []model.ProductRequest{
			{
				Name: "Kaos_Phincon",
				Price: 0,
			},
			{
				Name: "Lanyard_Phincon",
				Price: 80000,
			},
			{
				Name: "Tumbler_Phincon",
				Price: 30000,
			},
		}

		res, err := usecase.Create(req)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test create product fail : product already exist", func(t *testing.T) {
		mockSuccess := product.NewProductRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		req := []model.ProductRequest{
			{
				Name: "Kaos_Phincon",
				Price: 30000,
			},
			{
				Name: "Lanyard_Phincon",
				Price: 80000,
			},
			{
				Name: "Tumbler_Phincon",
				Price: 30000,
			},
		}

		// mock input output
		mockSuccess.On("GetProductByName", "Kaos_Phincon").Return(model.Product{
			Id: 1, 
			Name: "Kaos_Phincon", 
			Price: 30000,
		} , nil)
		mockSuccess.On("GetProductByName", "Lanyard_Phincon").Return(model.Product{
			Id: 2,
			Name: "Lanyard_Phincon", 
			Price: 80000,
		} , nil)
		mockSuccess.On("GetProductByName", "Tumbler_Phincon").Return(model.Product{
			Id: 3,
			Name: "Tumbler_Phincon", 
			Price: 80000,
		}, nil)

		res, err := usecase.Create(req)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test create product fail : create product", func(t *testing.T) {
		mockSuccess := product.NewProductRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		req := []model.ProductRequest{
			{
				Name: "Kaos_Phincon",
				Price: 30000,
			},
			{
				Name: "Lanyard_Phincon",
				Price: 80000,
			},
			{
				Name: "Tumbler_Phincon",
				Price: 30000,
			},
		}

		// mock input output
		mockSuccess.On("GetProductByName", "Kaos_Phincon").Return(model.Product{}, fmt.Errorf("product not found"))
		mockSuccess.On("GetProductByName", "Lanyard_Phincon").Return(model.Product{}, fmt.Errorf("product not found"))
		mockSuccess.On("GetProductByName", "Tumbler_Phincon").Return(model.Product{}, fmt.Errorf("product not found"))
		mockSuccess.On("Create", req).Return([]model.Product{}, fmt.Errorf("some error"))

		res, err := usecase.Create(req)
		require.Error(t, err)
		require.Empty(t, res)
	})
}
