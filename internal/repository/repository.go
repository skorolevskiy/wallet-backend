package repository

type Authorization interface {
}

type Wallet interface {
}

type Transaction interface {
}

type Repository struct {
	Authorization
	Wallet
	Transaction
}

func NewRepository() *Repository {
	return &Repository{}
}
