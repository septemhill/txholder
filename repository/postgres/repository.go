package postgres

import (
	"txholder/repository"

	"github.com/jmoiron/sqlx"
)

type repositoryImpl struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func NewRepository() *repositoryImpl {
	return &repositoryImpl{}
}

func (repo *repositoryImpl) NewApplicationRepository() repository.ApplicationRepository {
	return newApplicationRepository(WithDb(repo.db))
}

func (repo *repositoryImpl) NewUserRepository() repository.UserRepository {
	return nil
}

func (repo *repositoryImpl) NewApplicationTxHolderRepository() repository.ApplicationTxHolderRepository {
	return nil
}

func (repo *repositoryImpl) NewUserTxHolderRepository() repository.UserTxHolderRepository {
	return nil
}

var _ repository.Repository = (*repositoryImpl)(nil)
