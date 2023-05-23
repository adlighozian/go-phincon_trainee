package voucher

// import (
// 	"net/http"
// 	"testing"
// 	// "github.com/stretchr/testify/require"

// 	"sales-go/mocks/voucher"
// 	"sales-go/model"
// 	voucherUsecase "sales-go/usecase/voucher"
// )

// func TestGetList(t *testing.T) {
// 	t.Run("test get list voucher", func(t *testing.T) {
// 		mockSuccess := voucher.NewVoucherRepoMock()
// 		usecase := voucherUsecase.NewDBHTTPUsecase(mockSuccess)
// 		handler := NewDBHTTPHandler(usecase)

// 		mockSuccess.On("GetList").Return([]model.Voucher{
// 			{
// 				Id: 1,
// 				Code: "ph1ncon",
// 				Persen: 20,
// 			},
// 			{
// 				Id: 2,
// 				Code: "VouhcerPhincon",
// 				Persen: 20,
// 			},
// 			{
// 				Id: 3,
// 				Code: "VouhcerPhincon",
// 				Persen: 20,
// 			},
// 			{
// 				Id: 4,
// 				Code: "VouhcerPhincon",
// 				Persen: 20,
// 			},
// 		})

		
// 		w := new(http.ResponseWriter)
// 		r := new(*http.Request)

// 		handler.GetList(*w, *r)
// 	})
// }

// func TestCreate(t *testing.T) {
// 	t.Run("test create voucher", func(t *testing.T) {
// 		mockSuccess := voucher.NewVoucherRepoMock()
// 		usecase := voucherUsecase.NewDBHTTPUsecase(mockSuccess)
// 		handler := NewDBHTTPHandler(usecase)

// 		req := []model.VoucherRequest{
// 			{
// 				Code: "VouhcerPhincon",
// 				Persen: 20,
// 			},
// 			{
// 				Code: "Ph1nc0n",
// 				Persen: 30,
// 			},
// 		}

// 		mockSuccess.On("Create", req).Return([]model.Voucher{
// 			{
// 				Id: 4,
// 				Code: "VouhcerPhincon",
// 				Persen: 20,
// 			},
// 			{
// 				Id: 5,
// 				Code: "Ph1nc0n",
// 				Persen: 30,
// 			},
// 		})

// 		w := new(http.ResponseWriter)
// 		r := new(*http.Request)

// 		handler.Create(*w, *r)
// 	})
// }
