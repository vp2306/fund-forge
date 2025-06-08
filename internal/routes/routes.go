package routes

import (
	"net/http"

	"github.com/vp2306/fund-forge/internal/handlers"
)

func Routes(mux *http.ServeMux, h *handlers.Handler){
	HealthRoute(mux, h)
}