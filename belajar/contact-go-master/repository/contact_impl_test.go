package repository

import (
	"testing"

	"contact-go/model"
	"github.com/stretchr/testify/require"
)

func TestRepo(t *testing.T) {
	t.Run("test get list contact", func(t *testing.T) {
		repo := NewContactRepository()

		res, err := repo.List()
		require.NoError(t, err)
		require.Equal(t, []model.Contact{}, res)
	})

	t.Run("test create new contact", func(t *testing.T) {
		repo := NewContactRepository()

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

		res, err := repo.Add(req)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	t.Run("test update contact", func(t *testing.T) {
		repo := NewContactRepository()

		id := 2
		req := model.ContactRequest{
			Name: "Ardi",
			NoTelp: "082828329292",
		}

		err := repo.Update(id, req)
		require.NoError(t, err)
	})

	t.Run("test delete contact", func(t *testing.T) {
		repo := NewContactRepository()

		id := 2
		err := repo.Delete(id)
		require.NoError(t, err)
	})
}