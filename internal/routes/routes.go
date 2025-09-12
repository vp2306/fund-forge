package routes

import (
	"net/http"

	"github.com/vp2306/fund-forge/internal/handlers"
)

// RegisterETFRoutes registers ETF routes
func RegisterETFRoutes(mux *http.ServeMux, handler *handlers.ETFHandler) {
	mux.HandleFunc("/api/health", handlers.HealthHandler)

	mux.HandleFunc("/api/etfs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateETF(w, r)
		case http.MethodGet:
			idStr := r.URL.Query().Get("id")
			if idStr != "" {
				handler.GetETFByID(w, r)
			} else {
				handler.GetAllETFs(w, r)
			}
		case http.MethodDelete:
			handler.DeleteETF(w, r)
		case http.MethodPut, http.MethodPatch:
			handler.UpdateETF(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
