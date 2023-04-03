package repository

import (
	"errors"
	"fmt"
	"inventory/model"
	"math/rand"
)

type salesOrderRepository struct {
}

func NewSalesOrderRepository() SalesOrderRepository {
	return &salesOrderRepository{}
}

func (repo *salesOrderRepository) getIdSalesDetail() int {
	model := model.SalesOrderDetails
	tempId := 1
	for _, v := range model {
		tempId = int(v.Id) + 1
	}
	return tempId
}

// func (repo *salesOrderRepository) getIdProductSales() int {
// 	model := model.Products
// 	tempId := 1
// 	for _, v := range model {
// 		if tempId < int(v.Id) {
// 			tempId = int(v.Id) + 1
// 		}
// 	}
// 	return tempId
// }

func (repo *salesOrderRepository) getIdSales() int {
	model := model.SalesOrders
	tempId := 1
	for _, v := range model {
		tempId = int(v.Id) + 1
	}
	return tempId
}

func (repo *salesOrderRepository) searchItemSales(param string) bool {
	model := model.Products
	for _, v := range model {
		if param == v.Name {
			return true
		}
	}
	return false
}

func (repo *salesOrderRepository) ShowSalesOrderDetail(order int) (model.SalesOrder, error) {
	inventory := model.SalesOrders
	var kotak model.SalesOrder
	for _, v := range inventory {
		if v.OrderNumber == order {
			kotak = model.SalesOrder{
				Id:          v.Id,
				OrderNumber: v.OrderNumber,
				From:        v.From,
				Total:       v.Total,
				SalesOrderDetail: model.SalesOrderDetail{
					Item:     v.SalesOrderDetail.Item,
					Price:    v.SalesOrderDetail.Price,
					Quantity: v.SalesOrderDetail.Quantity,
				},
			}
			return kotak, nil
		}
	}
	return kotak, errors.New("kode order tidak ditemukan")
}

func (repo *salesOrderRepository) InputSalesOrder(req model.ReqSalesOrder) (model.SalesOrder, error) {
	// Get Id
	var order model.SalesOrder
	randomizer := rand.Intn(100)

	switch repo.searchItemSales(req.Item) {
	case false:
		return order, errors.New("barang tidak ditemukan")
	case true:
		invens := model.Products
		var index int
		for i, v := range invens {
			if v.Name == req.Item {
				index = i
			}
		}
		if invens[index].Stock < req.Total {
			fmt.Println(req.Total)
			return order, errors.New("stock barang tidak cukup")
		}
		inven := &invens[index]
		inven.Stock = inven.Stock - req.Total

		orderDetail := model.SalesOrderDetail{
			Id:       repo.getIdSalesDetail(),
			Item:     req.Item,
			Price:    req.Price,
			Quantity: req.Total,
			Total:    inven.Stock,
		}
		order = model.SalesOrder{
			Id:               repo.getIdSales(),
			OrderNumber:      randomizer,
			From:             req.From,
			Total:            inven.Stock,
			SalesOrderDetail: orderDetail,
		}

		model.SalesOrders = append(model.SalesOrders, order)
		return order, nil
	}
	return order, errors.New("kesalahan input")
}
