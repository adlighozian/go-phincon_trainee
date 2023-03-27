package model

type ReqPurchaseOrder struct {
	Item        string
	Price       int
	OrderNumber string
	From        string
	Total       int
}

type PurchaseOrder struct {
	Id          int
	OrderNumber int
	From        string
	Total       int
	PurchaseOrderDetail
}

type PurchaseOrderDetail struct {
	Id       int
	Item     string
	Price    int
	Quantity int
	Total    int
}

// Mengganti model

var PurchaseOrderDetails []PurchaseOrderDetail

var PurchaseOrders []PurchaseOrder
