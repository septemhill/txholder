package logic

import (
	"context"

	"github.com/septemhill/txholder/domain/application"
	"github.com/septemhill/txholder/domain/user"
	"github.com/septemhill/txholder/repository"
)

type Logic interface {
	UpdateApplicationAndUser(ctx context.Context, req *UpdateApplicationAndUserRequest) (UpdateApplicationAndUserResponse, error)
}

type logic struct {
	repo              repository.Repository
	appDomainFactory  func(repo repository.ApplicationRepository) application.Domain
	userDoaminFactory func(repo repository.UserRepository) user.Domain
}

func (l *logic) UpdateApplicationAndUser(ctx context.Context, req *UpdateApplicationAndUserRequest) (UpdateApplicationAndUserResponse, error) {
	appRepo := l.repo.NewApplicationTxHolderRepository()
	userRepo := appRepo.NewUserTxHolderRepository()

	appDomain := l.appDomainFactory(appRepo)
	userDmain := l.userDoaminFactory(userRepo)

	if err := appDomain.UpdateApplication(ctx, "Septem", "1.0.0"); err != nil {
		appRepo.Rollback()
	}

	if err := userDmain.UpdateUser(ctx, "Septem", "bb@test.com"); err != nil {
		userRepo.Rollback()
	}

	appRepo.Commit()

	return UpdateApplicationAndUserResponse{}, nil
}

func NewLogic(repo repository.Repository, appDomainFactory application.DomainFactory, userDomainFactory user.DomainFactory) *logic {
	return &logic{
		repo:              repo,
		appDomainFactory:  appDomainFactory,
		userDoaminFactory: userDomainFactory,
	}
}

var _ Logic = (*logic)(nil)
