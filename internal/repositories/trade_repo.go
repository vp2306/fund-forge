package repositories

import (
	"database/sql"
)

type TradeRepository struct {
	db *sql.DB
}