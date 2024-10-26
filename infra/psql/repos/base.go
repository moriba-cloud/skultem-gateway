package repos

import (
	"github.com/moriba-cloud/skultem-gateway/domain/feature"
	"github.com/moriba-cloud/skultem-gateway/domain/permission"
	"github.com/moriba-cloud/skultem-gateway/domain/role"
	"github.com/moriba-cloud/skultem-gateway/domain/user"
	"github.com/moriba-cloud/skultem-gateway/domain/year"
	featureModel "github.com/moriba-cloud/skultem-gateway/infra/psql/repos/feature"
	permissionModel "github.com/moriba-cloud/skultem-gateway/infra/psql/repos/permission"
	roleModel "github.com/moriba-cloud/skultem-gateway/infra/psql/repos/role"
	userModel "github.com/moriba-cloud/skultem-gateway/infra/psql/repos/user"
	yearModel "github.com/moriba-cloud/skultem-gateway/infra/psql/repos/year"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type (
	Repos struct {
		Year       year.Repo
		Feature    feature.Repo
		Role       role.Repo
		Permission permission.Repo
		User       user.Repo
	}
	Args struct {
		Db     *gorm.DB
		Logger *zap.Logger
	}
)

func NewRepos(args Args) *Repos {
	return &Repos{
		Year:       yearModel.New(args.Db, args.Logger),
		Feature:    featureModel.New(args.Db, args.Logger),
		Role:       roleModel.New(args.Db, args.Logger),
		Permission: permissionModel.New(args.Db, args.Logger),
		User:       userModel.New(args.Db, args.Logger),
	}
}
