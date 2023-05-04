package model

type ReqSales struct {
	Item        string
	Price       int
	OrderNumber string
	From        string
	Total       int
}

type SalesDetail struct {
	Id       int
	Item     string
	Price    int
	Quantity int
	Total    int
	Sales    Sales
}

type Sales struct {
	Id          int
	OrderNumber string
	From        string
	Total       int
}

var SalesDetails []SalesDetail

var Saless []Sales
