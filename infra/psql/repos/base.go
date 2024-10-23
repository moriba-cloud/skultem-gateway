package repos

import (
	"github.com/moriba-cloud/skultem-gateway/domain/feature"
	"github.com/moriba-cloud/skultem-gateway/domain/year"
	featureModel "github.com/moriba-cloud/skultem-gateway/infra/psql/repos/feature"
	yearModel "github.com/moriba-cloud/skultem-gateway/infra/psql/repos/year"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type (
	Repos struct {
		Year    year.Repo
		Feature feature.Repo
	}
	Args struct {
		Db     *gorm.DB
		Logger *zap.Logger
	}
)

func NewRepos(args Args) *Repos {
	return &Repos{
		Year:    yearModel.New(args.Db, args.Logger),
		Feature: featureModel.New(args.Db, args.Logger),
	}
}
