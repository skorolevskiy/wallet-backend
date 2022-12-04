package domain

import "time"

type Transaction struct {
	Id               int       `json:"id"`
	TypeTransaction  string    `json:"typeTransaction"`
	Amount           float64   `json:"amount"`
	BalanceAfter     float64   `json:"balanceAfter"`
	CommissionAmount float64   `json:"commissionAmount"`
	Currency         string    `json:"currency"`
	CreatedAt        time.Time `json:"createdAt"`
}

type WalletTransactions struct {
	Id            int
	WalletID      int
	TransactionId int
}
