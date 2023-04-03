package repository

import (
	"encoding/json"
	"errors"
	"inventory/model"
	"math/rand"
	"os"
	"time"
)

type purchaseOrderRepository struct {
	ProductRepository ProductRepository
}

func NewPurchaseOrderRepository(productRepository ProductRepository) PurchaseOrderRepository {
	return &purchaseOrderRepository{
		ProductRepository: productRepository,
	}
}

func (repo *purchaseOrderRepository) DecodePurchaseOrder() []model.PurchaseOrder {
	reader, err := os.Open("./assets/purchaseOrders.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(reader)
	decoder.Decode(&model.PurchaseOrders)
	return model.PurchaseOrders
}

func (repo *purchaseOrderRepository) EncodePurchaseOrder() {
	writer, err := os.Create("./assets/purchaseOrders.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(writer)
	encoder.Encode(model.PurchaseOrders)
}

func (repo *purchaseOrderRepository) getIdProduct() int {
	model := repo.DecodePurchaseOrder()
	tempId := 1
	for _, v := range model {
		tempId = int(v.Id) + 1
	}
	return tempId
}

func (repo *purchaseOrderRepository) getIdPurchase() int {
	model := repo.ProductRepository.DecodeProduct()
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

func (repo *purchaseOrderRepository) ShowPurchaseOrderDetail(order string) (model.PurchaseOrder, error) {
	inventory := repo.DecodePurchaseOrder()
	// repo.EncodePurchaseOrder()
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

	var order model.PurchaseOrder

	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 7)

	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	rand := string(b)

	switch repo.searchItem(req.Item) {
	case false:
		product := model.Product{
			Id:    repo.getIdProduct(),
			Name:  req.Item,
			Price: req.Price,
			Stock: req.Total,
		}

		orderDetail := model.PurchaseOrderDetail{
			Id:       repo.getIdPurchase(),
			Item:     req.Item,
			Price:    req.Price,
			Quantity: req.Total,
			Total:    req.Total,
		}

		order = model.PurchaseOrder{
			Id:                  repo.getIdPurchase(),
			OrderNumber:         rand,
			From:                req.From,
			Total:               req.Total,
			PurchaseOrderDetail: orderDetail,
		}

		// model.PurchaseOrderDetails = append(model.PurchaseOrderDetails, orderDetail)
		model.Products = append(repo.ProductRepository.DecodeProduct(), product)
		model.PurchaseOrders = append(repo.DecodePurchaseOrder(), order)
		repo.EncodePurchaseOrder()
		repo.ProductRepository.EncodeProduct()
		return order, nil
	case true:

		repo.ProductRepository.DecodeProduct()
		invens := model.Products
		var index int
		for i, v := range invens {
			if v.Name == req.Item {
				index = i
			}
		}
		inven := &invens[index]
		inven.Stock = req.Total + inven.Stock
		repo.ProductRepository.EncodeProduct()

		orderDetail := model.PurchaseOrderDetail{
			Id:       repo.getIdPurchase(),
			Item:     req.Item,
			Price:    req.Price,
			Quantity: req.Total,
			Total:    inven.Stock,
		}
		order = model.PurchaseOrder{
			Id:                  repo.getIdPurchase(),
			OrderNumber:         rand,
			From:                req.From,
			Total:               inven.Stock,
			PurchaseOrderDetail: orderDetail,
		}

		model.PurchaseOrders = append(repo.DecodePurchaseOrder(), order)
		repo.EncodePurchaseOrder()
		return order, nil
	}
	return order, errors.New("kesalahan input")
}
