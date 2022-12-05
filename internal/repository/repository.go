package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/skorolevskiy/wallet-backend/internal/domain"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type Wallet interface {
	CreateWallet(userId int, list domain.Wallet) (int, error)
	GetAllWallets(userId int) ([]domain.Wallet, error)
	GetWalletById(userId, walletId int) (domain.Wallet, error)
}

type Transaction interface {
}

type Repository struct {
	Authorization
	Wallet
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Wallet:        NewWalletPostgres(db),
	}
}
