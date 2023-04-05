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

var Contacts []Contact
