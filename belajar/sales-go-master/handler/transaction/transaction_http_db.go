package transaction

import (
	"encoding/json"
	"net/http"

	"sales-go/helpers/rest"
	"sales-go/model"
	"sales-go/usecase/transaction"
	"strconv"
)

type dbhttphandler struct {
	usecase        transaction.TransactionUseCase
}

func NewDBHTTPHandler(
	usecasesitorier transaction.TransactionUseCase,
) *dbhttphandler {
	return &dbhttphandler{
		usecase:        usecasesitorier,
	}
}

func (handler *dbhttphandler) GetTransactionByNumber(w http.ResponseWriter, r *http.Request) {
	transactionNumberStr := r.URL.Query().Get("transaction_id")
	transactionNumber, err := strconv.Atoi(transactionNumberStr)
	if err != nil {
		rest.ResponseError(w, r, http.StatusInternalServerError, err)
		return
	}

	response, err := handler.usecase.GetTransactionByNumber(transactionNumber)
	if err != nil {
		if err.Error() == "id must be > 0" {
			rest.ResponseError(w, r, http.StatusBadRequest, err)
			return
		} else {
			rest.ResponseError(w, r, http.StatusInternalServerError, err)
			return
		}
	}
	rest.ResponseData(w, r, http.StatusOK, response)
}

func (handler *dbhttphandler) CreateBulkTransactionDetail(w http.ResponseWriter, r *http.Request) {
	req := model.TransactionDetailBulkRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rest.ResponseError(w, r, http.StatusInternalServerError, err)
		return
	}

	voucherCode := r.URL.Query().Get("voucher_code")

	res, err := handler.usecase.CreateBulkTransactionDetail(voucherCode, req)
	if err != nil {
		if err.Error() == "quantity transaction should not be negative" || err.Error() == "item transaction hould not be empty" {
			rest.ResponseError(w, r, http.StatusBadRequest, err)
			return
		} else {
			rest.ResponseError(w, r, http.StatusInternalServerError, err)
			return
		}
	}
	rest.ResponseData(w, r, http.StatusCreated, res)
}
