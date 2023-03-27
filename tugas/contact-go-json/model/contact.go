package model

type Contact struct {
	ID     int64
	Name   string
	NoTelp string
}

type ContactRequest struct {
	Name   string
	NoTelp string
}

var Contacts []Contact


