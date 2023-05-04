package handler

type ContactHandler interface {
	List()
	Add()
	Detail()
	Update()
	Delete()
}
