package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moriba-build/ose/ddd/rest/dto"
	"github.com/moriba-build/ose/ddd/rest/validation"
	"github.com/moriba-cloud/skultem-gateway/api/rest/routes/middlewares"
	"github.com/moriba-cloud/skultem-gateway/domain/auth"
	"go.uber.org/zap"
)

type (
	apiAuth struct {
		validation *validation.Validation
		app        auth.App
		logger     *zap.Logger
	}
	Auth struct {
		Access  string `json:"access"`
		Refresh string `json:"refresh"`
	}
	AuthRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
)

func AuthResponse(o *auth.Domain) *Auth {
	return &Auth{
		Access:  o.Access(),
		Refresh: o.Refresh(),
	}
}

func (a apiAuth) login(c *fiber.Ctx) error {
	payload := new(AuthRequest)
	if err := c.BodyParser(payload); err != nil {
		return err
	}
	if err := a.validation.Run(payload); err != nil {
		return err
	}

	res, err := a.app.Login(c.Context(), payload.Email, payload.Password)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	record := AuthResponse(res.Record())

	return c.JSON(dto.NewResponse(dto.ResponseArgs[Auth]{
		Record: record,
	}))
}

func (a apiAuth) access(c *fiber.Ctx) error {
	refresh := c.Get("refresh")
	res, err := a.app.Access(c.Context(), refresh)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	record := AuthResponse(res.Record())

	return c.JSON(dto.NewResponse(dto.ResponseArgs[Auth]{
		Record: record,
	}))
}

func AuthRoute(api fiber.Router, app auth.App, logger *zap.Logger) {
	r := &apiAuth{
		app:        app,
		validation: validation.NewValidation(),
		logger:     logger,
	}

	api.Group("/auth").
		Post("login", r.login).
		Get("refresh", middlewares.RefreshTokenGuard, r.access)
}
