package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable              = "users"
	walletsTable            = "wallets"
	usersWalletTable        = "users_wallets"
	transactionTable        = "transactions"
	walletsTransactionTable = "wallets_transactions"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	return db, nil
}
