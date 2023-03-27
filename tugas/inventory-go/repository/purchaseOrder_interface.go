package repository

import (
	"inventory/model"
)

type PurchaseOrderRepository interface {
	ShowPurchaseOrderDetail(order int) (model.PurchaseOrder, error)
	InputPurchaseOrder(req model.ReqPurchaseOrder) (model.PurchaseOrder, error)
}
