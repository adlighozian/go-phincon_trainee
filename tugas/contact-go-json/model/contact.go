package model

// type Contact struct {
// 	Id     int `gorm:"primaryKey"`
// 	Name   string
// 	NoTelp string
// }

type Client struct {
	Id     int    `gorm:"column:id; primaryKey"`
	Name   string `gorm:"column:name" json:"name"`
	NoTelp string `gorm:"column:no_telp" json:"no_telp"`
}

type ContactRequest struct {
	Name   string
	NoTelp string
}

type ContactResponse struct {
	Status  int
	Message string
	Data    any
}

var Contacts []Client

var ContactReq []ContactRequest
