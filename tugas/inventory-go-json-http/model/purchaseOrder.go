package model

type ReqPurchaseOrder struct {
	Item        string
	Price       int
	OrderNumber string
	From        string
	Total       int
}

type ReqPurchase struct {
	Item  string
	Price int
	From  string
	Total int
}

type PurchaseOrderDetail struct {
	Id            int
	Item          string
	Price         int
	Quantity      int
	Total         int
	PurchaseOrder PurchaseOrder
}

type PurchaseOrder struct {
	Id          int
	OrderNumber string
	From        string
	Total       int
}

// Mengganti model

var PurchaseOrderDetails []PurchaseOrderDetail

var PurchaseOrders []PurchaseOrder
