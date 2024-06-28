package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/markraiter/spycat/internal/config"
	"github.com/markraiter/spycat/internal/lib/jwt"
)

func NewUserIdentity(cfg config.Auth) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "missing Authorization header")
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		uid, err := jwt.ParseToken(tokenString, cfg.SigningKey)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		c.Locals("uid", uid)
		c.Locals("refreshString", tokenString)

		return c.Next()
	}
}
