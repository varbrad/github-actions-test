package helloworld

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHelloWorld(t *testing.T) {

	t.Run("say hello!", func(t *testing.T) {
		assert.Equal(t, "Hello, world!", HelloWorld())
	})
}

func TestShoutHelloWorld(t *testing.T) {
	assert.Equal(t, "HELLO, WORLD!", ShoutHelloWorld())
}

func ExampleHelloWorld() {
	fmt.Println(HelloWorld())
	// Output: Hello, world!
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld()
	}
}

func BenchmarkShoutHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShoutHelloWorld()
	}
}
