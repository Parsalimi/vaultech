package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	// creates an HTTP requests in Memory, it doesn't send anything over the network.
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)

	// Creates a fake ResponseWriter that records your handler writes
	rec := httptest.NewRecorder()

	// Executes your handler using the fake request and recorder.
	healthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", rec.Code, http.StatusOK)
	}
}
