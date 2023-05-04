package client

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	t.Run("test config", func(t *testing.T) {
		config := GetDB("mysql").GetMysqlConnection()
		require.NotEmpty(t, config)
	})
}