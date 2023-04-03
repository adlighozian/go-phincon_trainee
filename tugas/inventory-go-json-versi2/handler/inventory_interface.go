package handler

type InventoryHandler interface {
	ShowProduct()
	ShowPurchaseOrderDetail()
	InputPurchaseOrder()
	ShowSalesOrderDetail()
	InputSalesOrder()
}
