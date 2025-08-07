package services

import (
	"errors"

	"github.com/vp2306/fund-forge/internal/models"
	"github.com/vp2306/fund-forge/internal/repositories"
)

type ETFService struct {
	repo *repositories.ETFRepository
}

func NewETfService(repo *repositories.ETFRepository) *ETFService {
	return &ETFService{repo: repo}
}

//validate input for createETF and call the repository 
func (s *ETFService) CreateETF(etf models.ETF) (models.ETF, error) {
	if etf.Name == "" {
		return models.ETF{}, errors.New("ETF Name is required")
	}

	if len(etf.Stocks) == 0 {
		return models.ETF{}, errors.New("ETF must contain at least one stock")
	}

	var weight float64
	for _, stock := range etf.Stocks {
		weight += stock.Weight
	}
	if weight != 1 {
		return models.ETF{}, errors.New("Total weight of all holdings must be 1")
	}

	return s.repo.Create(etf)
	
}