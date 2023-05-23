package model

type Product struct {
	Id    int
	Name  string
	Price float64
}

type ProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductResponseData struct {
	Status  int			`json:"status"`
	Message string		`json:"message"`
	Data	[]Product	`json:"data"`
}

var ProductSlice []Product = []Product{
	{
		Id:    1,
		Name:  "Kaos_Phincon",
		Price: 150000,
	},
	{
		Id:    2,
		Name:  "Lanyard_Phincon",
		Price: 20000,
	},
	{
		Id:    3,
		Name:  "Tumbler_Phincon",
		Price: 80000,
	},
}