package repository

import (
	"inventory/model"
)

type SalesOrderRepository interface {
	ShowSalesOrderDetail(order int) (model.SalesOrder, error)
	InputSalesOrder(req model.ReqSalesOrder) (model.SalesOrder, error)
}
