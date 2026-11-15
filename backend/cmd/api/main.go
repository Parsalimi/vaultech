package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe("localhost:8000", newRouter()))
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthHandler)
	return mux
}

type healthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rsp := healthResponse{Status: "ok"}

	if err := json.NewEncoder(w).Encode(rsp); err != nil {
		log.Print(err)
	}
}
