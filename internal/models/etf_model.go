package models

type Stock struct {
	Ticker string  `json:"ticker"`
	Weight float64 `json:"weight"`
}

type ETF struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Stocks []Stock `json:"stocks"`
}
