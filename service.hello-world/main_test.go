package main

import (
	"testing"
)

func TestSaysHelloWorld(t *testing.T) {
	result := salutation()
	expected := "Hello world!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
