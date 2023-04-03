package model

type ReqSalesOrder struct {
	Item        string
	Price       int
	OrderNumber string
	From        string
	Total       int
}

type SalesOrder struct {
	Id          int
	OrderNumber int
	From        string
	Total       int
	SalesOrderDetail
}

type SalesOrderDetail struct {
	Id       int
	Item     string
	Price    int
	Quantity int
	Total    int
}

var SalesOrderDetails []SalesOrderDetail

var SalesOrders []SalesOrder
