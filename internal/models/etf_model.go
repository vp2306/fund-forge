package models

type ETF struct {
	ID       int64    `json:"id"`
	Name     string   `json:"string"`
	Holdings []string `json:"holdings"`
}