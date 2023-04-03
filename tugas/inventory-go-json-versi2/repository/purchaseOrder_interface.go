package repository

import (
	"inventory/model"
)

type PurchaseOrderRepository interface {
	ShowPurchaseOrderDetail(order string) (model.PurchaseOrderDetail, error)
	InputPurchaseOrder(req model.ReqPurchaseOrder) (model.PurchaseOrderDetail, error)
	DecodePurchaseOrder() []model.PurchaseOrderDetail
}
