package handlers

import (
	"net/http"

	"github.com/vp2306/fund-forge/internal/services"
)

type TransactionHandler struct {
	service *services.TradeService
}

func NewTransactionHandler(service *services.TradeService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (s *TransactionHandler) BuyETF(w *http.ResponseWriter, r http.Request)  {

}