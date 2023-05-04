package handler

// import (
// 	"contact-go/mocks"
// 	"contact-go/model"
// 	"net/http"
// 	"testing"
// )

// func TestContactHTTPHandler(t *testing.T) {
// 	t.Run("test get list contact", func(t *testing.T) {
// 		ucMock := mocks.NewUseCaseMock()
// 		handler := NewContactHttpDbHandler(ucMock)

// 		ucMock.On("List").Return(model.ContactResponse{
// 			Status:  http.StatusInternalServerError,
// 			Message: "Internal server error",
// 			Data:    nil,
// 		}, nil)

// 		w := new(http.ResponseWriter)
// 		r := new(http.Request)
// 		handler.List(w, r)
// 	})
// 	t.Run("test create new contact", func(t *testing.T) {
// 		ucMock := mocks.NewUseCaseMock()
// 		handler := NewContactHttpDbHandler(ucMock)

// 		ucMock.On("Add").Return(model.ContactResponse{
// 			Status:  http.StatusInternalServerError,
// 			Message: "Internal server error",
// 			Data:    nil,
// 		}, nil)

// 		w := new(http.ResponseWriter)
// 		r := new(http.Request)
// 		handler.Add(w, r)
// 	})
// 	t.Run("test update contact", func(t *testing.T) {
// 		ucMock := mocks.NewUseCaseMock()
// 		handler := NewContactHttpDbHandler(ucMock)

// 		w := new(http.ResponseWriter)
// 		r := new(http.Request)
// 		handler.Update(w, r)
// 	})
// 	t.Run("test delete contact", func(t *testing.T) {
// 		ucMock := mocks.NewUseCaseMock()
// 		handler := NewContactHttpDbHandler(ucMock)

// 		w := new(http.ResponseWriter)
// 		r := new(http.Request)
// 		handler.Delete(w, r)
// 	})
// }
