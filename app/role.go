package app

import (
	"context"
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/role"
	"go.uber.org/zap"
)

type (
	aRole struct {
		repo   role.Repo
		logger *zap.Logger
	}
	argsRole struct {
		Repo   role.Repo
		Logger *zap.Logger
	}
)

func (a aRole) New(ctx context.Context, args role.Args) (*ddd.Response[role.Domain], error) {
	o, err := role.New(args)
	if err != nil {
		return nil, err
	}

	if _, err := a.repo.Check(o.Name()); err == nil {
		return nil, fmt.Errorf("role: %s already exists", o.Name())
	}

	record, err := a.repo.Save(*o)
	if err != nil {
		return nil, err
	}

	return ddd.NewResponse[role.Domain](ddd.ResponseArgs[role.Domain]{
		Record: record,
	}), nil
}

func (a aRole) Update(ctx context.Context, args role.Args) (*ddd.Response[role.Domain], error) {
	record, err := a.repo.FindById(args.Aggregation.Id)
	if err != nil {
		return nil, err
	}

	if check, _ := a.repo.Check(args.Name); check != nil {
		if check.ID() != record.ID() {
			return nil, fmt.Errorf("role: %s already exists", args.Name)
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

func (a aRole) ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[role.Domain], error) {
	return a.repo.ListByPage(args)
}

func (a aRole) List(ctx context.Context) (*ddd.Response[core.Option], error) {
	return a.repo.List()
}

func (a aRole) Remove(ctx context.Context, id string) (*ddd.Response[role.Domain], error) {
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
	return &aRole{
		repo:   args.Repo,
		logger: args.Logger,
	}
}
