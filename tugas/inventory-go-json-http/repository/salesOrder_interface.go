package repository

import (
	"inventory/model"
)

type SalesOrderRepository interface {
	ShowSalesOrderDetail(order string) (model.SalesOrderDetail, error)
	InputSalesOrder(req model.ReqSalesOrder) (model.SalesOrderDetail, error)
	DecodeSalesOrder() []model.SalesOrderDetail
}
