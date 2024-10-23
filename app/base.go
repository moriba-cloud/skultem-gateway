package app

import (
	"github.com/moriba-cloud/skultem-gateway/domain/feature"
	"github.com/moriba-cloud/skultem-gateway/domain/year"
	"github.com/moriba-cloud/skultem-gateway/infra/management"
	"github.com/moriba-cloud/skultem-gateway/infra/psql/repos"
	"go.uber.org/zap"
)

type (
	Apps struct {
		Year    year.App
		Feature feature.App
	}
	Args struct {
		Repos  *repos.Repos
		Logger *zap.Logger
		Bus    management.Bus
	}
)

func NewApps(args Args) *Apps {
	return &Apps{
		Year: NewYear(yArgs{
			Bus:    args.Bus.Year,
			Repo:   args.Repos.Year,
			Logger: args.Logger,
		}),
		Feature: NewFeature(argsFeature{
			Repo:   args.Repos.Feature,
			Logger: args.Logger,
		}),
	}
}
