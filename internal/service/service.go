package service

import (
	"github.com/skorolevskiy/wallet-backend/internal/domain"
	"github.com/skorolevskiy/wallet-backend/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Wallet interface {
	CreateWallet(userId int, wallet domain.Wallet) (int, error)
	GetAllWallets(userId int) ([]domain.Wallet, error)
	GetWalletById(userId, walletId int) (domain.Wallet, error)
}

type Transaction interface {
}

type Service struct {
	Authorization
	Wallet
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Wallet:        NewWalletService(repos.Wallet),
	}
}
