package application

import (
	"context"

	storage "github.com/septemhill/txholder/storage/application"
)

func (repo *applicationRepository) GetApplication(ctx context.Context, appId string) (*storage.Application, error) {
	return nil, nil
}
