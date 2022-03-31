package postgres

import (
	"context"
	"txholder/model"
	"txholder/repository"
)

type applicationRepository struct {
	config *postgresConfig
	base   *applicationCoreRepository
}

func newApplicationRepository(opts ...repository.Option) *applicationRepository {
	repo := &applicationRepository{
		config: &postgresConfig{},
		base:   newApplicationCoreRepository(),
	}

	for _, opt := range opts {
		opt.Apply(repo.config)
	}

	return repo
}

func (repo *applicationRepository) CreateApplication(ctx context.Context, application *model.Application) error {
	return repo.base.CreateApplication(ctx, repo.config.db, application)
}

func (repo *applicationRepository) DeleteApplication(ctx context.Context, appId string) error {
	return repo.base.DeleteApplication(ctx, repo.config.db, appId)
}

func (repo *applicationRepository) GetApplication(ctx context.Context) (*model.Application, error) {
	return repo.base.GetApplication(ctx, repo.config.db)
}

func (repo *applicationRepository) ListApplications(ctx context.Context) ([]*model.Application, error) {
	return repo.base.ListApplications(ctx, repo.config.db)
}

var _ repository.ApplicationRepository = (*applicationRepository)(nil)
