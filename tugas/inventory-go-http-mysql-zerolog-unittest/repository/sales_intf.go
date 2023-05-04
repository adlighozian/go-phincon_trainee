package repository

import "inventory/model"

type SalesRepository interface {
	InputSales(req []model.ReqSales) ([]model.SalesDetail, error)
	ShowSales(req string) (model.SalesDetail, error)
}
