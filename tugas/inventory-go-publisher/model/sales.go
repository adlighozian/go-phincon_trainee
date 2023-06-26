package model

type ReqSales struct {
	Item        string
	Price       int
	OrderNumber string
	From        string
	Total       int
}

type SendSales struct {
	Item        string
	Price       int
	From        string
	Total       int
	OrderNumber string `sql:"unique;field:order_number"`
}

type SalesDetail struct {
	Id       uint
	Sales_id uint
	Item     string
	Price    int
	Quantity int
	Total    int
	Sales    Sales
}

type Sales struct {
	Id          uint
	OrderNumber string
	From        string
	Total       int
}

type SalesReturn struct {
	Id          uint   `gorm:"primaryKey;column:id"`
	Sales_id    uint   `gorm:"column:sales_id"`
	Item        string `gorm:"column:item"`
	Price       int    `gorm:"column:price"`
	Quantity    int    `gorm:"column:quantity"`
	OrderNumber string `gorm:"column:order_number"`
	From        string `gorm:"column:orang"`
	Total       int    `grom:"table:sales_detail;column:total"`
}
