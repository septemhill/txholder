package repository

import (
	appRepo "github.com/septemhill/txholder/repository/application"
	userRepo "github.com/septemhill/txholder/repository/user"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (factory *repository) NewApplicationRepository() appRepo.ApplicationRepository {
	return appRepo.NewApplicationRepository(factory.db)
}

func (factory *repository) NewUserRepository() userRepo.UserRepository {
	return userRepo.NewUserRepository(factory.db.Begin())
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

var _ Repository = (*repository)(nil)
