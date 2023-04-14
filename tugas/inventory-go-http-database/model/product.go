package model

type Product struct {
	Id    int
	Name  string
	Price int
	Stock int
}

var Products []Product

type InventoryResponse struct {
	Status  int
	Message string
	Data    any
}
