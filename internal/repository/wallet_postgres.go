package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/skorolevskiy/wallet-backend/internal/domain"
)

type WalletPostgres struct {
	db *sqlx.DB
}

func NewWalletPostgres(db *sqlx.DB) *WalletPostgres {
	return &WalletPostgres{db: db}
}

func (r *WalletPostgres) CreateWallet(userId int, wallet domain.Wallet) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createWallerQuery := fmt.Sprintf("INSERT INTO %s (name, balance, currency) VALUES ($1, $2, $3)", walletsTable)
	row := tx.QueryRow(createWallerQuery, wallet.Name, wallet.Balance, wallet.Currency)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersWalletQuery := fmt.Sprintf("INSERT INTO %s (user_id, wallet_id) VALUES ($1, $2)", usersWalletTable)
	_, err = tx.Exec(createUsersWalletQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
