package routes

import (
	"net/http"

	"github.com/vp2306/fund-forge/internal/handlers"
)

//register ETF routes
func RegisterEtfRoutes(mux *http.ServeMux, handler *handlers.ETFHandler) {
	mux.HandleFunc("/api/health", handlers.HealthHandler)

	mux.HandleFunc("/api/etfs", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPost:
			handler.CreateETF(w, r)
		case http.MethodGet:
			handler.GetAllETFs(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		
	})
}