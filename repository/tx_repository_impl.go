package repository

import (
	"context"

	appRepo "github.com/septemhill/txholder/repository/application"
	userRepo "github.com/septemhill/txholder/repository/user"
	"gorm.io/gorm"
)

type txRepository struct {
	tx *gorm.DB
}

func (factory *txRepository) NewApplicationRepository() appRepo.ApplicationRepository {
	return nil
}

func (factory *txRepository) NewUserRepository() userRepo.UserRepository {
	return nil
}

func (factory *txRepository) Commit(ctx context.Context) error {
	return nil
}

func (factory *txRepository) Rollback(ctx context.Context) error {
	return nil
}

func NewTxRepository(tx *gorm.DB) *txRepository {
	return &txRepository{
		tx: tx,
	}
}

var _ TxRepository = (*txRepository)(nil)
var _ Transactional = (*txRepository)(nil)
