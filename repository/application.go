package repository

import (
	"context"
	"txholder/model"
)

type ApplicationRepository interface {
	CreateApplication(ctx context.Context, application *model.Application) error
	DeleteApplication(ctx context.Context, appId string) error
	GetApplication(ctx context.Context) (*model.Application, error)
	ListApplications(ctx context.Context) ([]*model.Application, error)
}

type ApplicationTxHolderRepository interface {
	ApplicationRepository
	TransactionHolderRepository
	TransactionalRepositoryCreator
}
