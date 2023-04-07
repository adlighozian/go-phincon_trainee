package model

type ReqSalesOrder struct {
	Item        string
	Price       int
	OrderNumber string
	From        string
	Total       int
}

type SalesOrderDetail struct {
	Id         int
	Item       string
	Price      int
	Quantity   int
	Total      int
	SalesOrder SalesOrder
}

type SalesOrder struct {
	Id          int
	OrderNumber string
	From        string
	Total       int
}

var SalesOrderDetails []SalesOrderDetail

var SalesOrders []SalesOrder
