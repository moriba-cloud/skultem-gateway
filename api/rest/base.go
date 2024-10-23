package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/moriba-build/ose/ddd/rest"
	"github.com/moriba-build/ose/ddd/rest/exception"
	"github.com/moriba-cloud/skultem-gateway/api/rest/routes"
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
	api := rest.New(fiber.Config{
		AppName:       "skultem",
		ServerHeader:  "x-skultem-gateway",
		StrictRouting: true,
		CaseSensitive: true,
		UnescapePath:  true,
		ErrorHandler:  exception.ErrorHandler,
	})

	api.Api().Use(cors.New())
	routes.Routes(routes.Args{
		Route:  api.Api(),
		Apps:   args.Apps,
		Logger: args.Logger,
	})
	api.Start()
}
