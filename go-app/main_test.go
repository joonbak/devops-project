package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	handler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resp Response
	err := json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Errorf("Could not parse response %v", err)
	}

	expected := "Hello from Go API Server!"
	if resp.Message != expected {
		t.Errorf("Unexpected message: got %v want %v", resp.Message, expected)
	}
}
