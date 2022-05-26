package repository

import "gorm.io/gorm"

type repositoryFactory struct {
	db *gorm.DB
}

func (factory *repositoryFactory) NewRepository() Repository {
	return NewRepository(factory.db)
}

func (factory *repositoryFactory) NewTxRepository() TxRepository {
	return nil
}

func NewRepositoryFactory() *repositoryFactory {
	return &repositoryFactory{}
}

var _ RepositoryFactory = (*repositoryFactory)(nil)
