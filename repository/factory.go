package repository

type RepositoryFactory interface {
	NewRepository() Repository
	NewTxRepository() TxRepository
}
