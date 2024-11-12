package app

import (
	"context"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/values"
	"go.uber.org/zap"
)

type (
	aValues struct {
		service values.Service
		logger  *zap.Logger
	}
	argsValue struct {
		Service values.Service
		Logger  *zap.Logger
	}
)

func (a aValues) New(ctx context.Context, args values.Args) (*ddd.Response[values.Domain], error) {
	payload, err := values.New(args)
	if err != nil {
		return nil, err
	}

	return a.service.Save(ctx, *payload)
}

func (a aValues) List(ctx context.Context) (*ddd.Response[core.Option], error) {
	return a.service.List(ctx)
}

func (a aValues) ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[values.Domain], error) {
	return a.service.ListByPage(ctx, args)
}

func (a aValues) ListByGroup(ctx context.Context, batch values.Batch) (*ddd.Response[core.Option], error) {
	return a.service.ListByBatch(ctx, batch)
}

func NewValues(args argsValue) values.App {
	return &aValues{
		service: args.Service,
		logger:  args.Logger,
	}
}
