package app

import (
	"context"
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-management/domain/core"
	"github.com/moriba-cloud/skultem-management/domain/values"
	"go.uber.org/zap"
)

type aValues struct {
	repo   values.Repo
	logger *zap.Logger
}

func (a aValues) New(ctx context.Context, args values.Args) (*ddd.Response[values.Domain], error) {
	payload, err := values.New(args)
	if err != nil {
		return nil, err
	}

	if check := a.repo.Check(payload.Key(), payload.Batch()); check {
		return nil, fmt.Errorf("value: %s already exists within this group: %s", payload.Key(), payload.Batch())
	}

	record, err := a.repo.Save(payload)
	if err != nil {
		return nil, err
	}

	return ddd.NewResponse[values.Domain](ddd.ResponseArgs[values.Domain]{
		Record: record,
	}), nil
}

func (a aValues) One(ctx context.Context, id string) (*ddd.Response[values.Domain], error) {
	record, err := a.repo.OneById(id)
	if err != nil {
		return nil, err
	}
	return ddd.NewResponse[values.Domain](ddd.ResponseArgs[values.Domain]{
		Record: record,
	}), nil
}

func (a aValues) List(ctx context.Context) (*ddd.Response[core.Option], error) {
	user := core.ActiveUser(ctx)
	return a.repo.List(user.School)
}

func (a aValues) ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[values.Domain], error) {
	user := core.ActiveUser(ctx)
	return a.repo.ListByPage(args, user.School)
}

func (a aValues) ListByGroup(ctx context.Context, group values.Batch) (*ddd.Response[core.Option], error) {
	user := core.ActiveUser(ctx)
	return a.repo.ListByGroup(group, user.School)
}

func NewValues(repo values.Repo, logger *zap.Logger) values.App {
	return &aValues{
		repo:   repo,
		logger: logger,
	}
}
