package service

import "inventory/model"

type PurchaseService interface {
	InputPurchase(req []model.ReqPurchase) (model.InventoryResponse, error)
	DetailPurchase(req string) (model.InventoryResponse, error)
}
