package main

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestSaysHelloWorld(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", bytes.NewReader([]byte{}))
	handler(w, r)
	result := w.Body.String()
	expected := "Ciao mondo!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
