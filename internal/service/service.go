package service

import "github.com/skorolevskiy/wallet-backend/internal/repository"

type Wallet interface {
}

type Transaction interface {
}

type Service struct {
	Wallet
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
