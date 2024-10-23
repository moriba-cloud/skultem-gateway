package app

import (
	"context"
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/feature"
	"go.uber.org/zap"
)

type (
	aFeature struct {
		repo   feature.Repo
		logger *zap.Logger
	}
	argsFeature struct {
		Repo   feature.Repo
		Logger *zap.Logger
	}
)

func (a aFeature) New(ctx context.Context, args feature.Args) (*ddd.Response[feature.Domain], error) {
	o, err := feature.New(args)
	if err != nil {
		return nil, err
	}

	if _, err := a.repo.Check(o.Name()); err == nil {
		return nil, fmt.Errorf("feature: %s already exists", o.Name())
	}

	record, err := a.repo.Save(*o)
	if err != nil {
		return nil, err
	}

	return ddd.NewResponse[feature.Domain](ddd.ResponseArgs[feature.Domain]{
		Record: record,
	}), nil
}

func (a aFeature) Update(ctx context.Context, args feature.Args) (*ddd.Response[feature.Domain], error) {
	record, err := a.repo.FindById(args.Aggregation.Id)
	if err != nil {
		return nil, err
	}

	if check, _ := a.repo.Check(args.Name); check != nil {
		if check.ID() != record.ID() {
			return nil, fmt.Errorf("feature: %s already exists", args.Name)
		}
	}

	err = record.Update(args)
	if err != nil {
		return nil, err
	}
	record, err = a.repo.Save(*record)

	return ddd.NewResponse[feature.Domain](ddd.ResponseArgs[feature.Domain]{
		Record: record,
	}), nil
}

func (a aFeature) ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[feature.Domain], error) {
	return a.repo.ListByPage(args)
}

func (a aFeature) List(ctx context.Context) (*ddd.Response[core.Option], error) {
	return a.repo.List()
}

func (a aFeature) Remove(ctx context.Context, id string) (*ddd.Response[feature.Domain], error) {
	record, err := a.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	record, err = a.repo.Remove(*record)
	if err != nil {
		return nil, err
	}

	return ddd.NewResponse[feature.Domain](ddd.ResponseArgs[feature.Domain]{
		Record: record,
	}), nil
}

func NewFeature(args argsFeature) feature.App {
	return &aFeature{
		repo:   args.Repo,
		logger: args.Logger,
	}
}
