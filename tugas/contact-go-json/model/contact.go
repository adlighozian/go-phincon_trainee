package model

type Contact struct {
	Id     int
	Name   string
	NoTelp string
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

var Contacts []Contact

var ContactReq []ContactRequest
