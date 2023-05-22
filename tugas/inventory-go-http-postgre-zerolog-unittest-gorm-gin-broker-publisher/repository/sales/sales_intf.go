package sales

import "inventory/model"

type SalesRepository interface {
	InputSales(req []model.ReqSales) ([]model.SalesDetail, error)
	DetailSales(req string) (model.SalesDetail, error)
}
