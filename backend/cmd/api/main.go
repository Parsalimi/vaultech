package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /healthz", healthHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
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
