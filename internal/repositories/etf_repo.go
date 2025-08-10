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

//create new
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

//get all
func (r *ETFRepository) GetAll() ([]models.ETF, error) {
	rows, err := r.db.Query(`SELECT id, name FROM etfs`)
	if err != nil{
		return nil, fmt.Errorf("get etfs: %w", err)
	}
	defer rows.Close()

	var etfs []models.ETF
	for rows.Next() {
		var etf models.ETF
		if err:= rows.Scan(&etf.ID, &etf.Name); err != nil {
			return nil, fmt.Errorf("scan etfs: %w", err)
		}

		//get holdings
		holdingRows, err := r.db.Query(`SELECT ticker, weight FROM etf_holdings WHERE etf_id = $1`, etf.ID)
		
		if err != nil {
			return nil, fmt.Errorf("get holdings: %w", err)
		}

		for holdingRows.Next() {
			var stock models.Stock
			if err := holdingRows.Scan(&stock.Ticker, &stock.Weight); err != nil {
				return nil, fmt.Errorf("scan holdings: %w", err)
			}
			etf.Stocks = append(etf.Stocks, stock)
		}
		holdingRows.Close()

		etfs = append(etfs, etf)

	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}	

	return etfs, nil
}


