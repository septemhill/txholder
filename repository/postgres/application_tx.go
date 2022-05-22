package postgres

import (
	"context"

	repository "github.com/septemhill/txholder/repository"
	storage "github.com/septemhill/txholder/storage/application"
)

type applicationTxHolderRepository struct {
	config *postgresConfig
	base   *applicationCoreRepository
}

func newApplicationTxHolderRepository(opts ...repository.Option) *applicationTxHolderRepository {
	repo := &applicationTxHolderRepository{
		config: &postgresConfig{},
		base:   newApplicationCoreRepository(),
	}

	for _, opt := range opts {
		opt.Apply(repo.config)
	}

	return repo
}

// Implemtnation of repository.ApplicationRepository

func (repo *applicationTxHolderRepository) Create(ctx context.Context, application *storage.Application) error {
	return repo.base.Create(ctx, repo.config.tx, application)
}

func (repo *applicationTxHolderRepository) Delete(ctx context.Context, appId string) error {
	return repo.base.Delete(ctx, repo.config.tx, appId)
}

func (repo *applicationTxHolderRepository) Update(ctx context.Context, application *storage.Application) error {
	return repo.base.Update(ctx, repo.config.tx, application)
}

func (repo *applicationTxHolderRepository) GetApplication(ctx context.Context) (*storage.Application, error) {
	return repo.base.GetApplication(ctx, repo.config.tx)
}

func (repo *applicationTxHolderRepository) ListApplications(ctx context.Context) ([]*storage.Application, error) {
	return repo.base.ListApplications(ctx, repo.config.tx)
}

// Implementation of repository.TransactionHolderRepository

func (repo *applicationTxHolderRepository) Commit() error {
	return repo.config.tx.Commit()
}

func (repo *applicationTxHolderRepository) Rollback() error {
	return repo.config.tx.Rollback()
}

// Implementation of repository.TransactionalRepositoryCreator

func (repo *applicationTxHolderRepository) NewApplicationTxHolderRepository() repository.ApplicationTxHolderRepository {
	return repo
}

func (repo *applicationTxHolderRepository) NewUserTxHolderRepository() repository.UserTxHolderRepository {
	return newUserTxHolderRepository(WithTx(repo.config.tx))
}

var _ repository.ApplicationTxHolderRepository = (*applicationTxHolderRepository)(nil)
