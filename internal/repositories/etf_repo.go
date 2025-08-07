package repositories

import (
	"database/sql"
	"fmt"

	"github.com/vp2306/fund-forge/internal/models"
)

type ETFRepository struct {
	db *sql.DB
}

func NewETFRepository (db *sql.DB) *ETFRepository {
	return &ETFRepository{db: db}
}

func (r *ETFRepository) Create(etf models.ETF) (models.ETF, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return models.ETF{}, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	//insert etf
	query := `INSERT INTO etfs (name) VALUES ($1) RETURNING id`
	if err := tx.QueryRow(query, etf.Name).Scan(&etf.ID); err != nil {
		return models.ETF{}, fmt.Errorf("insert etf: %w", err)
	}

	//insert etf holdings
	holdingQuery := `INSERT into etf_holdings (etf_id, ticker, weight) VALUES ($1, $2, $3)`
	for _, stock := range etf.Stocks {
		_, err := tx.Exec(holdingQuery, etf.ID, stock.Ticker, stock.Weight)
		if err != nil {
			return models.ETF{}, fmt.Errorf("insert holding: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return models.ETF{}, fmt.Errorf("commit tx: %w", err)
	}

	return etf, nil
}


