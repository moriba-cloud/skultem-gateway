package app

import (
	"context"
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/feature"
	"github.com/moriba-cloud/skultem-gateway/domain/permission"
	"github.com/moriba-cloud/skultem-gateway/domain/role"
	"go.uber.org/zap"
	"strings"
)

type (
	aPermission struct {
		repo    permission.Repo
		feature feature.Repo
		role    role.Repo
		logger  *zap.Logger
	}
	argsPermission struct {
		Feature feature.Repo
		Role    role.Repo
		Repo    permission.Repo
		Logger  *zap.Logger
	}
)

func (a aPermission) Update(ctx context.Context, args []*permission.Args, role string) (*ddd.Response[permission.Domain], error) {
	permissions := make([]*permission.Domain, len(args))
	errors := make([]string, 0)

	if _, err := a.role.FindById(role); err != nil {
		return nil, fmt.Errorf("feature '%s' not found", role)
	}

	for i, arg := range args {
		var o *permission.Domain

		if _, err := a.feature.FindById(arg.Feature.Id); err != nil {
			errors = append(errors, fmt.Sprintf("feature '%s' not found", arg.Feature))
		}

		check, err := a.repo.Check(arg.Feature.Id, role)
		if check != nil {
			arg.Aggregation.Id = check.ID()
			arg.Aggregation.CreatedAt = check.CreatedAt()
			arg.Aggregation.UpdatedAt = check.UpdatedAt()
			arg.Aggregation.State = check.State()

			o, err = permission.Existing(*arg)
			if err != nil {
				errors = append(errors, err.Error())
			}
		} else {
			o, err = permission.New(*arg)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}

		permissions[i] = o
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf(strings.Join(errors, ", "))
	}

	return a.repo.Save(permissions, role)
}

func NewPermission(args argsPermission) permission.App {
	return &aPermission{
		repo:    args.Repo,
		feature: args.Feature,
		role:    args.Role,
		logger:  args.Logger,
	}
}
