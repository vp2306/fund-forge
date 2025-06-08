package routes

import (
	"net/http"

	"github.com/vp2306/fund-forge/internal/handlers"
)

func HealthRoute(mux *http.ServeMux, h *handlers.Handler) {
	mux.HandleFunc("/health", h.HealthHandler())
}