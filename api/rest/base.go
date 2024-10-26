package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/moriba-build/ose/ddd/config"
	"github.com/moriba-build/ose/ddd/rest"
	"github.com/moriba-build/ose/ddd/rest/exception"
	"github.com/moriba-cloud/skultem-gateway/api/rest/routes"
	"github.com/moriba-cloud/skultem-gateway/api/rest/routes/middlewares"
	"github.com/moriba-cloud/skultem-gateway/app"
	"go.uber.org/zap"
)

type (
	Args struct {
		Apps   *app.Apps
		Logger *zap.Logger
	}
)

func Api(args Args) {
	x := config.NewEnvs().EnvStr("X_HEADER")

	api := rest.New(fiber.Config{
		AppName:       "skultem",
		ServerHeader:  x,
		StrictRouting: true,
		CaseSensitive: true,
		UnescapePath:  true,
		ErrorHandler:  exception.ErrorHandler,
	})

	api.Api().Use(cors.New())
	api.Api().Use(middlewares.Auth)
	routes.Routes(routes.Args{
		Route:  api.Api(),
		Apps:   args.Apps,
		Logger: args.Logger,
	})
	api.Start()
}
