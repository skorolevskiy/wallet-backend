package service

import (
	"github.com/skorolevskiy/wallet-backend/internal/domain"
	"github.com/skorolevskiy/wallet-backend/internal/repository"
)

type WalletService struct {
	repo repository.Wallet
}

func NewWalletService(repo repository.Wallet) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) CreateWallet(userId int, wallet domain.Wallet) (int, error) {
	wallet.Balance = 0
	return s.repo.CreateWallet(userId, wallet)
}
