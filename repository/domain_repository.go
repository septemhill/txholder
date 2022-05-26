package repository

import (
	appRepo "github.com/septemhill/txholder/repository/application"
	userRepo "github.com/septemhill/txholder/repository/user"
)

type DomainRepositoryFactory interface {
	NewApplicationRepository() appRepo.ApplicationRepository
	NewUserRepository() userRepo.UserRepository
}
