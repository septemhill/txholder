package repository

import (
	"context"
	"txholder/model"
)

//go:generate mockery --name=UserRepository --inpackage --case underscore
type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, userId string) error
	GetUser(ctx context.Context, userId string) (*model.User, error)
	ListUsers(ctx context.Context) ([]*model.User, error)
}

//go:generate mockery --name=UserTxHolderRepository --inpackage --case underscore
type UserTxHolderRepository interface {
	UserRepository
	TransactionHolderRepository
	TransactionalRepositoryCreator
}
