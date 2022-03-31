package postgres

import (
	"txholder/repository"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository() *repositoryImpl {
	return &repositoryImpl{}
}

func (repo *repositoryImpl) NewApplicationRepository() repository.ApplicationRepository {
	return newApplicationRepository(WithDb(repo.db))
}

func (repo *repositoryImpl) NewUserRepository() repository.UserRepository {
	return newUserRepository(WithDb(repo.db))
}

func (repo *repositoryImpl) NewApplicationTxHolderRepository() repository.ApplicationTxHolderRepository {
	return newApplicationTxHolderRepository(WithTx(repo.db.Begin()))
}

func (repo *repositoryImpl) NewUserTxHolderRepository() repository.UserTxHolderRepository {
	return newUserTxHolderRepository(WithTx(repo.db.Begin()))
}

var _ repository.Repository = (*repositoryImpl)(nil)
