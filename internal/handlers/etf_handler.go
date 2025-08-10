package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vp2306/fund-forge/internal/models"
	"github.com/vp2306/fund-forge/internal/services"
)

type ETFHandler struct {
	service *services.ETFService
}

func NewETFHandler(service *services.ETFService) *ETFHandler {
	return &ETFHandler{service: service}
}

//handle post
func (h *ETFHandler) CreateETF(w http.ResponseWriter, r *http.Request){
	var etf models.ETF

	//decode json body
	if err := json.NewDecoder(r.Body).Decode(&etf); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//call service layer to create etf
	createdETF, err := h.service.CreateETF(etf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//respond with created etf
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdETF)

}

func (h *ETFHandler) GetAllETFs(w http.ResponseWriter, r *http.Request) {
	
	allETFs, err := h.service.GetAllETFs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allETFs)
}