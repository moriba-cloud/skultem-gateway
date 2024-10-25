package app

import (
	"context"
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/year"
	"go.uber.org/zap"
)

type (
	aYear struct {
		repo   year.Repo
		bus    year.Bus
		logger *zap.Logger
	}
	argsYear struct {
		Bus    year.Bus
		Repo   year.Repo
		Logger *zap.Logger
	}
)

func (a aYear) New(ctx context.Context, args year.Args) (*ddd.Response[year.Domain], error) {
	o, err := year.New(args)
	if err != nil {
		return nil, err
	}

	if check := a.repo.Check(o.Start(), o.End()); check {
		return nil, fmt.Errorf("academic: %d - %d already exists", o.Start(), o.End())
	}

	record, err := a.repo.Save(*o)
	if err != nil {
		return nil, err
	}

	a.bus.TakeOff(ctx, *record)

	return ddd.NewResponse[year.Domain](ddd.ResponseArgs[year.Domain]{
		Record: record,
	}), nil
}

func (a aYear) One(ctx context.Context, id string) (*ddd.Response[year.Domain], error) {
	record, err := a.repo.OneById(id)
	if err != nil {
		return nil, err
	}
	return ddd.NewResponse[year.Domain](ddd.ResponseArgs[year.Domain]{
		Record: record,
	}), nil
}

func (a aYear) ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[year.Domain], error) {
	return a.repo.ListByPage(args)
}

func (a aYear) List(ctx context.Context) (*ddd.Response[core.Option], error) {
	return a.repo.List()
}

func (a aYear) Active(ctx context.Context, args ddd.FindByArgs) (*ddd.Response[year.Domain], error) {
	//TODO implement me
	panic("implement me")
}

func NewYear(args argsYear) year.App {
	return &aYear{
		repo:   args.Repo,
		logger: args.Logger,
		bus:    args.Bus,
	}
}
