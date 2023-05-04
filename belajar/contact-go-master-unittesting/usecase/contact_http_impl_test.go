package usecase

import (
	"net/http"
	"testing"

	"contact-go/mocks"
	"contact-go/model"

	"github.com/stretchr/testify/require"
)

func TestContactHTTP(t *testing.T) {
	t.Run("get-list", func(t *testing.T) {
		mockSuccess := mocks.NewContactRepoMock()
		uc := NewUseCase(mockSuccess)

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
		
		res, err := uc.List()
		// test
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
		require.Equal(t, "Ok", res.Message)
	})

	t.Run("add-contact", func(t *testing.T) {
		mockSuccess := mocks.NewContactRepoMock()
		uc := NewUseCase(mockSuccess)

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

		res, err := uc.Add(req)

		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, res.Status)
		require.Equal(t, "Created", res.Message)
	})

	t.Run("update-contact", func(t *testing.T) {
		mockSuccess := mocks.NewContactRepoMock()
		uc := NewUseCase(mockSuccess)

		id := 2
		req := model.ContactRequest{
			Name: "Ardi",
			NoTelp: "082828329292",
		}
		mockSuccess.On("Update", id, req).Return(nil)

		idStr := "2"
		res, err := uc.Update(idStr, req)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
		require.Equal(t, "Updated", res.Message)
	})

	t.Run("delete-contact", func(t *testing.T) {
		mockSuccess := mocks.NewContactRepoMock()
		uc := NewUseCase(mockSuccess)

		id := 2
		mockSuccess.On("Delete", id).Return(nil)

		idStr := "2"
		res, err := uc.Delete(idStr)
		
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
		require.Equal(t, "Deleted", res.Message)
	})
}