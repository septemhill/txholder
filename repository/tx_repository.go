package repository

type TxRepository interface {
	DomainRepositoryFactory
	Transactional
}
