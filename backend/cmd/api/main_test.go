package main

import (
	"net/http"
	"testing"
)

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthHandler)
	return mux
}

func TestHealthzHandler(t *testing.T) {

}
