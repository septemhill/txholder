package postgres

import (
	"context"
	"txholder/model"
	"txholder/repository"
)

type userRepository struct {
	config *postgresConfig
	base   *userCoreRepository
}

func newUserRepository(opts ...repository.Option) *userRepository {
	repo := &userRepository{
		config: &postgresConfig{},
		base:   newUserCoreRepository(),
	}

	for _, opt := range opts {
		opt.Apply(repo.config)
	}

	return repo
}

func (repo *userRepository) CreateUser(ctx context.Context, application *model.User) error {
	return repo.base.CreateUser(ctx, repo.config.tx, application)
}

func (repo *userRepository) DeleteUser(ctx context.Context, userId string) error {
	return repo.base.DeleteUser(ctx, repo.config.tx, userId)
}

func (repo *userRepository) GetUser(ctx context.Context, userId string) (*model.User, error) {
	return repo.base.GetUser(ctx, repo.config.tx, userId)
}

func (repo *userRepository) ListUsers(ctx context.Context) ([]*model.User, error) {
	return repo.base.ListUsers(ctx, repo.config.tx)
}

var _ repository.UserRepository = (*userRepository)(nil)
