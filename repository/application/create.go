package application

import (
	"context"

	storage "github.com/septemhill/txholder/storage/application"
)

func (repo *applicationRepository) Create(ctx context.Context, app *storage.Application) error {
	return nil
}
