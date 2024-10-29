package app

import (
	"context"
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/config"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/school"
	"github.com/moriba-cloud/skultem-gateway/domain/user"
	"go.uber.org/zap"
)

type (
	aSchool struct {
		repo   school.Repo
		user   user.Repo
		logger *zap.Logger
	}
	argsSchool struct {
		Repo   school.Repo
		User   user.Repo
		Logger *zap.Logger
	}
)

func (a aSchool) New(ctx context.Context, args school.Args) (*ddd.Response[school.Domain], error) {
	if _, err := a.repo.Check("name", args.Name); err == nil {
		return nil, fmt.Errorf("school already exist with this name: %s", args.Name)
	}

	if _, err := a.repo.Check("domain", args.Domain); err == nil {
		return nil, fmt.Errorf("school already exist with this domain: %s", args.Domain)
	}

	if _, err := a.user.CheckByPhone(args.Owner.Phone); err == nil {
		return nil, fmt.Errorf("owner already exist with this phone: %d", args.Owner.Phone)
	}

	if _, err := a.user.CheckByEmail(args.Owner.Email); err == nil {
		return nil, fmt.Errorf("owner already exist with this email: %s", args.Owner.Email)
	}

	role := config.NewEnvs().EnvStr("OWNER_ROLE")
	owner := user.Args{
		GivenNames: args.Owner.GivenNames,
		FamilyName: args.Owner.FamilyName,
		Phone:      args.Owner.Phone,
		Email:      args.Owner.Email,
		Role: core.Reference{
			Id: role,
		},
	}
	args.Owner = owner

	o, err := school.New(args)
	if err != nil {
		return nil, err
	}
	return a.repo.Save(*o)
}

func (a aSchool) FindById(ctx context.Context, id string) (*ddd.Response[school.Domain], error) {
	return a.repo.FindById(id)
}

func (a aSchool) Remove(ctx context.Context, id string) (*ddd.Response[school.Domain], error) {
	domain, err := a.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	record, err := a.repo.Remove(*domain.Record())
	if err != nil {
		return nil, err
	}

	return ddd.NewResponse(ddd.ResponseArgs[school.Domain]{
		Record: record,
	}), nil
}

func (a aSchool) Update(ctx context.Context, args school.Args) (*ddd.Response[school.Domain], error) {
	check, err := a.repo.FindById(args.Aggregation.Id)
	if err != nil {
		return nil, err
	}
	record := check.Record()

	err = record.Update(args)
	if err != nil {
		return nil, err
	}
	return a.repo.Save(*record)
}

func (a aSchool) ListByPage(ctx context.Context, args ddd.PaginationArgs) (*ddd.Response[school.Domain], error) {
	return a.repo.ListByPage(args)
}

func (a aSchool) List(ctx context.Context) (*ddd.Response[core.Option], error) {
	return a.repo.List()
}

func NewSchool(args argsSchool) school.App {
	return &aSchool{
		user:   args.User,
		repo:   args.Repo,
		logger: args.Logger,
	}
}
