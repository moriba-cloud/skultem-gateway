package app

import (
	"context"
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
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
	school := core.GetContextValue(ctx, "school")
	args.School = school

	o, err := user.New(args)
	if err != nil {
		return nil, err
	}

	if _, err := a.repo.CheckByPhone(args.Phone); err != nil {
		return nil, fmt.Errorf("user already exist with this phone: %d", args.Phone)
	}
	if _, err := a.repo.CheckByEmail(args.Email); err != nil {
		return nil, fmt.Errorf("user already exist with this email: %d", args.Email)
	}

	record, err := a.repo.Save(*o)
	if err != nil {
		return nil, err
	}

	return ddd.NewResponse(ddd.ResponseArgs[user.Domain]{
		Record: record,
	}), nil
}

func (a aUser) FindById(ctx context.Context, id string) (*ddd.Response[user.Domain], error) {
	return a.repo.FindById(id)
}

func (a aUser) Remove(ctx context.Context, id string) (*ddd.Response[user.Domain], error) {
	domain, err := a.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	record, err := a.repo.Remove(*domain.Record())
	if err != nil {
		return nil, err
	}

	return ddd.NewResponse(ddd.ResponseArgs[user.Domain]{
		Record: record,
	}), nil
}

func (a aUser) Update(ctx context.Context, args user.Args) (*ddd.Response[user.Domain], error) {
	check, err := a.repo.FindById(args.Aggregation.Id)
	if err != nil {
		return nil, err
	}
	record := check.Record()

	if _, err := a.repo.CheckByPhone(args.Phone); err != nil {
		return nil, fmt.Errorf("user already exist with this phone: %d", args.Phone)
	}
	if _, err := a.repo.CheckByEmail(args.Email); err != nil {
		return nil, fmt.Errorf("user already exist with this email: %d", args.Email)
	}

	err = record.Update(args)
	if err != nil {
		return nil, err
	}
	record, err = a.repo.Save(*record)

	return ddd.NewResponse(ddd.ResponseArgs[user.Domain]{
		Record: record,
	}), nil
}

func (a aUser) ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[user.Domain], error) {
	school := core.GetContextValue(ctx, "school")
	return a.repo.ListByPage(args, school)
}

func (a aUser) List(ctx context.Context) (*ddd.Response[core.Option], error) {
	return a.repo.List()
}

func NewUser(args argsUser) user.App {
	return &aUser{
		repo:   args.Repo,
		logger: args.Logger,
	}
}
