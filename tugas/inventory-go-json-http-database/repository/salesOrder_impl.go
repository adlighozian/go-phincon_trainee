package repository

import (
	"encoding/json"
	"errors"
	"inventory/model"
	"math/rand"
	"os"
	"time"
)

type salesOrderRepository struct {
	ProductRepository ProductRepository
}

func NewSalesOrderRepository(productRepository ProductRepository) SalesOrderRepository {
	return &salesOrderRepository{
		ProductRepository: productRepository,
	}
}

func (repo *salesOrderRepository) DecodeSalesOrder() []model.SalesOrderDetail {
	reader, err := os.Open("./assets/salesOrders.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(reader)
	decoder.Decode(&model.SalesOrderDetails)
	return model.SalesOrderDetails
}

func (repo *salesOrderRepository) EncodeSalesOrder() {
	writer, err := os.Create("./assets/salesOrders.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(writer)
	encoder.Encode(model.SalesOrderDetails)
}

func (repo *salesOrderRepository) getIdSales() int {
	model := repo.DecodeSalesOrder()
	tempId := 1
	for _, v := range model {
		tempId = int(v.Id) + 1
	}
	return tempId
}

func (repo *salesOrderRepository) ShowSalesOrderDetail(order string) (model.SalesOrderDetail, error) {
	sod := repo.DecodeSalesOrder()

	var kotak model.SalesOrderDetail
	for _, v := range sod {

		if v.SalesOrder.OrderNumber == order {
			kotak = model.SalesOrderDetail{
				Id:       v.Id,
				Item:     v.Item,
				Price:    v.Price,
				Quantity: v.Quantity,
				SalesOrder: model.SalesOrder{
					OrderNumber: v.SalesOrder.OrderNumber,
					From:        v.SalesOrder.From,
					Total:       v.SalesOrder.Total,
				},
			}
			return kotak, nil
		}
	}

	return kotak, errors.New("kode sales order tidak ditemukan")
}

func (repo *salesOrderRepository) InputSalesOrder(req model.ReqSalesOrder) (model.SalesOrderDetail, error) {
	// Get Id
	var orderDetail model.SalesOrderDetail

	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 7)

	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	rand := string(b)

	switch repo.ProductRepository.SearchItem(req.Item) {
	case false:
		return orderDetail, errors.New("barang tidak ditemukan")
	case true:
		invens := repo.ProductRepository.DecodeProduct()
		var index int
		for i, v := range invens {
			if v.Name == req.Item {
				index = i
			}
		}
		if invens[index].Stock < req.Total {
			return orderDetail, errors.New("stock barang tidak cukup")
		}
		inven := &invens[index]
		inven.Stock = inven.Stock - req.Total
		repo.ProductRepository.EncodeProduct()

		order := model.SalesOrder{
			Id:          repo.getIdSales(),
			OrderNumber: rand,
			From:        req.From,
			Total:       inven.Stock,
		}

		orderDetail = model.SalesOrderDetail{
			Id:         repo.getIdSales(),
			Item:       req.Item,
			Price:      req.Price,
			Quantity:   req.Total,
			Total:      inven.Stock,
			SalesOrder: order,
		}

		model.SalesOrderDetails = append(repo.DecodeSalesOrder(), orderDetail)
		repo.EncodeSalesOrder()
		return orderDetail, nil
	}
	return orderDetail, errors.New("kesalahan input")
}
