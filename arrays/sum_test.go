package arrays

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [...]int{1, 2, 3, 4, 5}

		expected := 15
		actual := Sum(numbers[:])

		assert.Equal(t, expected, actual)
	})
}
