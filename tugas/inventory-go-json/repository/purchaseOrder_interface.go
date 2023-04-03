package repository

import (
	"inventory/model"
)

type PurchaseOrderRepository interface {
	ShowPurchaseOrderDetail(order string) (model.PurchaseOrder, error)
	InputPurchaseOrder(req model.ReqPurchaseOrder) (model.PurchaseOrder, error)
	DecodePurchaseOrder() []model.PurchaseOrder
}
