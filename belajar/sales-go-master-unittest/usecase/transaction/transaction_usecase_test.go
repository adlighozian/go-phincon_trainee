package transaction

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/require"

	"sales-go/model"
	productMock "sales-go/mocks/product"
	transactionMock "sales-go/mocks/transaction"
	voucherMock "sales-go/mocks/voucher"
)

func TestUseCaseTransaction(t *testing.T) {
	t.Run("test get transaction by number success", func(t *testing.T) {
		mockProductSuccess := productMock.NewProductRepoMock()
		mockTransactionSuccess := transactionMock.NewTransactionRepoMock()
		mockVoucherSuccess := voucherMock.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockTransactionSuccess, mockProductSuccess, mockVoucherSuccess)

		// request
		transactionNumber := 288029617

		// mock input output
		mockTransactionSuccess.On("GetTransactionByNumber", transactionNumber).Return([]model.TransactionDetail{
			{
				Id: 1,
				Item: "Tumbler_Phincon",
				Price: 30000,
				Quantity: 3,
				Total: 90000,
				Transaction: model.Transaction{
					Id: 1,
					TransactionNumber: 288029617,
					Name: "Utsman",
					Quantity: 11,
					Discount: 0,
					Total: 480000,
					Pay: 1000000,
				},
			},
			{
				Id: 2,
				Item: "Kaos_Phincon",
				Price: 30000,
				Quantity: 5,
				Total: 150000,
				Transaction: model.Transaction{
					Id: 1,
					TransactionNumber: 288029617,
					Name: "Utsman",
					Quantity: 11,
					Discount: 0,
					Total: 480000,
					Pay: 1000000,
				},
			},
			{
				Id: 3,
				Item: "Lanyard_Phincon",
				Price: 80000,
				Quantity: 3,
				Total: 240000,
				Transaction: model.Transaction{
					Id: 1,
					TransactionNumber: 288029617,
					Name: "Utsman",
					Quantity: 11,
					Discount: 0,
					Total: 480000,
					Pay: 1000000,
				},
			},
		}, nil)

		res, err := usecase.GetTransactionByNumber(transactionNumber)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	t.Run("test get transaction by number failed : transaction number is not a positive number", func(t *testing.T) {
		mockProductSuccess := productMock.NewProductRepoMock()
		mockTransactionSuccess := transactionMock.NewTransactionRepoMock()
		mockVoucherSuccess := voucherMock.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockTransactionSuccess, mockProductSuccess, mockVoucherSuccess)

		// request
		transactionNumber := -288029617

		res, err := usecase.GetTransactionByNumber(transactionNumber)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test get transaction by number failed : get transaction by number", func(t *testing.T) {
		mockProductSuccess := productMock.NewProductRepoMock()
		mockTransactionSuccess := transactionMock.NewTransactionRepoMock()
		mockVoucherSuccess := voucherMock.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockTransactionSuccess, mockProductSuccess, mockVoucherSuccess)

		// request
		transactionNumber := 288029617

		// mock input output
		mockTransactionSuccess.On("GetTransactionByNumber", transactionNumber).Return([]model.TransactionDetail{}, fmt.Errorf("some error"))

		res, err := usecase.GetTransactionByNumber(transactionNumber)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test create bulk transaction detail success", func(t *testing.T) {
		mockProductSuccess := productMock.NewProductRepoMock()
		mockTransactionSuccess := transactionMock.NewTransactionRepoMock()
		mockVoucherSuccess := voucherMock.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockTransactionSuccess, mockProductSuccess, mockVoucherSuccess)

		// request
		voucherCode := "Ph1ncon"
		req := model.TransactionDetailBulkRequest{
			Items: []model.TransactionDetailItemRequest{
				{
					Item:"Tumbler_Phincon",
					Quantity:3,
				},
				{
					Item:"Kaos_Phincon",
					Quantity:5,
				},
				{
					Item:"Lanyard_Phincon",
					Quantity:3,
				},
			},
			Name: "Utsman",
			Pay:  1000000,
		}
		
		// mock input output
		mockProductSuccess.On("GetProductByName", "Kaos_Phincon").Return(model.Product{
			Id: 1,
			Name: "Kaos Phincon",
			Price: 50000,
		}, nil)
		mockProductSuccess.On("GetProductByName", "Lanyard_Phincon").Return(model.Product{
			Id: 2,
			Name: "Lanyard_Phincon",
			Price: 80000,
		}, nil)
		mockProductSuccess.On("GetProductByName", "Tumbler_Phincon").Return(model.Product{
			Id: 3,
			Name: "Tumbler_Phincon",
			Price: 30000,
		}, nil)
		mockVoucherSuccess.On("GetVoucherByCode", voucherCode).Return(model.Voucher{
			Id: 1,
			Code: "Ph1ncon",
			Persen: 20,
		}, nil)
		mockTransactionSuccess.On("CreateBulkTransactionDetail", 
			model.VoucherRequest{
				Code: "Ph1ncon",
				Persen: 20,
			}, 
			[]model.TransactionDetail{
				{
					Id: 0,
					Item: "Tumbler_Phincon",
					Price: 30000,
					Quantity: 3,
					Total: 90000,
					Transaction: model.Transaction{},
				},
				{
					Id: 0,
					Item: "Kaos_Phincon",
					Price: 50000,
					Quantity: 5,
					Total: 250000,
					Transaction: model.Transaction{},
				},
				{
					Id: 0,
					Item: "Lanyard_Phincon",
					Price: 80000,
					Quantity: 3,
					Total: 240000,
					Transaction: model.Transaction{},
				},
			}, req,
		).Return([]model.TransactionDetail{
				{
					Id: 1,
					Item: "Tumbler_Phincon",
					Price: 30000,
					Quantity: 3,
					Total: 90000,
					Transaction: model.Transaction{
						Id: 1,
						TransactionNumber: 288029617,
						Name: "Utsman",
						Quantity: 11,
						Discount: 0,
						Total: 480000,
						Pay: 1000000,
					},
				},
				{
					Id: 2,
					Item: "Kaos_Phincon",
					Price: 30000,
					Quantity: 5,
					Total: 150000,
					Transaction: model.Transaction{
						Id: 1,
						TransactionNumber: 288029617,
						Name: "Utsman",
						Quantity: 11,
						Discount: 0,
						Total: 480000,
						Pay: 1000000,
					},
				},
				{
					Id: 3,
					Item: "Lanyard_Phincon",
					Price: 80000,
					Quantity: 3,
					Total: 240000,
					Transaction: model.Transaction{
						Id: 1,
						TransactionNumber: 288029617,
						Name: "Utsman",
						Quantity: 11,
						Discount: 0,
						Total: 480000,
						Pay: 1000000,
					},
				},
			}, nil)

		res, err := usecase.CreateBulkTransactionDetail(voucherCode, req)
		if err != nil {
			t.Error(err)
		}
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	t.Run("test create bulk transaction detail failed : get product by name", func(t *testing.T) {
		mockProductSuccess := productMock.NewProductRepoMock()
		mockTransactionSuccess := transactionMock.NewTransactionRepoMock()
		mockVoucherSuccess := voucherMock.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockTransactionSuccess, mockProductSuccess, mockVoucherSuccess)

		// request
		voucherCode := "Ph1ncon"
		req := model.TransactionDetailBulkRequest{
			Items: []model.TransactionDetailItemRequest{
				{
					Item:"Tumbler_Phincon",
					Quantity:3,
				},
				{
					Item:"Kaos_Phincon",
					Quantity:5,
				},
				{
					Item:"Lanyard_Phincon",
					Quantity:3,
				},
			},
			Name: "Utsman",
			Pay:  1000000,
		}
		
		// mock input output
		mockProductSuccess.On("GetProductByName", "Tumbler_Phincon").Return(model.Product{}, fmt.Errorf("some error"))

		res, err := usecase.CreateBulkTransactionDetail(voucherCode, req)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test create bulk transaction detail failed : get voucher by code", func(t *testing.T) {
		mockProductSuccess := productMock.NewProductRepoMock()
		mockTransactionSuccess := transactionMock.NewTransactionRepoMock()
		mockVoucherSuccess := voucherMock.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockTransactionSuccess, mockProductSuccess, mockVoucherSuccess)

		// request
		voucherCode := "Ph1ncon"
		req := model.TransactionDetailBulkRequest{
			Items: []model.TransactionDetailItemRequest{
				{
					Item:"Tumbler_Phincon",
					Quantity:3,
				},
				{
					Item:"Kaos_Phincon",
					Quantity:5,
				},
				{
					Item:"Lanyard_Phincon",
					Quantity:3,
				},
			},
			Name: "Utsman",
			Pay:  1000000,
		}
		
		// mock input output
		mockProductSuccess.On("GetProductByName", "Kaos_Phincon").Return(model.Product{
			Id: 1,
			Name: "Kaos Phincon",
			Price: 50000,
		}, nil)
		mockProductSuccess.On("GetProductByName", "Lanyard_Phincon").Return(model.Product{
			Id: 2,
			Name: "Lanyard_Phincon",
			Price: 80000,
		}, nil)
		mockProductSuccess.On("GetProductByName", "Tumbler_Phincon").Return(model.Product{
			Id: 3,
			Name: "Tumbler_Phincon",
			Price: 30000,
		}, nil)
		mockVoucherSuccess.On("GetVoucherByCode", voucherCode).Return(model.Voucher{}, fmt.Errorf("some error"))

		res, err := usecase.CreateBulkTransactionDetail(voucherCode, req)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test create bulk transaction detail failed : item is empty", func(t *testing.T) {
		mockProductSuccess := productMock.NewProductRepoMock()
		mockTransactionSuccess := transactionMock.NewTransactionRepoMock()
		mockVoucherSuccess := voucherMock.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockTransactionSuccess, mockProductSuccess, mockVoucherSuccess)

		// request
		voucherCode := "Ph1ncon"
		req := model.TransactionDetailBulkRequest{
			Items: []model.TransactionDetailItemRequest{
				{
					Item:"",
					Quantity:3,
				},
				{
					Item:"Kaos_Phincon",
					Quantity:5,
				},
				{
					Item:"Lanyard_Phincon",
					Quantity:3,
				},
			},
			Name: "Utsman",
			Pay:  1000000,
		}
		
		// mock input output
		mockProductSuccess.On("GetProductByName", "Kaos_Phincon").Return(model.Product{
			Id: 1,
			Name: "Kaos Phincon",
			Price: 50000,
		}, nil)
		mockProductSuccess.On("GetProductByName", "Lanyard_Phincon").Return(model.Product{
			Id: 2,
			Name: "Lanyard_Phincon",
			Price: 80000,
		}, nil)
		mockProductSuccess.On("GetProductByName", "Tumbler_Phincon").Return(model.Product{
			Id: 3,
			Name: "Tumbler_Phincon",
			Price: 30000,
		}, nil)
		mockVoucherSuccess.On("GetVoucherByCode", voucherCode).Return(model.Voucher{}, fmt.Errorf("some error"))

		res, err := usecase.CreateBulkTransactionDetail(voucherCode, req)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test create bulk transaction detail failed : quantity is not a positive number", func(t *testing.T) {
		mockProductSuccess := productMock.NewProductRepoMock()
		mockTransactionSuccess := transactionMock.NewTransactionRepoMock()
		mockVoucherSuccess := voucherMock.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockTransactionSuccess, mockProductSuccess, mockVoucherSuccess)

		// request
		voucherCode := "Ph1ncon"
		req := model.TransactionDetailBulkRequest{
			Items: []model.TransactionDetailItemRequest{
				{
					Item:"Tumbler_Phincon",
					Quantity:-3,
				},
				{
					Item:"Kaos_Phincon",
					Quantity:5,
				},
				{
					Item:"Lanyard_Phincon",
					Quantity:3,
				},
			},
			Name: "Utsman",
			Pay:  1000000,
		}
		
		// mock input output
		mockProductSuccess.On("GetProductByName", "Kaos_Phincon").Return(model.Product{
			Id: 1,
			Name: "Kaos Phincon",
			Price: 50000,
		}, nil)
		mockProductSuccess.On("GetProductByName", "Lanyard_Phincon").Return(model.Product{
			Id: 2,
			Name: "Lanyard_Phincon",
			Price: 80000,
		}, nil)
		mockProductSuccess.On("GetProductByName", "Tumbler_Phincon").Return(model.Product{
			Id: 3,
			Name: "Tumbler_Phincon",
			Price: 30000,
		}, nil)
		mockVoucherSuccess.On("GetVoucherByCode", voucherCode).Return(model.Voucher{}, fmt.Errorf("some error"))

		res, err := usecase.CreateBulkTransactionDetail(voucherCode, req)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test create bulk transaction detail failed : create bulk transaction detail", func(t *testing.T) {
		mockProductSuccess := productMock.NewProductRepoMock()
		mockTransactionSuccess := transactionMock.NewTransactionRepoMock()
		mockVoucherSuccess := voucherMock.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockTransactionSuccess, mockProductSuccess, mockVoucherSuccess)

		// request
		voucherCode := "Ph1ncon"
		req := model.TransactionDetailBulkRequest{
			Items: []model.TransactionDetailItemRequest{
				{
					Item:"Tumbler_Phincon",
					Quantity:3,
				},
				{
					Item:"Kaos_Phincon",
					Quantity:5,
				},
				{
					Item:"Lanyard_Phincon",
					Quantity:3,
				},
			},
			Name: "Utsman",
			Pay:  1000000,
		}
		
		// mock input output
		mockProductSuccess.On("GetProductByName", "Kaos_Phincon").Return(model.Product{
			Id: 1,
			Name: "Kaos Phincon",
			Price: 50000,
		}, nil)
		mockProductSuccess.On("GetProductByName", "Lanyard_Phincon").Return(model.Product{
			Id: 2,
			Name: "Lanyard_Phincon",
			Price: 80000,
		}, nil)
		mockProductSuccess.On("GetProductByName", "Tumbler_Phincon").Return(model.Product{
			Id: 3,
			Name: "Tumbler_Phincon",
			Price: 30000,
		}, nil)
		mockVoucherSuccess.On("GetVoucherByCode", voucherCode).Return(model.Voucher{
			Id: 1,
			Code: "Ph1ncon",
			Persen: 20,
		}, nil)
		mockTransactionSuccess.On("CreateBulkTransactionDetail", 
			model.VoucherRequest{
				Code: "Ph1ncon",
				Persen: 20,
			}, 
			[]model.TransactionDetail{
				{
					Id: 0,
					Item: "Tumbler_Phincon",
					Price: 30000,
					Quantity: 3,
					Total: 90000,
					Transaction: model.Transaction{},
				},
				{
					Id: 0,
					Item: "Kaos_Phincon",
					Price: 50000,
					Quantity: 5,
					Total: 250000,
					Transaction: model.Transaction{},
				},
				{
					Id: 0,
					Item: "Lanyard_Phincon",
					Price: 80000,
					Quantity: 3,
					Total: 240000,
					Transaction: model.Transaction{},
				},
			}, req,
		).Return([]model.TransactionDetail{}, fmt.Errorf("some error"))

		res, err := usecase.CreateBulkTransactionDetail(voucherCode, req)
		require.Error(t, err)
		require.Empty(t, res)
	})
}