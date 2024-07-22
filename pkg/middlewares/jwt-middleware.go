package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/spf13/viper"
	"gogram/config"
	"gogram/pkg/helpers"
)

func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(viper.GetString(config.JwtSecret)),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			res := helpers.NewResponseHelper(
				fiber.StatusUnauthorized,
				"Unauthorized",
				nil,
				nil,
			)

			return c.Status(res.Code).JSON(res)
		},
	})
}
