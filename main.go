package main

import (
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/api/rest"
	"github.com/moriba-cloud/skultem-gateway/app"
	"github.com/moriba-cloud/skultem-gateway/infra/management"
	"github.com/moriba-cloud/skultem-gateway/infra/psql"
)

func main() {
	logger, err := ddd.NewLogger()
	if err != nil {
		panic(err)
	}

	db := psql.Database(logger)
	managementService := management.NeeService(management.Args{
		Logger: logger,
	})
	apps := app.NewApps(app.Args{
		Repos:      db,
		Logger:     logger,
		Management: managementService,
	})

	rest.Api(rest.Args{
		Apps:   apps,
		Logger: logger,
	})
}
