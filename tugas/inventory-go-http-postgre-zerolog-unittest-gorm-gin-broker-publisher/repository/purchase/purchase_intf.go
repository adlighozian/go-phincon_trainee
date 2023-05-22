package purchase

import "inventory/model"

type PurchaseRepository interface {
	InputPurchase(req []model.ReqPurchase) ([]model.PurchaseDetail, error)
	DetailPurchase(req string) (model.PurchaseDetail, error)
}
