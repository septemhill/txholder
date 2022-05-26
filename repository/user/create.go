package user

import (
	"context"

	storage "github.com/septemhill/txholder/storage/user"
)

func (repo *userRepository) Create(ctx context.Context, user *storage.User) error {
	return nil
}
