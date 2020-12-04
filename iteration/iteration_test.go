package iteration

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestRepeat(t *testing.T) {
	t.Run("can repeat just once", func(t *testing.T) {
		expected := "a"
		actual := Repeat("a", 1)

		assert.Equal(t, expected, actual)
	})

	t.Run("can repeat many times", func(t *testing.T) {
		expected := "abcabcabc"
		actual := Repeat("abc", 3)

		assert.Equal(t, expected, actual)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("abcdef", 100)
	}
}
