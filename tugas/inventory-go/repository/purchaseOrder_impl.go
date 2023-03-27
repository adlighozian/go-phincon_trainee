package repository

import (
	"errors"
	"inventory/model"
	"math/rand"
)

type purchaseOrderRepository struct {
}

func NewPurchaseOrderRepository() PurchaseOrderRepository {
	return &purchaseOrderRepository{}
}

func (repo *purchaseOrderRepository) getIdPurchaseDetail() int {
	model := model.PurchaseOrderDetails
	tempId := 1
	for _, v := range model {
		tempId = int(v.Id) + 1
	}
	return tempId
}

func (repo *purchaseOrderRepository) getIdProduct() int {
	model := model.Products
	tempId := 1
	for _, v := range model {
		tempId = int(v.Id) + 1
	}
	return tempId
}
func (repo *purchaseOrderRepository) getIdPurchase() int {
	model := model.PurchaseOrders
	tempId := 1
	for _, v := range model {
		tempId = int(v.Id) + 1
	}
	return tempId
}

func (repo *purchaseOrderRepository) searchItem(param string) bool {
	model := model.Products
	for _, v := range model {
		if param == v.Name {
			return true
		}
	}
	return false
}

func (repo *purchaseOrderRepository) ShowPurchaseOrderDetail(order int) (model.PurchaseOrder, error) {
	inventory := model.PurchaseOrders
	var kotak model.PurchaseOrder
	for _, v := range inventory {
		if v.OrderNumber == order {
			kotak = model.PurchaseOrder{
				Id:          v.Id,
				OrderNumber: v.OrderNumber,
				From:        v.From,
				Total:       v.Total,
				PurchaseOrderDetail: model.PurchaseOrderDetail{
					Item:     v.PurchaseOrderDetail.Item,
					Price:    v.PurchaseOrderDetail.Price,
					Quantity: v.PurchaseOrderDetail.Quantity,
				},
			}
			return kotak, nil
		}
	}
	return kotak, errors.New("kode order tidak ditemukan")
}

func (repo *purchaseOrderRepository) InputPurchaseOrder(req model.ReqPurchaseOrder) (model.PurchaseOrder, error) {
	// Get Id
	var order model.PurchaseOrder
	randomizer := rand.Intn(100)

	switch repo.searchItem(req.Item) {
	case false:
		product := model.Product{
			Id:    repo.getIdProduct(),
			Name:  req.Item,
			Price: req.Price,
			Stock: req.Total,
		}

		orderDetail := model.PurchaseOrderDetail{
			Id:       repo.getIdPurchaseDetail(),
			Item:     req.Item,
			Price:    req.Price,
			Quantity: req.Total,
			Total:    req.Total,
		}

		order = model.PurchaseOrder{
			Id:                  repo.getIdPurchase(),
			OrderNumber:         randomizer,
			From:                req.From,
			Total:               req.Total,
			PurchaseOrderDetail: orderDetail,
		}

		// model.PurchaseOrderDetails = append(model.PurchaseOrderDetails, orderDetail)
		model.Products = append(model.Products, product)
		model.PurchaseOrders = append(model.PurchaseOrders, order)
		return order, nil
	case true:
		invens := model.Products
		var index int
		for i, v := range invens {
			if v.Name == req.Item {
				index = i
			}
		}
		inven := &invens[index]
		inven.Stock = req.Total + inven.Stock

		orderDetail := model.PurchaseOrderDetail{
			Id:       repo.getIdPurchaseDetail(),
			Item:     req.Item,
			Price:    req.Price,
			Quantity: req.Total,
			Total:    inven.Stock,
		}
		order = model.PurchaseOrder{
			Id:                  repo.getIdPurchase(),
			OrderNumber:         randomizer,
			From:                req.From,
			Total:               inven.Stock,
			PurchaseOrderDetail: orderDetail,
		}

		model.PurchaseOrders = append(model.PurchaseOrders, order)
		return order, nil
	}
	return order, errors.New("kesalahan input")
}
