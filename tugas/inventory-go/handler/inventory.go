package handler

import (
	"fmt"
	"inventory/helper"
	"inventory/model"
	"inventory/repository"
)

type inventoryHandler struct {
	ProductRepository       repository.ProductRepository
	PurchaseOrderRepository repository.PurchaseOrderRepository
	SalesOrderRepository    repository.SalesOrderRepository
}

func NewInventoryHandler(productRepository repository.ProductRepository, purchaseOrderRepository repository.PurchaseOrderRepository, salesOrderRepository repository.SalesOrderRepository) InventoryHandler {
	return &inventoryHandler{
		ProductRepository:       productRepository,
		PurchaseOrderRepository: purchaseOrderRepository,
		SalesOrderRepository:    salesOrderRepository,
	}
}

func (handler *inventoryHandler) ShowProduct() {
	helper.CallClear()
	fmt.Println("Show product")
	fmt.Println("==========================")
	fmt.Printf("ID\t\t| Name\t\t| Price\t\t| Stock\t\t\n")
	inventory := handler.ProductRepository.ShowProduct()
	for _, v := range inventory {
		fmt.Printf("%d\t\t| %s\t\t| %d\t\t| %d\t\t\n", v.Id, v.Name, v.Price, v.Stock)
	}
}

func (handler *inventoryHandler) ShowPurchaseOrderDetail() {
	helper.CallClear()
	fmt.Println("Show purchase order detail")
	fmt.Println("==========================")
	model := model.PurchaseOrders
	fmt.Println("Order Number")
	for _, v := range model {
		fmt.Printf("%d ", v.OrderNumber)
	}
	fmt.Println()
	fmt.Println()
	fmt.Print("Masukkan order mumber: ")
	var orderNumber int
	fmt.Scanln(&orderNumber)
	inventory, err := handler.PurchaseOrderRepository.ShowPurchaseOrderDetail(orderNumber)

	helper.CallClear()
	fmt.Printf("Order number %d", orderNumber)
	fmt.Println()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Order Number\t\t| From\t\t| Item\t\t| Price\t\t| Total\t\t\n")
		fmt.Printf("%d\t\t| %s\t\t| %s\t\t| %d\t\t| %d\t\t\n", inventory.OrderNumber, inventory.From, inventory.PurchaseOrderDetail.Item, inventory.PurchaseOrderDetail.Price, inventory.Total)
	}

}

func (handler *inventoryHandler) InputPurchaseOrder() {
	helper.CallClear()
	var item string
	var price int
	var from string
	var total int
	fmt.Println("Input purchase order")
	fmt.Println("==========================")
	fmt.Print("Nama barang: ")
	fmt.Scan(&item)
	fmt.Print("Nama orang: ")
	fmt.Scan(&from)
	fmt.Print("Harga: ")
	fmt.Scan(&price)
	fmt.Print("jumlah total: ")
	fmt.Scan(&total)

	inputReq := model.ReqPurchaseOrder{
		Item:  item,
		Price: price,
		From:  from,
		Total: total,
	}
	var test int
	val, err := handler.PurchaseOrderRepository.InputPurchaseOrder(inputReq)
	if err != nil {
		fmt.Println()
		fmt.Println(err)
		fmt.Scanln(&test)
	} else {
		helper.CallClear()
		fmt.Println("clear input 2")
		fmt.Printf("Order Number\t| From\t| Item\t| Price\t| Quantity\t| Total\t\n")
		fmt.Printf("%d\t| %s\t| %s\t| %d\t| %d\t| %d\t\n", val.OrderNumber, val.From, val.PurchaseOrderDetail.Item, val.PurchaseOrderDetail.Price, val.Quantity, val.Total)
		fmt.Scanln(&test)
	}
	fmt.Println()
	fmt.Println("2. utnuk input lagi")
}

func (handler *inventoryHandler) ShowSalesOrderDetail() {
	helper.CallClear()
	fmt.Println("Show sales order detail")
	fmt.Println("==========================")
	model := model.SalesOrders
	fmt.Println("Order Number")
	for _, v := range model {
		fmt.Printf("%d ", v.OrderNumber)
	}
	fmt.Println()
	fmt.Println()
	fmt.Print("Masukkan order mumber: ")
	var orderNumber int
	fmt.Scanln(&orderNumber)
	inventory, err := handler.SalesOrderRepository.ShowSalesOrderDetail(orderNumber)

	helper.CallClear()
	fmt.Printf("Order number %d", orderNumber)
	fmt.Println()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Order Number\t\t| From\t\t| Item\t\t| Price\t\t| Total\t\t\n")
		fmt.Printf("%d\t\t| %s\t\t| %s\t\t| %d\t\t| %d\t\t\n", inventory.OrderNumber, inventory.From, inventory.SalesOrderDetail.Item, inventory.SalesOrderDetail.Price, inventory.Total)
	}

}

func (handler *inventoryHandler) InputSalesOrder() {
	helper.CallClear()
	var item string
	var price int
	var from string
	var total int
	fmt.Println("Input sales order")
	fmt.Println("==========================")
	fmt.Print("Nama barang: ")
	fmt.Scan(&item)
	fmt.Print("Nama orang: ")
	fmt.Scan(&from)
	fmt.Print("Harga: ")
	fmt.Scan(&price)
	fmt.Print("jumlah total: ")
	fmt.Scan(&total)

	inputReq := model.ReqSalesOrder{
		Item:  item,
		Price: price,
		From:  from,
		Total: total,
	}
	var test int
	val, err := handler.SalesOrderRepository.InputSalesOrder(inputReq)
	if err != nil {
		fmt.Println()
		fmt.Println(err)
		fmt.Scanln(&test)
	} else {
		helper.CallClear()
		// fmt.Println("clear input 2")
		fmt.Printf("Order Number\t| From\t| Item\t| Price\t| Quantity\t| Total\t\n")
		fmt.Printf("%d\t| %s\t| %s\t| %d\t| %d\t| %d\t\n", val.OrderNumber, val.From, val.SalesOrderDetail.Item, val.SalesOrderDetail.Price, val.Quantity, val.Total)
		fmt.Scanln(&test)
	}
	fmt.Println()
	fmt.Println("3. utnuk input lagi")
}
