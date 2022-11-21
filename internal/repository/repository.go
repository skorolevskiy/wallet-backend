package repository

type Wallet interface {
}

type Transaction interface {
}

type Repository struct {
	Wallet
	Transaction
}

func NewRepository() *Repository {
	return &Repository{}
}
