package user

import (
	"context"

	storage "github.com/septemhill/txholder/storage/user"
)

func (repo *userRepository) GetUser(ctx context.Context, userId string) (*storage.User, error) {
	return nil, nil
}
