package user

import (
	"context"

	storage "github.com/septemhill/txholder/storage/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *storage.User) error
	Update(ctx context.Context, user *storage.User) error
	GetUser(ctx context.Context, userId string) (*storage.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

var _ UserRepository = (*userRepository)(nil)
