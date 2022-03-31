package postgres

import (
	"context"
	"txholder/model"
	"txholder/repository"
)

type userTxHolderRepository struct {
	config *postgresConfig
	base   *userCoreRepository
}

func newUserTxHolderRepository(opts ...repository.Option) *userTxHolderRepository {
	repo := &userTxHolderRepository{
		config: &postgresConfig{},
		base:   newUserCoreRepository(),
	}

	for _, opt := range opts {
		opt.Apply(repo.config)
	}

	return repo
}

// Implemtnation of repository.ApplicationRepository

func (repo *userTxHolderRepository) CreateUser(ctx context.Context, application *model.User) error {
	return repo.base.CreateUser(ctx, repo.config.tx, application)
}

func (repo *userTxHolderRepository) DeleteUser(ctx context.Context, userId string) error {
	return repo.base.DeleteUser(ctx, repo.config.tx, userId)
}

func (repo *userTxHolderRepository) GetUser(ctx context.Context, userId string) (*model.User, error) {
	return repo.base.GetUser(ctx, repo.config.tx, userId)
}

func (repo *userTxHolderRepository) ListUsers(ctx context.Context) ([]*model.User, error) {
	return repo.base.ListUsers(ctx, repo.config.tx)
}

// Implementation of repository.TransactionHolderRepository

func (repo *userTxHolderRepository) Commit() error {
	return repo.config.tx.Commit().Error
}

func (repo *userTxHolderRepository) Rollback() error {
	return repo.config.tx.Rollback().Error
}

// Implementation of repository.TransactionalRepositoryCreator

func (repo *userTxHolderRepository) NewApplicationTxHolderRepository() repository.ApplicationTxHolderRepository {
	return newApplicationTxHolderRepository(WithTx(repo.config.tx))
}
func (repo *userTxHolderRepository) NewUserTxHolderRepository() repository.UserTxHolderRepository {
	return repo
}

var _ repository.UserTxHolderRepository = (*userTxHolderRepository)(nil)
