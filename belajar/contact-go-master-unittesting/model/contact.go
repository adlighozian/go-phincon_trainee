package model

type ContactRequest struct {
	Name   string `json:"name"`
	NoTelp string `json:"no_telp"`
}

type Contact struct {
	Id     int
	Name   string
	NoTelp string
}

type ContactResponse struct {
	Status  int			`json:"status"`
	Message string		`json:"message"`
	Data	[]Contact	`json:"data"`
}

var ContactSlice []Contact = []Contact{}