package app

import (
	"context"
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/role"
	"github.com/moriba-cloud/skultem-gateway/domain/user"
	"go.uber.org/zap"
)

type (
	aUser struct {
		repo   user.Repo
		logger *zap.Logger
	}
	argsUser struct {
		Repo   user.Repo
		Logger *zap.Logger
	}
)

func (a aUser) New(ctx context.Context, args user.Args) (*ddd.Response[user.Domain], error) {
	o, err := user.New(args)
	if err != nil {
		return nil, err
	}

	if _, err := a.repo.Check(o.Phone()); err == nil {
		return nil, fmt.Errorf("user with this phone: %d already exists", o.Phone())
	}

	record, err := a.repo.Save(*o)
	if err != nil {
		return nil, err
	}

	return ddd.NewResponse(ddd.ResponseArgs[user.Domain]{
		Record: record,
	}), nil
}

func (a aUser) Update(ctx context.Context, args user.Args) (*ddd.Response[user.Domain], error) {
	record, err := a.repo.FindById(args.Aggregation.Id)
	if err != nil {
		return nil, err
	}

	if check, _ := a.repo.Check(args.Phone); check != nil {
		if check.ID() != record.ID() {
			return nil, fmt.Errorf("user with this phone: %d already exists", args.Phone)
		}
	}

	err = record.Update(args)
	if err != nil {
		return nil, err
	}
	record, err = a.repo.Save(*record)

	return ddd.NewResponse(ddd.ResponseArgs[role.Domain]{
		Record: record,
	}), nil
}

func (a aUser) ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[role.Domain], error) {
	return a.repo.ListByPage(args)
}

func (a aUser) List(ctx context.Context) (*ddd.Response[core.Option], error) {
	return a.repo.List()
}

func (a aUser) Remove(ctx context.Context, id string) (*ddd.Response[role.Domain], error) {
	record, err := a.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	record, err = a.repo.Remove(*record)
	if err != nil {
		return nil, err
	}

	return ddd.NewResponse(ddd.ResponseArgs[role.Domain]{
		Record: record,
	}), nil
}

func NewRole(args argsRole) role.App {
	return &aUser{
		repo:   args.Repo,
		logger: args.Logger,
	}
}
