package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

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
func (h *ETFHandler) CreateETF(w http.ResponseWriter, r *http.Request) {
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
	if err := json.NewEncoder(w).Encode(createdETF); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}

}

func (h *ETFHandler) GetAllETFs(w http.ResponseWriter, r *http.Request) {

	allETFs, err := h.service.GetAllETFs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allETFs); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *ETFHandler) GetETFByID(w http.ResponseWriter, r *http.Request) {
	// Parse the "id" query parameter
	ids := r.URL.Query().Get("id")
	if ids == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	etf, err := h.service.GetETFByID(id)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			http.Error(w, "ETF not found", http.StatusNotFound)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(etf); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *ETFHandler) DeleteETF(w http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query().Get("id")
	if ids == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	err = h.service.DeleteETFByID(id)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			http.Error(w, "ETF not found", http.StatusNotFound)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ETFHandler) UpdateETF(w http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query().Get("id")
	if ids == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	var etf models.ETF
	if err := json.NewDecoder(r.Body).Decode(&etf); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	etf.ID = id // Ensure the ID from the URL is used

	if err := h.service.UpdateETF(etf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
