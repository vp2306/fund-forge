package models

import "time"

// buy/sell order
type Trade struct {
    ID          int64      `json:"id"`
    EtfID       int64      `json:"etf_id"`
    TradeType   string     `json:"trade_type"`
    Amount      float64    `json:"amount"`
    Status      string     `json:"status"`
    CreatedAt   time.Time  `json:"created_at"`
    ExecutedAt  *time.Time `json:"executed_at,omitempty"`
    
    // trade can have multiple executions
    Executions  []TradeExecution `json:"executions,omitempty"`
}

//purchase/sell within the trade
type TradeExecution struct {
	ID              int64      `json:"id"`
    TradeID         int64      `json:"trade_id"`
    Ticker          string     `json:"ticker"`
    Shares          float64    `json:"shares"`
    PricePerShare   float64    `json:"price_per_share"`
    TotalCost       float64    `json:"total_cost"`
}

//current holdings
type Position struct {
    ID              int64   `json:"id"`
    Ticker          string  `json:"ticker"`
    TotalShares     float64 `json:"total_shares"`
    AverageCost     float64 `json:"average_cost"`
}

type BuyRequest struct {
    Amount float64 `json:"amount"`
}

type SellRequest struct {
    Amount float64 `json:"amount"`
}

