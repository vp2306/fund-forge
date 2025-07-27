package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Message string `json:"message"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HealthResponse{Message: "server is healthy"})
}