package model

type ReqPurchase struct {
	Item        string
	Price       int
	OrderNumber string
	From        string
	Total       int
}

type PurchaseDetail struct {
	Id       int
	Item     string
	Price    int
	Quantity int
	Total    int
	Purchase Purchase
}

type Purchase struct {
	Id          int
	OrderNumber string
	From        string
	Total       int
}

// Mengganti model

var PurchaseDetails []PurchaseDetail

var Purchases []Purchase
