package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moriba-cloud/skultem-gateway/domain/auth"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/user"
	"strings"
)

func AccessTokenGuard(c *fiber.Ctx) error {
	if c.GetReqHeaders()["Authorization"] != nil {
		authorization := c.GetReqHeaders()["Authorization"][0]
		userToken := strings.Replace(authorization, "Bearer ", "", 1)
		token, err := auth.VerifyAccessToken(userToken)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		id := token["id"].(string)
		record, err := fetchUser(c, id)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		c.Locals("user", record)
		return c.Next()
	}

	return fiber.NewError(fiber.StatusUnauthorized, "missing authentication token")
}

func RefreshTokenGuard(c *fiber.Ctx) error {
	if c.GetReqHeaders()["Authorization"] != nil {
		authorization := c.GetReqHeaders()["Authorization"][0]
		userToken := strings.Replace(authorization, "Bearer ", "", 1)
		_, err := auth.VerifyRefreshToken(userToken)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		c.Set("refresh", userToken)
		return c.Next()
	}

	return fiber.NewError(fiber.StatusUnauthorized, "missing authentication token")
}

func fetchUser(c *fiber.Ctx, id string) (*auth.User, error) {
	app := c.Locals("uApp").(user.App)
	res, err := app.FindById(c.UserContext(), id)
	if err != nil {
		return nil, err
	}

	record := res.Record()
	return &auth.User{
		Id:         record.ID(),
		GivenNames: record.GivenNames(),
		FamilyName: record.FamilyName(),
		Phone:      record.Phone(),
		Email:      record.Email(),
		Role: core.Reference{
			Id:    record.Role().Id,
			Value: record.Role().Value,
		},
		School: record.School(),
		State:  string(record.State()),
	}, nil
}
