package postgres

import (
	"context"

	storage "github.com/septemhill/txholder/storage/user"
)

type userCoreRepository struct {
}

func (repo *userCoreRepository) Create(ctx context.Context, db sqlxDB, application *storage.User) error {
	return nil
}
func (repo *userCoreRepository) Delete(ctx context.Context, db sqlxDB, userId string) error {
	return nil
}

func (repo *userCoreRepository) Update(ctx context.Context, db sqlxDB, application *storage.User) error {
	return nil
}

func (repo *userCoreRepository) GetUser(ctx context.Context, db sqlxDB, userId string) (*storage.User, error) {
	return nil, nil
}

func (repo *userCoreRepository) ListUsers(ctx context.Context, db sqlxDB) ([]*storage.User, error) {
	return nil, nil
}

func newUserCoreRepository() *userCoreRepository {
	return &userCoreRepository{}
}
