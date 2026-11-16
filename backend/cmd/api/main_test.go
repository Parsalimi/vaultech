package main

import (
	"encoding/json"
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
	// healthHandler(rec, req)

	mux := newRouter()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestContentTypeHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	mux := newRouter()
	mux.ServeHTTP(rec, req)

	contentType := rec.Result().Header.Get("Content-Type")

	if contentType != "application/json" {
		t.Errorf("Got Content-Type %q, want %q", contentType, "application/json")
	}
}

func TestVerifyStatusHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	mux := newRouter()
	mux.ServeHTTP(rec, req)

	var body healthResponse
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	if body.Status != "ok" {
		t.Errorf("Got JSON status %q, want %q", body.Status, "ok")
	}
}

func TestHealthMethodAllowed(t *testing.T) {
	tests := []struct {
		name   string
		method string
		want   int
	}{
		{
			name:   "POST",
			method: http.MethodPost,
			want:   http.StatusMethodNotAllowed,
		},
		{
			name:   "DELETE",
			method: http.MethodDelete,
			want:   http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/healthz", nil)
			rec := httptest.NewRecorder()

			mux := newRouter()
			mux.ServeHTTP(rec, req)

			if rec.Code != tt.want {
				t.Errorf("got status %d, want %d", rec.Code, tt.want)
			}
		})
	}
}

func TestUnknownHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	rec := httptest.NewRecorder()

	mux := newRouter()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("got status %d, want %d", rec.Code, http.StatusNotFound)
	}
}
