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

func (s *WalletService) GetAllWallets(userId int) ([]domain.Wallet, error) {
	return s.repo.GetAllWallets(userId)
}

func (s *WalletService) GetWalletById(userId, walletId int) (domain.Wallet, error) {
	return s.repo.GetWalletById(userId, walletId)
}
