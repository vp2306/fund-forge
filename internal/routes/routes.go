package routes

import (
	"net/http"

	"github.com/vp2306/fund-forge/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/health", handlers.HealthHandler)
}