package helloworld

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, "Hello, world!", HelloWorld())
}

func TestShoutHelloWorld(t *testing.T) {
	assert.Equal(t, "HELLO, WORLD!", ShoutHelloWorld())
}
