package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, "Hello, world!", HelloWorld())
}
