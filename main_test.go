package main

import (
	"context"
	"testing"
	"txholder/model"
	"txholder/repository"
)

func TestTxRepoSample(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		repo       *repository.MockRepository
		appTxRepo  *repository.MockApplicationTxHolderRepository
		userTxRepo *repository.MockUserTxHolderRepository
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		mock   func(fields, args)
	}{
		{
			name: "sample-test",
			fields: fields{
				repo:       &repository.MockRepository{},
				appTxRepo:  &repository.MockApplicationTxHolderRepository{},
				userTxRepo: &repository.MockUserTxHolderRepository{},
			},
			args: args{
				ctx: ctx,
			},
			mock: func(fields fields, args args) {
				fields.repo.On("NewApplicationTxHolderRepository").Return(fields.appTxRepo).Once()
				fields.appTxRepo.On("NewUserTxHolderRepository").Return(fields.userTxRepo).Once()
				fields.appTxRepo.On("CreateApplication", args.ctx, (*model.Application)(nil)).Return(nil).Once()
				fields.userTxRepo.On("CreateUser", args.ctx, (*model.User)(nil)).Return(nil).Once()
				fields.userTxRepo.On("Commit").Return(nil).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock(tt.fields, tt.args)
			}
			TxRepoSample(tt.args.ctx, tt.fields.repo)
		})
	}
}

func TestRepoSample(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		repo    *repository.MockRepository
		appRepo *repository.MockApplicationRepository
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		mock   func(fields, args)
	}{
		{
			name: "sample-test",
			fields: fields{
				repo:    &repository.MockRepository{},
				appRepo: &repository.MockApplicationRepository{},
			},
			args: args{
				ctx: ctx,
			},
			mock: func(fields fields, args args) {
				fields.repo.On("NewApplicationRepository").Return(fields.appRepo).Once()
				fields.appRepo.On("CreateApplication", args.ctx, (*model.Application)(nil)).Return(nil).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock(tt.fields, tt.args)
			}
			RepoSample(tt.args.ctx, tt.fields.repo)
		})
	}
}
