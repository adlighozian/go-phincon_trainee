package handler

import (
	"bytes"
	"contact-go/mocks"
	"contact-go/model"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

const sresult string = "Internal Database Error"

func TestHandler(t *testing.T) {
	t.Run("list handler success", func(t *testing.T) {
		mock := mocks.NewHandlerMock()
		handler := NewContactHandlerHttp(mock)

		mock.On("List").Return(model.ContactResponse{
			Status:  http.StatusOK,
			Message: "oke",
			Data: []model.Client{
				{
					Id:     1,
					Name:   "Andi",
					NoTelp: "08987534895",
				},
				{
					Id:     2,
					Name:   "Umar",
					NoTelp: "08987534895",
				},
			},
		}, nil)

		request := httptest.NewRequest("GET", "http://localhost:5000/contact", nil)
		recorder := httptest.NewRecorder()

		handler.HandlerGet(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusOK, response.StatusCode)
	})

	t.Run("list handler fail-1", func(t *testing.T) {
		mock := mocks.NewHandlerMock()
		handler := NewContactHandlerHttp(mock)

		mock.On("List").Return(model.ContactResponse{
			Status:  http.StatusBadGateway,
			Message: sresult,
			Data:    nil,
		}, errors.New("testing error"))

		request := httptest.NewRequest("GET", "http://localhost:5000/contact", nil)
		recorder := httptest.NewRecorder()

		handler.HandlerGet(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusBadGateway, response.StatusCode)
	})

	t.Run("add handler success", func(t *testing.T) {
		mock := mocks.NewHandlerMock()
		handler := NewContactHandlerHttp(mock)

		req := []model.ContactRequest{
			{
				Name:   "adli",
				NoTelp: "1221",
			},
		}

		mock.On("Add", req).Return(model.ContactResponse{
			Status:  http.StatusCreated,
			Message: "oke",
			Data: []model.Client{
				{
					Id:     1,
					Name:   "adli",
					NoTelp: "1221",
				},
			},
		}, nil)

		json, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(json)

		request := httptest.NewRequest("POST", "http://localhost:5000/contact", body)
		recorder := httptest.NewRecorder()

		handler.HandlerPost(recorder, request)

		response := recorder.Result()

		require.Equal(t, http.StatusCreated, response.StatusCode)

	})

	t.Run("add handler fail-1", func(t *testing.T) {
		mock := mocks.NewHandlerMock()
		handler := NewContactHandlerHttp(mock)

		req := []model.ContactRequest{
			{
				Name:   "test",
				NoTelp: "test",
			},
		}

		mock.On("Add", req).Return(model.ContactResponse{
			Status:  http.StatusBadGateway,
			Message: sresult,
			Data:    nil,
		}, errors.New("testing error"))

		json, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(json)

		request := httptest.NewRequest("POST", "http://localhost:5000/contact", body)
		recorder := httptest.NewRecorder()

		handler.HandlerPost(recorder, request)

		response := recorder.Result()

		require.Equal(t, http.StatusBadGateway, response.StatusCode)

	})

	// t.Run("add handler fail-2", func(t *testing.T) {
	// 	mock := mocks.NewHandlerMock()
	// 	handler := NewContactHandlerHttp(mock)

	// 	req := []model.ContactRequest{
	// 		{
	// 			Name:   "",
	// 			NoTelp: "",
	// 		},
	// 	}

	// 	mock.On("Add", req).Return(model.ContactResponse{
	// 		Status:  http.StatusBadRequest,
	// 		Message: "Status Bad Request",
	// 		Data:    nil,
	// 	}, errors.New("data tidak ada"))

	// 	json, err := json.Marshal(req)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	body := bytes.NewReader(json)

	// 	request := httptest.NewRequest("POST", "http://localhost:5000/contact", body)
	// 	recorder := httptest.NewRecorder()

	// 	handler.HandlerPost(recorder, request)

	// 	response := recorder.Result()

	// 	require.Equal(t, http.StatusBadGateway, response.StatusCode)

	// })
}
