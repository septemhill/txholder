package application

import (
	"context"

	repository "github.com/septemhill/txholder/repository/application"
)

type Domain interface {
	UpdateApplication(ctx context.Context, name, version string) error
}

type DomainFactory func(repo repository.ApplicationRepository) Domain

type applicationDomain struct {
	repo repository.ApplicationRepository
}

func (a *applicationDomain) UpdateApplication(ctx context.Context, name, version string) error {
	ap, err := a.repo.GetApplication(ctx, "")
	if err != nil {
		return err
	}

	app := marshalToModel(ap)

	if name != "" {
		app.updateName(name)
	}

	if version != "" {
		app.updateVersion(version)
	}

	return a.repo.Update(ctx, app.unmarshalToStorage())
}

func NewDomain(repo repository.ApplicationRepository) *applicationDomain {
	return &applicationDomain{
		repo: repo,
	}
}

func NewDomainFactory() DomainFactory {
	return func(repo repository.ApplicationRepository) Domain {
		return NewDomain(repo)
	}
}

var _ Domain = (*applicationDomain)(nil)
