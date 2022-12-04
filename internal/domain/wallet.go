package domain

import "time"

type Wallet struct {
	ID         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name" binding:"required"`
	Balance    float64   `json:"balance" db:"balance"`
	Currency   string    `json:"currency" db:"currency"`
	RegisterAt time.Time `db:"register_at"`
}

type UserWallets struct {
	id       int
	UserId   int
	WalletId int
}
