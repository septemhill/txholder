package logic

import (
	"context"

	"github.com/septemhill/txholder/domain/application"
	"github.com/septemhill/txholder/domain/user"
	"github.com/septemhill/txholder/repository"
	apprepo "github.com/septemhill/txholder/repository/application"
	userrepo "github.com/septemhill/txholder/repository/user"
)

type Logic interface {
	UpdateApplicationAndUser(ctx context.Context, req *UpdateApplicationAndUserRequest) (UpdateApplicationAndUserResponse, error)
}

type logic struct {
	repo              repository.Repository
	repoFactory       repository.RepositoryFactory
	appDomainFactory  func(repo apprepo.ApplicationRepository) application.Domain
	userDoaminFactory func(repo userrepo.UserRepository) user.Domain
}

func (l *logic) UpdateApplicationAndUser(ctx context.Context, req *UpdateApplicationAndUserRequest) (UpdateApplicationAndUserResponse, error) {
	root := l.repoFactory.NewTxRepository()

	appDomain := l.appDomainFactory(root.NewApplicationRepository())
	userDmain := l.userDoaminFactory(root.NewUserRepository())

	if err := appDomain.UpdateApplication(ctx, "Septem", "1.0.0"); err != nil {
		root.Rollback(ctx)
	}

	if err := userDmain.UpdateUser(ctx, "Septem", "bb@test.com"); err != nil {
		root.Rollback(ctx)
	}

	root.Commit(ctx)

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
