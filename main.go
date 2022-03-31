package main

import (
	"context"
	"txholder/repository"
	"txholder/repository/postgres"
)

func TxRepoSample(ctx context.Context, repo repository.Repository) {
	appTxRepo := repo.NewApplicationTxHolderRepository()
	userTxRepo := appTxRepo.NewUserTxHolderRepository()

	appTxRepo.CreateApplication(ctx, nil)
	userTxRepo.CreateUser(ctx, nil)

	userTxRepo.Commit() // or appTxRepo.Commit()
}

func RepoSample(ctx context.Context, repo repository.Repository) {
	appRepo := repo.NewApplicationRepository()
	appRepo.CreateApplication(ctx, nil)
}

func main() {
	ctx := context.Background()
	repo := postgres.NewRepository()

	TxRepoSample(ctx, repo)
	RepoSample(ctx, repo)
}
