package postgres

import (
	"context"

	repository "github.com/septemhill/txholder/repository"
	storage "github.com/septemhill/txholder/storage/user"
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

func (repo *userTxHolderRepository) Create(ctx context.Context, user *storage.User) error {
	return repo.base.Create(ctx, repo.config.tx, user)
}

func (repo *userTxHolderRepository) Delete(ctx context.Context, userId string) error {
	return repo.base.Delete(ctx, repo.config.tx, userId)
}

func (repo *userTxHolderRepository) Update(ctx context.Context, user *storage.User) error {
	return repo.base.Update(ctx, repo.config.tx, user)
}

func (repo *userTxHolderRepository) GetUser(ctx context.Context, userId string) (*storage.User, error) {
	return repo.base.GetUser(ctx, repo.config.tx, userId)
}

func (repo *userTxHolderRepository) ListUsers(ctx context.Context) ([]*storage.User, error) {
	return repo.base.ListUsers(ctx, repo.config.tx)
}

// Implementation of repository.TransactionHolderRepository

func (repo *userTxHolderRepository) Commit() error {
	return repo.config.tx.Commit()
}

func (repo *userTxHolderRepository) Rollback() error {
	return repo.config.tx.Rollback()
}

// Implementation of repository.TransactionalRepositoryCreator

func (repo *userTxHolderRepository) NewApplicationTxHolderRepository() repository.ApplicationTxHolderRepository {
	return newApplicationTxHolderRepository(WithTx(repo.config.tx))
}
func (repo *userTxHolderRepository) NewUserTxHolderRepository() repository.UserTxHolderRepository {
	return repo
}

var _ repository.UserTxHolderRepository = (*userTxHolderRepository)(nil)
