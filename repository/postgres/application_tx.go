package postgres

import (
	"context"
	"txholder/model"
	"txholder/repository"
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

func (repo *applicationTxHolderRepository) CreateApplication(ctx context.Context, application *model.Application) error {
	return repo.base.CreateApplication(ctx, repo.config.tx, application)
}

func (repo *applicationTxHolderRepository) DeleteApplication(ctx context.Context, appId string) error {
	return repo.base.DeleteApplication(ctx, repo.config.tx, appId)
}

func (repo *applicationTxHolderRepository) GetApplication(ctx context.Context) (*model.Application, error) {
	return repo.base.GetApplication(ctx, repo.config.tx)
}

func (repo *applicationTxHolderRepository) ListApplications(ctx context.Context) ([]*model.Application, error) {
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
