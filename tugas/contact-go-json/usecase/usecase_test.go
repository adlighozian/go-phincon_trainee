package usecase

import (
	"errors"
	"net/http"
	"testing"

	"contact-go/mocks"
	"contact-go/model"

	"github.com/stretchr/testify/require"
)

func TestUseCaseHTTP(t *testing.T) {
	t.Run("get-list", func(t *testing.T) {
		mockSuccess := mocks.NewRepoMock()
		uc := NewContactUseCase(mockSuccess)

		mockSuccess.On("List").Return([]model.Client{
			{
				Id:     1,
				Name:   "Ardi",
				NoTelp: "082828329292",
			},
			{
				Id:     2,
				Name:   "Amar",
				NoTelp: "082828329292",
			},
		}, nil)

		res, err := uc.List()
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
		require.Equal(t, "oke", res.Message)
	})

	t.Run("get-list-fail", func(t *testing.T) {
		mockSuccess := mocks.NewRepoMock()
		uc := NewContactUseCase(mockSuccess)

		mockSuccess.On("List").Return([]model.Client{}, errors.New("error"))

		res, _ := uc.List()

		require.Equal(t, http.StatusBadGateway, res.Status)
		require.Equal(t, "Internal Database Error", res.Message)
	})

	t.Run("add-contact", func(t *testing.T) {
		mockSuccess := mocks.NewRepoMock()
		uc := NewContactUseCase(mockSuccess)

		req := []model.ContactRequest{
			{
				Name:   "Ardi",
				NoTelp: "082828329292",
			},
			{
				Name:   "Amar",
				NoTelp: "082828329292",
			},
		}

		mockSuccess.On("Add", req).Return([]model.Client{
			{
				Id:     1,
				Name:   "Ardi",
				NoTelp: "082828329292",
			},
			{
				Id:     2,
				Name:   "Amar",
				NoTelp: "082828329292",
			},
		}, nil)

		res, err := uc.Add(req)

		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, res.Status)
		require.Equal(t, "oke", res.Message)
	})

	t.Run("add-contact-fail", func(t *testing.T) {
		mockSuccess := mocks.NewRepoMock()
		uc := NewContactUseCase(mockSuccess)

		req := []model.ContactRequest{
			{
				Name:   "Ardi",
				NoTelp: "082828329292",
			},
		}

		mockSuccess.On("Add", req).Return([]model.Client{}, errors.New("error"))

		res, _ := uc.Add(req)

		require.Equal(t, http.StatusBadGateway, res.Status)
		require.Equal(t, "Internal Database Error", res.Message)
	})

	t.Run("update-contact", func(t *testing.T) {
		mockSuccess := mocks.NewRepoMock()
		uc := NewContactUseCase(mockSuccess)

		id := 2
		req := model.ContactRequest{
			Name:   "Ardi",
			NoTelp: "082828329292",
		}
		mockSuccess.On("Update", id, req).Return(nil)

		res, err := uc.Update(id, req)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
		require.Equal(t, "oke", res.Message)
	})

	t.Run("update-contact-fail", func(t *testing.T) {
		mockSuccess := mocks.NewRepoMock()
		uc := NewContactUseCase(mockSuccess)

		id := -2
		req := model.ContactRequest{
			Name:   "Ardi",
			NoTelp: "082828329292",
		}
		mockSuccess.On("Update", id, req).Return(errors.New("error"))

		res, _ := uc.Update(id, req)

		require.Equal(t, http.StatusBadGateway, res.Status)
		require.Equal(t, "Internal Database Error", res.Message)
	})

	t.Run("delete-contact", func(t *testing.T) {
		mockSuccess := mocks.NewRepoMock()
		uc := NewContactUseCase(mockSuccess)

		id := 2
		mockSuccess.On("Delete", id).Return(nil)

		res, err := uc.Delete(id)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
		require.Equal(t, "oke", res.Message)
	})

	t.Run("delete-contact-fail", func(t *testing.T) {
		mockSuccess := mocks.NewRepoMock()
		uc := NewContactUseCase(mockSuccess)

		id := 2
		mockSuccess.On("Delete", id).Return(errors.New("error"))

		res, _ := uc.Delete(id)

		require.Equal(t, http.StatusBadGateway, res.Status)
		require.Equal(t, "Internal Database Error", res.Message)
	})

	t.Run("delete-contact-fail by id", func(t *testing.T) {
		mockSuccess := mocks.NewRepoMock()
		uc := NewContactUseCase(mockSuccess)

		id := -2
		mockSuccess.On("Delete", id).Return(errors.New("error"))

		res, _ := uc.Delete(id)

		require.Equal(t, http.StatusNotFound, res.Status)
		require.Equal(t, "id tidak ditemukan", res.Message)
	})
}
