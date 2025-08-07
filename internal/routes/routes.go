package routes

import (
	"net/http"

	"github.com/vp2306/fund-forge/internal/handlers"
)

//register ETF routes
func RegisterEtfRoutes(mux *http.ServeMux, handler *handlers.ETFHandler) {
	mux.HandleFunc("/api/health", handlers.HealthHandler)

	mux.HandleFunc("/api/etfs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.CreateETF(w, r)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})
}