package domain

import "time"

type Wallet struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}
