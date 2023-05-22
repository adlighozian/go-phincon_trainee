package model

// type ReqPurchase struct {
// 	Item  string
// 	Price int
// 	From  string
// 	Total int
// }

type SendPurchase struct {
	Item        string
	Price       int
	From        string
	Total       int
	OrderNumber string
}

type PurchaseDetail struct {
	Id          uint     `gorm:"primaryKey;column:id"`
	Purchase_id uint     `gorm:"column:purchase_id"`
	Item        string   `gorm:"column:id"`
	Price       int      `gorm:"column:price"`
	Quantity    int      `gorm:"column:quantity"`
	Total       int      `gorm:"column:total"`
	Purchase    Purchase `gorm:"foreignKey:Purchase_id"`
}

type Purchase struct {
	Id          uint   `gorm:"primaryKey;column:id"`
	OrderNumber string `gorm:"column:order_number"`
	From        string `gorm:"column:orang"`
	Total       int    `grom:"column:total"`
}

var PurchaseDetails []PurchaseDetail

var Purchases []Purchase
