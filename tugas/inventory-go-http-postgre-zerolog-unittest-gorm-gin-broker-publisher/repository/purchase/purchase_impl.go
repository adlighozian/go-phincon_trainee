package purchase

import (
	"errors"
	"inventory/helper"
	"inventory/helper/middleware"
	"inventory/model"
	"inventory/publisher"
	"log"
	"time"

	"gorm.io/gorm"
)

type purchaseRepository struct {
	db *gorm.DB
}

func NewPurchaseRepository(dbs *gorm.DB) PurchaseRepository {
	return &purchaseRepository{
		db: dbs,
	}
}

func (r *purchaseRepository) InputPurchase(req []model.ReqPurchase) ([]model.PurchaseDetail, error) {
	log.Println("purchase repository")

	var err error
	var send []model.SendPurchase
	var returns []model.PurchaseDetail

	for _, v := range req {
		if v.Total <= 0 {
			continue
		}
		sending := model.SendPurchase{
			Item:        v.Item,
			Price:       v.Price,
			From:        v.From,
			Total:       v.Total,
			OrderNumber: helper.Randomizer(),
		}
		send = append(send, sending)
	}

	publisher.PubPurchase(send)

	// get return

	time.Sleep(1 * time.Second)

	tx := r.db.Begin()
	selectPurchaseDetail := `select * from purchase p join purchase_detail pd on p.id = pd.purchase_id where order_number = $1`
	for _, d := range send {
		var returnP model.PurchaseReturn
		err = tx.Raw(selectPurchaseDetail, d.OrderNumber).Scan(&returnP).Error
		if err != nil {
			tx.Rollback()
		}
		resultDetail := model.PurchaseDetail{
			Id:          returnP.Id,
			Purchase_id: returnP.Purchase_id,
			Item:        returnP.Item,
			Price:       returnP.Price,
			Quantity:    returnP.Quantity,
			Total:       returnP.Total,
			Purchase: model.Purchase{
				Id:          returnP.Purchase_id,
				OrderNumber: returnP.OrderNumber,
				From:        returnP.From,
				Total:       returnP.Quantity,
			},
		}
		returns = append(returns, resultDetail)
	}
	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
	}

	return returns, nil
}

func (r *purchaseRepository) DetailPurchase(req string) (model.PurchaseDetail, error) {
	log.Println("purchase repository")
	log.Println(req)

	var returnP model.PurchaseReturn
	var idOrder uint
	var err error

	checkOrder := `select id from purchase where order_number = $1`
	selectPurchaseDetail := `select * from purchase p join purchase_detail pd on p.id = pd.purchase_id where order_number = $1`

	err = r.db.Raw(checkOrder, req).Scan(&idOrder).Error
	middleware.FailError(err, "")

	if idOrder == 0 {
		return model.PurchaseDetail{}, errors.New("order tidak ditemukan")
	}

	err = r.db.Raw(selectPurchaseDetail, req).Scan(&returnP).Error
	middleware.FailError(err, "")

	resultDetail := model.PurchaseDetail{
		Id:          returnP.Id,
		Purchase_id: returnP.Purchase_id,
		Item:        returnP.Item,
		Price:       returnP.Price,
		Quantity:    returnP.Quantity,
		Total:       returnP.Total,
		Purchase: model.Purchase{
			Id:          returnP.Purchase_id,
			OrderNumber: returnP.OrderNumber,
			From:        returnP.From,
			Total:       returnP.Quantity,
		},
	}

	return resultDetail, nil
}
