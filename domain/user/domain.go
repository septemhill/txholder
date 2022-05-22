package user

import (
	"context"

	repository "github.com/septemhill/txholder/repository"
)

type Domain interface {
	UpdateUser(ctx context.Context, name, email string) error
}

type DomainFactory func(repo repository.UserRepository) Domain

type userDomain struct {
	repo repository.UserRepository
}

func (d *userDomain) UpdateUser(ctx context.Context, name, email string) error {
	u, err := d.repo.GetUser(ctx, "")
	if err != nil {
		return err
	}

	user := marshalToModel(u)

	if name != "" {
		user.updateName(name)
	}

	if email != "" {
		user.updateEmail(email)
	}

	return d.repo.Update(ctx, user.unmarshalToStorage())
}

func NewDomain(repo repository.UserRepository) *userDomain {
	return &userDomain{
		repo: repo,
	}
}

func NewDomainFactory() DomainFactory {
	return func(repo repository.UserRepository) Domain {
		return NewDomain(repo)
	}
}

var _ Domain = (*userDomain)(nil)
