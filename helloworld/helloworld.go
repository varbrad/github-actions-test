package helloworld

import (
	"strings"
)

// HelloWorld returns a Hello World string!
func HelloWorld() string {
	return "Hello, world!"
}

// ShoutHelloWorld returns a HELLO WORLD string!
func ShoutHelloWorld() string {
	return strings.ToUpper(HelloWorld())
}
