package repository

import (
	"context"

	storage "github.com/septemhill/txholder/storage/user"
)

type UserRepository interface {
	Create(ctx context.Context, user *storage.User) error
	Delete(ctx context.Context, userId string) error
	Update(ctx context.Context, user *storage.User) error
	GetUser(ctx context.Context, userId string) (*storage.User, error)
	ListUsers(ctx context.Context) ([]*storage.User, error)
}

type UserTxHolderRepository interface {
	UserRepository
	TransactionHolderRepository
	TransactionalRepositoryCreator
}
