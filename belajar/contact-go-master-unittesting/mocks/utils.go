package mocks

import (
	"contact-go/model"
)

func NewMockSuccessList() *RepoMock {
	mockSuccess := NewContactRepoMock()
	// Get
	mockSuccess.On("List").Return([]model.Contact{
		{
			Id: 1,
			Name: "Ardi",
			NoTelp: "082828329292",
		},
		{
			Id: 2,
			Name: "Amar",
			NoTelp: "082828329292",
		},
	})
	// Create
	req := []model.ContactRequest{
		{
			Name: "Ardi",
			NoTelp: "082828329292",
		},
		{
			Name: "Amar",
			NoTelp: "082828329292",
		},
	}
	mockSuccess.On("Add", req).Return([]model.Contact{
		{
			Id: 1,
			Name: "Ardi",
			NoTelp: "082828329292",
		},
		{
			Id: 2,
			Name: "Amar",
			NoTelp: "082828329292",
		},
	})
	// Update
	id1 := 2
	req2 := model.ContactRequest{
		Name: "Ardi",
		NoTelp: "082828329292",
	}
	mockSuccess.On("Update", id1, req2).Return(nil)
	// Delete
	id2 := 2
	mockSuccess.On("Delete", id2).Return(nil)
	return mockSuccess
}