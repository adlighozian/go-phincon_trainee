package service

import "inventory/model"

type SalesService interface {
	InputSales(req []model.ReqSales) (model.InventoryResponse, error)
	DetailSales(req string) (model.InventoryResponse, error)
}
