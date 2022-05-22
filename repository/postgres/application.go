package postgres

import (
	"context"

	repository "github.com/septemhill/txholder/repository"
	storage "github.com/septemhill/txholder/storage/application"
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

func (repo *applicationRepository) Create(ctx context.Context, application *storage.Application) error {
	return repo.base.Create(ctx, repo.config.db, application)
}

func (repo *applicationRepository) Delete(ctx context.Context, appId string) error {
	return repo.base.Delete(ctx, repo.config.db, appId)
}

func (repo *applicationRepository) Update(ctx context.Context, app *storage.Application) error {
	return nil
}

func (repo *applicationRepository) GetApplication(ctx context.Context) (*storage.Application, error) {
	return repo.base.GetApplication(ctx, repo.config.db)
}

func (repo *applicationRepository) ListApplications(ctx context.Context) ([]*storage.Application, error) {
	return repo.base.ListApplications(ctx, repo.config.db)
}

var _ repository.ApplicationRepository = (*applicationRepository)(nil)
