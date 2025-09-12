package services

import (
	"errors"
	"math"

	"github.com/vp2306/fund-forge/internal/models"
	"github.com/vp2306/fund-forge/internal/repositories"
)

// ErrNotFound is returned when a requested resource does not exist.
var ErrNotFound = repositories.ErrNotFound

type ETFService struct {
	repo *repositories.ETFRepository
}

func NewETFService(repo *repositories.ETFRepository) *ETFService {
	return &ETFService{repo: repo}
}

// validate input for CreateETF and call the repository
func (s *ETFService) CreateETF(etf models.ETF) (models.ETF, error) {
	if etf.Name == "" {
		return models.ETF{}, errors.New("ETF Name is required")
	}

	if len(etf.Stocks) == 0 {
		return models.ETF{}, errors.New("ETF must contain at least one stock")
	}

	var total float64
	for _, stock := range etf.Stocks {
		if stock.Weight < 0 {
			return models.ETF{}, errors.New("holding weight cannot be negative")
		}
		total += stock.Weight
	}
	if math.Abs(total-1.0) > 1e-9 {
		return models.ETF{}, errors.New("total weight of all holdings must be 1")
	}

	return s.repo.Create(etf)
}

func (s *ETFService) GetAllETFs() ([]models.ETF, error) {
	return s.repo.GetAll()
}

func (s *ETFService) GetETFByID(id int64) (models.ETF, error) {
	return s.repo.GetByID(id)
}

func (s *ETFService) DeleteETFByID(id int64) error {
	return s.repo.DeleteByID(id)
}

func (s *ETFService) UpdateETF(etf models.ETF) error {
	return s.repo.Update(etf)
}
