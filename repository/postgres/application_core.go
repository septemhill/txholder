package postgres

import (
	"context"
	"txholder/model"
)

type applicationCoreRepository struct {
}

func (repo *applicationCoreRepository) CreateApplication(ctx context.Context, db sqlxDB, application *model.Application) error {
	return nil
}
func (repo *applicationCoreRepository) DeleteApplication(ctx context.Context, db sqlxDB, appId string) error {
	return nil
}

func (repo *applicationCoreRepository) GetApplication(ctx context.Context, db sqlxDB) (*model.Application, error) {
	return nil, nil
}

func (repo *applicationCoreRepository) ListApplications(ctx context.Context, db sqlxDB) ([]*model.Application, error) {
	return nil, nil
}

func newApplicationCoreRepository() *applicationCoreRepository {
	return &applicationCoreRepository{}
}
