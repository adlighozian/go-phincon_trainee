package transaction

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"sales-go/model"
	"sales-go/repository/product"
	"sales-go/repository/transaction"
	"sales-go/repository/voucher"
	"strconv"
	"strings"
)

type jsonhttphandler struct {
	repo        transaction.Repositorier
	productrepo product.Repositorier
	voucherrepo voucher.Repositorier
}

func NewJsonHTTPHandler(
	repositorier transaction.Repositorier,
	productRepository product.Repositorier,
	voucherRepository voucher.Repositorier,
) *jsonhttphandler {
	return &jsonhttphandler{
		repo:        repositorier,
		productrepo: productRepository,
		voucherrepo: voucherRepository,
	}
}

func (handler *jsonhttphandler) GetTransactionByNumber(w http.ResponseWriter, r *http.Request) {
	transactionNumberStr := r.URL.Query().Get("transaction_id")
	transactionNumber, err := strconv.Atoi(transactionNumberStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
		log.Println("[ERROR] convert string id to integer:", err.Error())
		return
	}
	if transactionNumber < 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("message : id must be > 0"))
		log.Println("[ERROR] transaction id must be > 0:", err.Error())
		return
	}

	result, err := handler.repo.GetTransactionByNumber(transactionNumber)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
		log.Println("[ERROR] get transaction by number :", err.Error())
		return
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
		log.Println("[ERROR] marshal list transaction :", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))
}

func (handler *jsonhttphandler) CreateBulkTransactionDetail(w http.ResponseWriter, r *http.Request) {
	req := model.TransactionDetailBulkRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
		log.Println("[ERROR] decode request :", err.Error())
		return
	}

	listTransactionDetail := []model.TransactionDetail{}
	for _, v := range req.Items {
		product, err := handler.productrepo.GetProductByName(v.Item)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
			log.Println("[ERROR] get product by name :", err.Error())
			return
		}
	
		if v.Quantity < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("[ERROR] quantity should not negative."))
			log.Println("[ERROR] quantity should not negative.")
			return
		}

		listTransactionDetail = append(listTransactionDetail, model.TransactionDetail{
			Item:     v.Item,
			Price:    product.Price,
			Quantity: v.Quantity,
			Total:    product.Price * float64(v.Quantity),
		})
	}

	voucherData := model.VoucherRequest{}
	voucherCode := r.URL.Query().Get("voucher_code")
	if voucherCode != "" {
		voucher, err := handler.voucherrepo.GetVoucherByCode(voucherCode)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
			log.Println("[ERROR] get voucher by code :", err.Error())
			return
		} else if err == nil {
			voucherData = model.VoucherRequest{
				Code:   voucher.Code,
				Persen: voucher.Persen,
			}
		}
	}

	// input transaction detail to transaction slice
	res, err := handler.repo.CreateBulkTransactionDetail(voucherData, listTransactionDetail, req)
	if err != nil {
		if strings.EqualFold(err.Error(), "pay must be > total") {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
			log.Println("[ERROR] create bulk transaction :", err.Error())
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("message : %s", err.Error())))
		log.Println("[ERROR] create bulk transaction :", err.Error())
		return
	}

	var ids []int
	for _, v := range res {
		ids = append(ids, v.Id)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Created new transaction with ids : %d", ids)))
}
