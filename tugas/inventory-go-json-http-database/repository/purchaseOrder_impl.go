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

func (repo *purchaseOrderRepository) DecodePurchaseOrder() []model.PurchaseOrderDetail {
	reader, err := os.Open("./assets/purchaseOrders.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(reader)
	decoder.Decode(&model.PurchaseOrderDetails)
	return model.PurchaseOrderDetails
}

func (repo *purchaseOrderRepository) EncodePurchaseOrder() {
	writer, err := os.Create("./assets/purchaseOrders.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(writer)
	encoder.Encode(model.PurchaseOrderDetails)
}

func (repo *purchaseOrderRepository) getIdPurchase() int {
	model := repo.DecodePurchaseOrder()
	tempId := 1
	for _, v := range model {
		tempId = int(v.Id) + 1
	}
	return tempId
}

func (repo *purchaseOrderRepository) ShowPurchaseOrderDetail(order string) (model.PurchaseOrderDetail, error) {
	pod := repo.DecodePurchaseOrder()

	var kotak model.PurchaseOrderDetail
	for _, v := range pod {

		if v.PurchaseOrder.OrderNumber == order {
			kotak = model.PurchaseOrderDetail{
				Id:       v.Id,
				Item:     v.Item,
				Price:    v.Price,
				Quantity: v.Quantity,
				PurchaseOrder: model.PurchaseOrder{
					OrderNumber: v.PurchaseOrder.OrderNumber,
					From:        v.PurchaseOrder.From,
					Total:       v.PurchaseOrder.Total,
				},
			}
			return kotak, nil
		}
	}

	return kotak, errors.New("kode purchase order tidak ditemukan")
}

func (repo *purchaseOrderRepository) InputPurchaseOrder(req model.ReqPurchaseOrder) (model.PurchaseOrderDetail, error) {

	var orderDetail model.PurchaseOrderDetail

	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 7)

	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	rand := string(b)

	switch repo.ProductRepository.SearchItem(req.Item) {
	case false:
		product := model.Product{
			Id:    repo.ProductRepository.GetIdProduct(),
			Name:  req.Item,
			Price: req.Price,
			Stock: req.Total,
		}

		order := model.PurchaseOrder{
			Id:          repo.getIdPurchase(),
			OrderNumber: rand,
			From:        req.From,
			Total:       req.Total,
		}

		orderDetail = model.PurchaseOrderDetail{
			Id:            repo.getIdPurchase(),
			Item:          req.Item,
			Price:         req.Price,
			Quantity:      req.Total,
			Total:         req.Total,
			PurchaseOrder: order,
		}

		model.Products = append(repo.ProductRepository.DecodeProduct(), product)
		model.PurchaseOrderDetails = append(repo.DecodePurchaseOrder(), orderDetail)
		repo.ProductRepository.EncodeProduct()
		repo.EncodePurchaseOrder()

		return orderDetail, nil

	case true:

		invens := repo.ProductRepository.DecodeProduct()
		var index int
		for i, v := range invens {
			if v.Name == req.Item {
				index = i
			}
		}
		inven := &invens[index]
		inven.Stock = req.Total + inven.Stock
		repo.ProductRepository.EncodeProduct()

		order := model.PurchaseOrder{
			Id:          repo.getIdPurchase(),
			OrderNumber: rand,
			From:        req.From,
			Total:       inven.Stock,
		}

		orderDetail = model.PurchaseOrderDetail{
			Id:            repo.getIdPurchase(),
			Item:          req.Item,
			Price:         req.Price,
			Quantity:      req.Total,
			Total:         inven.Stock,
			PurchaseOrder: order,
		}

		model.PurchaseOrderDetails = append(repo.DecodePurchaseOrder(), orderDetail)
		repo.EncodePurchaseOrder()
		return orderDetail, nil
	}
	return orderDetail, errors.New("kesalahan input")
}
