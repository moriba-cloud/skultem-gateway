package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/moriba-cloud/skultem-gateway/domain/auth"
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

		c.Set("user", token["id"].(string))
		return c.Next()
	}

	return fmt.Errorf("missing authentication token")
}

func RefreshTokenGuard(c *fiber.Ctx) error {
	if c.GetReqHeaders()["Authorization"] != nil {
		authorization := c.GetReqHeaders()["Authorization"][0]
		userToken := strings.Replace(authorization, "Bearer ", "", 1)
		token, err := auth.VerifyRefreshToken(userToken)

		if err != nil {
			return err
		}

		c.Set("user", token["id"].(string))
		c.Set("refresh", userToken)
		return c.Next()
	}

	return fmt.Errorf("missing authentication token")
}
