package repository

import (
	"context"
	"txholder/model"
)

//go:generate mockery --name=ApplicationRepository --inpackage --case underscore
type ApplicationRepository interface {
	CreateApplication(ctx context.Context, application *model.Application) error
	DeleteApplication(ctx context.Context, appId string) error
	GetApplication(ctx context.Context) (*model.Application, error)
	ListApplications(ctx context.Context) ([]*model.Application, error)
}

//go:generate mockery --name=ApplicationTxHolderRepository --inpackage --case underscore
type ApplicationTxHolderRepository interface {
	ApplicationRepository
	TransactionHolderRepository
	TransactionalRepositoryCreator
}
