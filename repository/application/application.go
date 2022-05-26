package application

import (
	"context"

	storage "github.com/septemhill/txholder/storage/application"
	"gorm.io/gorm"
)

type ApplicationRepository interface {
	Create(ctx context.Context, app *storage.Application) error
	Update(ctx context.Context, app *storage.Application) error
	GetApplication(ctx context.Context, appId string) (*storage.Application, error)
}

type applicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) *applicationRepository {
	return &applicationRepository{
		db: db,
	}
}

var _ ApplicationRepository = (*applicationRepository)(nil)
