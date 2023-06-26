package sales

import (
	"errors"
	"inventory/helper"
	"inventory/helper/middleware"
	"inventory/model"
	"inventory/publisher"
	"time"

	"gorm.io/gorm"
)

type salesRepository struct {
	db        *gorm.DB
	publisher publisher.SalesInterface
	random    helper.RandomInterface
}

func NewSalesRepository(dbs *gorm.DB, publish publisher.SalesInterface, randoms helper.RandomInterface) SalesRepository {
	return &salesRepository{
		db:        dbs,
		publisher: publish,
		random:    randoms,
	}
}

func (r *salesRepository) InputSales(req []model.ReqSales) ([]model.SalesDetail, error) {

	var send []model.SendSales
	var returns []model.SalesDetail

	for _, v := range req {
		if v.Total <= 0 {
			continue
		}
		sending := model.SendSales{
			Item:        v.Item,
			Price:       v.Price,
			From:        v.From,
			Total:       v.Total,
			OrderNumber: r.random.Randomizer(),
		}
		send = append(send, sending)
	}

	r.publisher.PubSales(send)

	// get return
	time.Sleep(1 * time.Second)

	selectSalesDetail := `select * from sales p join sales_detail pd on p.id = pd.sales_id where order_number = $1`

	for _, d := range send {
		var returnP model.SalesReturn
		r.db.Raw(selectSalesDetail, d.OrderNumber).Scan(&returnP)

		// if returnP.Id == 0 {
		// 	continue
		// }

		resultDetail := model.SalesDetail{
			Id:       returnP.Id,
			Sales_id: returnP.Sales_id,
			Item:     returnP.Item,
			Price:    returnP.Price,
			Quantity: returnP.Quantity,
			Total:    returnP.Total,
			Sales: model.Sales{
				Id:          returnP.Sales_id,
				OrderNumber: returnP.OrderNumber,
				From:        returnP.From,
				Total:       returnP.Quantity,
			},
		}
		returns = append(returns, resultDetail)
	}

	if returns == nil {
		return returns, errors.New("error")
	}

	return returns, nil
}

func (r *salesRepository) DetailSales(req string) (model.SalesDetail, error) {
	var returnP model.SalesReturn
	var idOrder uint
	var err error

	checkOrder := `select id from sales where order_number = $1`
	selectsalesDetail := `select * from sales p join sales_detail pd on p.id = pd.sales_id where order_number = $1`

	err = r.db.Raw(checkOrder, req).Scan(&idOrder).Error
	middleware.FailError(err, "")

	if idOrder == 0 {
		return model.SalesDetail{}, errors.New("order tidak ditemukan")
	}

	err = r.db.Raw(selectsalesDetail, req).Scan(&returnP).Error
	middleware.FailError(err, "")

	resultDetail := model.SalesDetail{
		Id:       returnP.Id,
		Sales_id: returnP.Sales_id,
		Item:     returnP.Item,
		Price:    returnP.Price,
		Quantity: returnP.Quantity,
		Total:    returnP.Total,
		Sales: model.Sales{
			Id:          returnP.Sales_id,
			OrderNumber: returnP.OrderNumber,
			From:        returnP.From,
			Total:       returnP.Quantity,
		},
	}

	return resultDetail, nil
}
