package main

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

// A small addition just to trigger tests...
func TestSaysHelloWorld(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", bytes.NewReader([]byte{}))
	handler(w, r)
	result := w.Body.String()
	expected := "¡Hola mundo!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
