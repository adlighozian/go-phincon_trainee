package model

type Voucher struct {
	Id     int		
	Code   string
	Persen float64
}

type VoucherRequest struct {
	Code   string	`json:"code"`
	Persen float64	`json:"persen"`
}

var VoucherSlice []Voucher = []Voucher{
	{
		Id:    1,
		Code:  "Ph1ncon",
		Persen: 30,
	},
	{
		Id:    2,
		Code:  "Phintraco",
		Persen: 20,
	},
}