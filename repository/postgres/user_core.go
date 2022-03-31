package postgres

import (
	"context"
	"txholder/model"
)

type userCoreRepository struct {
}

func (repo *userCoreRepository) CreateUser(ctx context.Context, db sqlxDB, application *model.User) error {
	return nil
}
func (repo *userCoreRepository) DeleteUser(ctx context.Context, db sqlxDB, userId string) error {
	return nil
}

func (repo *userCoreRepository) GetUser(ctx context.Context, db sqlxDB, userId string) (*model.User, error) {
	return nil, nil
}

func (repo *userCoreRepository) ListUsers(ctx context.Context, db sqlxDB) ([]*model.User, error) {
	return nil, nil
}

func newUserCoreRepository() *userCoreRepository {
	return &userCoreRepository{}
}
