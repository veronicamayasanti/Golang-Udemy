package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("maya")
	if result != "Hello maya" {
		panic("result is not 'Hello maya'")
	}
}
