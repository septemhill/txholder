package postgres

import (
	"context"

	storage "github.com/septemhill/txholder/storage/application"
)

type applicationCoreRepository struct {
}

func (repo *applicationCoreRepository) Create(ctx context.Context, db sqlxDB, application *storage.Application) error {
	return nil
}
func (repo *applicationCoreRepository) Delete(ctx context.Context, db sqlxDB, appId string) error {
	return nil
}

func (repo *applicationCoreRepository) Update(ctx context.Context, db sqlxDB, application *storage.Application) error {
	return nil
}

func (repo *applicationCoreRepository) GetApplication(ctx context.Context, db sqlxDB) (*storage.Application, error) {
	return nil, nil
}

func (repo *applicationCoreRepository) ListApplications(ctx context.Context, db sqlxDB) ([]*storage.Application, error) {
	return nil, nil
}

func newApplicationCoreRepository() *applicationCoreRepository {
	return &applicationCoreRepository{}
}
