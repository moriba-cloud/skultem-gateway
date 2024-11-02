package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moriba-build/ose/ddd/rest/dto"
	"github.com/moriba-build/ose/ddd/rest/validation"
	"github.com/moriba-cloud/skultem-gateway/api/rest/middlewares"
	"github.com/moriba-cloud/skultem-gateway/domain/auth"
	"go.uber.org/zap"
)

type (
	ApiAuth struct {
		validation *validation.Validation
		app        auth.App
		logger     *zap.Logger
	}
	Auth struct {
		Access  string `json:"access"`
		Refresh string `json:"refresh"`
	}
	AuthUser struct {
		Id         string
		GivenNames string
		FamilyName string
		Phone      int
		Email      string
		Role       Reference
		School     string
		State      string
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

func AuthUserResponse(o *auth.User) *AuthUser {
	return &AuthUser{
		Id:         o.Id,
		GivenNames: o.GivenNames,
		FamilyName: o.FamilyName,
		Phone:      o.Phone,
		Email:      o.Email,
		Role: Reference{
			Id:    o.Role.Id,
			Value: o.Role.Value,
		},
		School: o.School,
		State:  o.State,
	}
}

func (a ApiAuth) login(c *fiber.Ctx) error {
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

func (a ApiAuth) access(c *fiber.Ctx) error {
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

func (a ApiAuth) me(c *fiber.Ctx) error {
	res, err := a.app.Me(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	record := AuthUserResponse(res.Record())
	return c.JSON(dto.NewResponse(dto.ResponseArgs[AuthUser]{
		Record: record,
	}))
}

func AuthRoute(api fiber.Router, app auth.App, logger *zap.Logger) {
	r := &ApiAuth{
		app:        app,
		validation: validation.NewValidation(),
		logger:     logger,
	}

	api.Group("/auth").
		Post("login", r.login).
		Get("refresh", middlewares.RefreshTokenGuard, r.access).
		Get("me", middlewares.AccessTokenGuard, r.me)
}
