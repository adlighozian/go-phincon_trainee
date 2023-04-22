package transaction

import (
	"net/http"
)

type Handlerer interface {
	GetTransactionByNumber(w http.ResponseWriter, r *http.Request)
	CreateBulkTransactionDetail(w http.ResponseWriter, r *http.Request)
}
