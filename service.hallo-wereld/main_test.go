package main

import (
	"testing"
)

func TestSaysHelloWorld(t *testing.T) {
	result := salutation()
	expected := "Hallo werld!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
