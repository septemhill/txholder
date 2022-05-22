package main

import (
	"context"

	"github.com/septemhill/txholder/domain/application"
	"github.com/septemhill/txholder/domain/user"
	"github.com/septemhill/txholder/logic"
	"github.com/septemhill/txholder/repository/postgres"
)

func main() {
	repo := postgres.NewRepository()
	appDomainFactory := application.NewDomainFactory()
	userDomainFactory := user.NewDomainFactory()

	l := logic.NewLogic(repo, appDomainFactory, userDomainFactory)

	l.UpdateApplicationAndUser(context.Background(), &logic.UpdateApplicationAndUserRequest{})
}
