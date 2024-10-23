package psql

import (
	"github.com/moriba-build/ose/ddd/psql"
	"github.com/moriba-cloud/skultem-gateway/infra/psql/repos"
	"github.com/moriba-cloud/skultem-gateway/infra/psql/repos/feature"
	"github.com/moriba-cloud/skultem-gateway/infra/psql/repos/year"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Database(logger *zap.Logger) *repos.Repos {
	db := psql.New()
	db.Connect()

	migrate(db.DB())
	return repos.NewRepos(repos.Args{
		Db:     db.DB(),
		Logger: logger,
	})
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&year.Year{}, feature.Feature{},
	)
	if err != nil {
		panic(err)
	}
}
