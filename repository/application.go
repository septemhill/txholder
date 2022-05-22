package repository

import (
	"context"

	storage "github.com/septemhill/txholder/storage/application"
)

// type ApplicationRepository interface {
// 	CreateApplication(ctx context.Context, application *model.Application) error
// 	DeleteApplication(ctx context.Context, appId string) error
// 	GetApplication(ctx context.Context) (*model.Application, error)
// 	ListApplications(ctx context.Context) ([]*model.Application, error)
// }

type ApplicationRepository interface {
	Create(ctx context.Context, app *storage.Application) error
	Delete(ctx context.Context, appId string) error
	Update(ctx context.Context, app *storage.Application) error
	GetApplication(ctx context.Context) (*storage.Application, error)
	ListApplications(ctx context.Context) ([]*storage.Application, error)
}

type ApplicationTxHolderRepository interface {
	ApplicationRepository
	TransactionHolderRepository
	TransactionalRepositoryCreator
}
