package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moriba-build/ose/ddd/rest/dto"
	"github.com/moriba-build/ose/ddd/rest/validation"
	"github.com/moriba-cloud/skultem-gateway/domain/user"
	"go.uber.org/zap"
)

type (
	AuthRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
)

func (a apiUser) login(c *fiber.Ctx) error {
	payload := new(AuthRequest)
	if err := c.BodyParser(payload); err != nil {
		return err
	}
	if err := a.validation.Run(payload); err != nil {
		return err
	}

	res, err := a.app.Login(c.Context(), user.AuthArgs{
		Email:    payload.Email,
		Password: payload.Password,
	})

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	record := UserResponse(res.Record())
	record.AccessToken = res.Record().Access()
	record.RefreshToken = res.Record().Access()

	return c.JSON(dto.NewResponse(dto.ResponseArgs[User]{
		Record: record,
	}))
}

func AuthRoute(api fiber.Router, app user.App, logger *zap.Logger) {
	r := &apiUser{
		app:        app,
		validation: validation.NewValidation(),
		logger:     logger,
	}

	api.Group("/auth").
		Post("login", r.login)
}
