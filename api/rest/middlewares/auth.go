package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/moriba-cloud/skultem-gateway/domain/auth"
	"github.com/moriba-cloud/skultem-gateway/domain/user"
	"strings"
)

func AccessTokenGuard(c *fiber.Ctx) error {
	if c.GetReqHeaders()["Authorization"] != nil {
		authorization := c.GetReqHeaders()["Authorization"][0]
		userToken := strings.Replace(authorization, "Bearer ", "", 1)
		token, err := auth.VerifyAccessToken(userToken)
		if err != nil {
			return err
		}

		id := token["id"].(string)
		record, err := fetchUser(c, id)
		if err != nil {
			return err
		}

		c.Locals("user", record.ID())
		c.Locals("school", record.School())
		c.Locals("role", record.Role().Id)
		return c.Next()
	}

	return fmt.Errorf("missing authentication token")
}

func RefreshTokenGuard(c *fiber.Ctx) error {
	if c.GetReqHeaders()["Authorization"] != nil {
		authorization := c.GetReqHeaders()["Authorization"][0]
		userToken := strings.Replace(authorization, "Bearer ", "", 1)
		_, err := auth.VerifyRefreshToken(userToken)
		if err != nil {
			return err
		}

		c.Set("refresh", userToken)

		return c.Next()
	}

	return fmt.Errorf("missing authentication token")
}

func fetchUser(c *fiber.Ctx, id string) (*user.Domain, error) {
	app := c.Locals("uApp").(user.App)
	record, err := app.FindById(c.UserContext(), id)
	if err != nil {
		return nil, err
	}

	return record.Record(), nil
}
