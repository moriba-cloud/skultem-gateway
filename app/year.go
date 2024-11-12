package app

import (
	"context"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/year"
	"go.uber.org/zap"
)

type (
	aYear struct {
		service year.Service
		logger  *zap.Logger
	}
	argsYear struct {
		Service year.Service
		Logger  *zap.Logger
	}
)

func (a aYear) New(ctx context.Context, args year.Args) (*ddd.Response[year.Domain], error) {
	o, err := year.New(args)
	if err != nil {
		return nil, err
	}
	return a.service.Save(ctx, *o)
}

func (a aYear) One(ctx context.Context, id string) (*ddd.Response[year.Domain], error) {
	return nil, nil
}

func (a aYear) ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[year.Domain], error) {
	return a.service.ListByPage(ctx, args)
}

func (a aYear) List(ctx context.Context) (*ddd.Response[core.Option], error) {
	return nil, nil
}

func (a aYear) Active(ctx context.Context, args ddd.FindByArgs) (*ddd.Response[year.Domain], error) {
	//TODO implement me
	panic("implement me")
}

func NewYear(args argsYear) year.App {
	return &aYear{
		service: args.Service,
		logger:  args.Logger,
	}
}
