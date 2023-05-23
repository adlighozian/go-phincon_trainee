package voucher

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"sales-go/mocks/voucher"
	"sales-go/model"
)

func TestUseCaseVoucher(t *testing.T) {
	t.Run("get list voucher success", func(t *testing.T) {
		mockSuccess := voucher.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// mock input output
		mockSuccess.On("GetList").Return([]model.Voucher{
			{
				Id: 1,
				Code: "ph1ncon",
				Persen: 20,
			},
			{
				Id: 2,
				Code: "VouhcerPhincon",
				Persen: 20,
			},
			{
				Id: 3,
				Code: "VouhcerPhincon",
				Persen: 20,
			},
			{
				Id: 4,
				Code: "VouhcerPhincon",
				Persen: 20,
			},
			{
				Id: 5,
				Code: "Ph1nc0n",
				Persen: 30,
			},
			{
				Id: 6,
				Code: "ph1ncon2",
				Persen: 20,
			},
		}, nil)

		res, err := usecase.GetList()
		if err != nil {
			t.Error(err)
		}

		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	t.Run("get list voucher failed", func(t *testing.T) {
		mockSuccess := voucher.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// mock input output
		mockSuccess.On("GetList").Return([]model.Voucher{}, fmt.Errorf("some error"))

		res, err := usecase.GetList()
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test get voucher by code success", func(t *testing.T) {
		mockSuccess := voucher.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		code := "Ph1ncon"

		// mock input output
		mockSuccess.On("GetVoucherByCode", code).Return(model.Voucher{
			Id: 1,
			Code: "Ph1ncon",
			Persen: 20,
		}, nil)

		res, err := usecase.GetVoucherByCode(code)
		if err != nil {
			t.Error(err)
		}

		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	t.Run("test get voucher by code failed : get voucher by code error", func(t *testing.T) {
		mockSuccess := voucher.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		code := "Ph1ncon"

		// mock input output
		mockSuccess.On("GetVoucherByCode", code).Return(model.Voucher{}, fmt.Errorf("some error"))

		res, err := usecase.GetVoucherByCode(code)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test get voucher by code failed : return empty struct", func(t *testing.T) {
		mockSuccess := voucher.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		code := "Ph1ncon"

		// mock input output
		mockSuccess.On("GetVoucherByCode", code).Return(model.Voucher{}, nil)

		res, err := usecase.GetVoucherByCode(code)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test create voucher success", func(t *testing.T) {
		mockSuccess := voucher.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		req := []model.VoucherRequest{
			{
				Code: "VouhcerPhincon",
				Persen: 20,
			},
			{
				Code: "Ph1nc0n",
				Persen: 30,
			},
		}

		// mock input output
		mockSuccess.On("GetVoucherByCode", "VouhcerPhincon").Return(model.Voucher{}, fmt.Errorf("voucher not found"))
		mockSuccess.On("GetVoucherByCode", "Ph1nc0n").Return(model.Voucher{}, fmt.Errorf("voucher not found"))
		mockSuccess.On("Create", req).Return([]model.Voucher{
			{
				Id: 1,
				Code: "VouhcerPhincon",
				Persen: 20,
			},
			{
				Id: 2,
				Code: "Ph1nc0n",
				Persen: 30,
			},
		}, nil)

		res, err := usecase.Create(req)
		if err != nil {
			t.Error(err)
		}

		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	t.Run("test create voucher failed : some code is empty", func(t *testing.T) {
		mockSuccess := voucher.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		req := []model.VoucherRequest{
			{
				Code: "",
				Persen: 20,
			},
			{
				Code: "Ph1nc0n",
				Persen: 30,
			},
		}

		res, err := usecase.Create(req)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test create voucher failed : some persen is not positive number", func(t *testing.T) {
		mockSuccess := voucher.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		req := []model.VoucherRequest{
			{
				Code: "Phintraco",
				Persen: -1,
			},
			{
				Code: "Ph1nc0n",
				Persen: 30,
			},
		}

		res, err := usecase.Create(req)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test create voucher failed : get voucher by code", func(t *testing.T) {
		mockSuccess := voucher.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		req := []model.VoucherRequest{
			{
				Code: "VouhcerPhincon",
				Persen: 20,
			},
			{
				Code: "Ph1nc0n",
				Persen: 30,
			},
		}

		// mock input output
		mockSuccess.On("GetVoucherByCode", "VouhcerPhincon").Return(model.Voucher{
			Id: 1,
			Code: "VouhcerPhincon",
			Persen: 20,
		}, nil)
		mockSuccess.On("GetVoucherByCode", "Ph1nc0n").Return(model.Voucher{
			Id: 2,
			Code: "Ph1nc0n",
			Persen: 30,	
		}, nil)

		res, err := usecase.Create(req)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test create voucher failed : create voucher", func(t *testing.T) {
		mockSuccess := voucher.NewVoucherRepoMock()
		usecase := NewDBHTTPUsecase(mockSuccess)

		// request
		req := []model.VoucherRequest{
			{
				Code: "VouhcerPhincon",
				Persen: 20,
			},
			{
				Code: "Ph1nc0n",
				Persen: 30,
			},
		}

		// mock input output
		mockSuccess.On("GetVoucherByCode", "VouhcerPhincon").Return(model.Voucher{}, fmt.Errorf("voucher not found"))
		mockSuccess.On("GetVoucherByCode", "Ph1nc0n").Return(model.Voucher{}, fmt.Errorf("voucher not found"))
		mockSuccess.On("Create", req).Return([]model.Voucher{}, fmt.Errorf("some error"))

		res, err := usecase.Create(req)
		require.Error(t, err)
		require.Empty(t, res)
	})
}
