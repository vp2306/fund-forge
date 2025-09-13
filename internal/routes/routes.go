package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/vp2306/fund-forge/internal/handlers"
)

// RegisterETFRoutes registers ETF routes
func RegisterETFRoutes(r chi.Router, handler *handlers.ETFHandler) {

	r.Get("/api/health", handlers.Health)
    r.Route("/api/etfs", func(r chi.Router) {
        r.Post("/", handler.CreateETF)
        r.Get("/", handler.GetAllETFs)
        r.Get("/{id}", handler.GetETFByID)
        r.Put("/{id}", handler.UpdateETF)
        r.Patch("/{id}", handler.UpdateETF)
        r.Delete("/{id}", handler.DeleteETF)
    })
}
