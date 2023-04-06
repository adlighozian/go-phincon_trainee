package model

type Transaction struct {
	Id                int
	TransactionNumber int
	Name			  string
	Quantity          int
	Discount          float64
	Total             float64
	Pay				  float64
}

type TransactionDetail struct {
	Id          int
	Item        string
	Price       float64
	Quantity    int
	Total       float64
	Transaction Transaction
}

type TransactionRequest struct {
	Name		string	`json:"name"`
	Quantity	int		`json:"quantity"`
	Discount	float64	`json:"discount"`
	Total		float64	`json:"total"`
	Pay			float64	`json:"pay"`
}

type TransactionDetailItemRequest struct {	
	Item        string	`json:"item"`
	Quantity    int		`json:"quantity"`
}

type TransactionDetailBulkRequest struct {
	Items	[]TransactionDetailItemRequest
	Name	string `json:"name"`
	Pay		float64 `json:"pay"`
}

var Total float64

var TransactionSlice []Transaction = []Transaction{}

var TransactionDetailSlice []TransactionDetail = []TransactionDetail{}
