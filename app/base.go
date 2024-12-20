package app

import (
	"github.com/moriba-cloud/skultem-gateway/domain/auth"
	"github.com/moriba-cloud/skultem-gateway/domain/feature"
	"github.com/moriba-cloud/skultem-gateway/domain/permission"
	"github.com/moriba-cloud/skultem-gateway/domain/role"
	"github.com/moriba-cloud/skultem-gateway/domain/school"
	"github.com/moriba-cloud/skultem-gateway/domain/user"
	"github.com/moriba-cloud/skultem-gateway/domain/values"
	"github.com/moriba-cloud/skultem-gateway/domain/year"
	"github.com/moriba-cloud/skultem-gateway/infra/management"
	"github.com/moriba-cloud/skultem-gateway/infra/psql/repos"
	"go.uber.org/zap"
)

type (
	Apps struct {
		Year       year.App
		Value      values.App
		Feature    feature.App
		Role       role.App
		Permission permission.App
		User       user.App
		Auth       auth.App
		School     school.App
	}
	Args struct {
		Repos      *repos.Repos
		Management management.Services
		Logger     *zap.Logger
	}
)

func NewApps(args Args) *Apps {
	return &Apps{
		Year: NewYear(argsYear{
			Service: args.Management.Year,
			Logger:  args.Logger,
		}),
		Value: NewValues(argsValue{
			Service: args.Management.Value,
			Logger:  args.Logger,
		}),
		Feature: NewFeature(argsFeature{
			Repo:   args.Repos.Feature,
			Logger: args.Logger,
		}),
		Role: NewRole(argsRole{
			Repo:       args.Repos.Role,
			Permission: args.Repos.Permission,
			Logger:     args.Logger,
		}),
		Permission: NewPermission(argsPermission{
			Repo:    args.Repos.Permission,
			Role:    args.Repos.Role,
			Feature: args.Repos.Feature,
			Logger:  args.Logger,
		}),
		User: NewUser(argsUser{
			Repo:   args.Repos.User,
			Logger: args.Logger,
		}),
		Auth: NewAuth(argsAuth{
			Repo:   args.Repos.User,
			Logger: args.Logger,
		}),
		School: NewSchool(argsSchool{
			Repo:   args.Repos.School,
			User:   args.Repos.User,
			Logger: args.Logger,
		}),
	}
}
