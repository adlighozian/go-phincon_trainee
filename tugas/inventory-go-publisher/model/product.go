package model

type Product struct {
	Id    uint   `gorm:"primaryKey;column:id"`
	Name  string `gorm:"column:name"`
	Price int    `gorm:"column:price"`
	Stock int    `gorm:"column:stock"`
}

var Products []Product

type InventoryResponse struct {
	Status  int
	Message string
	Data    any
}
