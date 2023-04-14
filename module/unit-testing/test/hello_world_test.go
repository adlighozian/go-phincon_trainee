package helper

import (
	"fmt"
	"testing"
	"time"
	"unit-testing/model"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("start test")
	start := time.Now()
	m.Run()
	end := time.Since(start)
	fmt.Println("end test")

	fmt.Println("duration", end)
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("adli")

	if result != "hello adli" {
		panic("bukan hello adli")
	}

}

func TestSapaNama(t *testing.T) {

	t.Run("aman", func(t *testing.T) {
		result := SapaNama("adli")
		require.Equal(t, true, result, "input bukan adli")
	})

	t.Run("gagal", func(t *testing.T) {
		result := SapaNama("adlis")
		require.Equal(t, false, result, "input harus yang salah")
	})
}

func TestGetNama(t *testing.T) {

	for _, v := range model.Names {
		t.Run(v.SubTest, func(t *testing.T) {
			result := SapaNama(v.Input)
			require.Equal(t, v.Status, result, v.Message)
		})
	}

}
