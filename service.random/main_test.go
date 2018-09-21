package main

import (
	"testing"
)

func TestSaysHelloWorld(t *testing.T) {
	result := salutation()
	expected := "Baaaaaaah!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
