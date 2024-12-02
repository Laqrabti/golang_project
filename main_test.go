package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld()
	expected := "hello, go"

	if result != expected {
		t.Errorf("Expected '%v', but got '%v'", expected, result)
	}

}
