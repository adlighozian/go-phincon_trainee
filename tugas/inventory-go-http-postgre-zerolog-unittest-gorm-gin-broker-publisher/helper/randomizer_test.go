package helper

import (
	"testing"
)

func TestRandomizer(t *testing.T) {
	t.Run("Randomizer", func(t *testing.T) {
		result := Randomizer()

		if result == "" {
			t.Error(result)
		}
	})
}
